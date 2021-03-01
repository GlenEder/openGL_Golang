package main

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"

	//"github.com/go-gl/mathgl/mgl32"
	"openGL_Golang/shaders"
	"runtime"
	"strings"
)

const (
	//window width and height
	WIDTH = 480
	HEIGHT = 480
)

//Points of triangle that are used for rendering our first triangle
var trianglePoints = []float32{
	-1.0, -1.0, 0.0,
	1.0, -1.0, 0.0,
	0.0, 1.0, 0.0,
}


func init() {
	//lock onto main thread for glfw and gl calls
	runtime.LockOSThread()

	//Initialize glfw
	if err := glfw.Init(); err != nil {
		panic(err)
	}
}

func main() {

	//generate window
	window := createMainWindow()
	defer glfw.Terminate()

	//init opengl now that we have a window in context
	program := initOpenGL()

	setupView(program)

	//Ensure that we can read input
	window.SetInputMode(glfw.StickyKeysMode, glfw.True)

	//load up our triangle
	triangle := genVAO(trianglePoints)

	//render loop
	for !window.ShouldClose() && window.GetKey(glfw.KeyEscape) != glfw.Press {
		draw(window, program, triangle)
	}

}

//Render method
func draw(window *glfw.Window, program uint32, vao uint32) {
	//clear screen
	gl.ClearColor(0, 0, 1, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)




	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)

	//Swap buffers and poll for events
	window.SwapBuffers()
	glfw.PollEvents()
}

func setupView(program uint32) {

	//use our opengl program
	gl.UseProgram(program)

	projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(WIDTH)/HEIGHT, 0.1, 10.0)
	projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	camera := mgl32.LookAtV(mgl32.Vec3{4, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	cameraUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])

	model := mgl32.Ident4()
	modelUniform := gl.GetUniformLocation(program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])


}

func genVAO(points []float32) uint32 {

	//create vertex buffer object to load data into
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	//Signal that rest of calls are about our buffer we made
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	//Load data into buffer (using static draw since we are not modifying point data)
	gl.BufferData(gl.ARRAY_BUFFER, 4 * len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

func createMainWindow() *glfw.Window {

	//Describe window we're about to make
	// 4x antialiasing
	glfw.WindowHint(glfw.Samples, 4)
	//Request openGl v4.1
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	//For mac compatibility
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	//create window
	window, err := glfw.CreateWindow(WIDTH, HEIGHT, "OpenGL-Tutorial.org", nil, nil)
	if err != nil {
		glfw.Terminate()
		panic(err)
	}

	//bring focus to window
	window.MakeContextCurrent()

	return window
}

//Only call after a window's context has been made current
func initOpenGL() uint32 {
	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	//Load in shaders
	vShader, err := compileShader(shaders.VertexShader, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	defer gl.DeleteShader(vShader)

	fShader, err := compileShader(shaders.FragmentShader, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}
	defer gl.DeleteShader(fShader)

	//Create opengl program
	program := gl.CreateProgram()
	gl.AttachShader(program, vShader)
	gl.AttachShader(program, fShader)
	gl.LinkProgram(program)

	gl.DetachShader(program, vShader)
	gl.DetachShader(program, fShader)

	return program
}


func compileShader(source string, shaderType uint32) (uint32, error) {
	 shader := gl.CreateShader(shaderType)

	 csources, free := gl.Strs(source)
	 gl.ShaderSource(shader, 1, csources, nil)
	 free()
	 gl.CompileShader(shader)

	 var status int32
	 gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	 if status == gl.FALSE {
				 var logLength int32
				 gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

				 log := strings.Repeat("\x00", int(logLength+1))
				 gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

				 return 0, fmt.Errorf("failed to compile %v: %v", source, log)
		 }

	 return shader, nil
}