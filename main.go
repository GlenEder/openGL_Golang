package main

import (
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gl/v4.1-core/gl"
	"runtime"
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
	initOpenGL()

	//Ensure that we can read input
	window.SetInputMode(glfw.StickyKeysMode, glfw.True)

	//load up our triangle
	triangle := genVBO(trianglePoints)

	//render loop
	for !window.ShouldClose() && window.GetKey(glfw.KeyEscape) != glfw.Press {
		draw(window, triangle)
	}

}

//Render method
func draw(window *glfw.Window, vbo uint32) {
	//clear screen
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.EnableVertexAttribArray(0)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	gl.DisableVertexAttribArray(0)

	//Swap buffers and poll for events
	window.SwapBuffers()
	glfw.PollEvents()
}

func genVBO(points []float32) uint32 {

	//create vertex buffer object to load data into
	var vbo uint32
	gl.GenBuffers(1, &vbo)

	//Signal that rest of calls are about our buffer we made
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	//Load data into buffer (using static draw since we are not modifying point data)
	gl.BufferData(gl.ARRAY_BUFFER, len(points), gl.Ptr(points), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vbo
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
func initOpenGL() {
	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)
}