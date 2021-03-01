package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"runtime"
)

const (
	//window width and height
	WIDTH = 480
	HEIGHT = 480
)

func init() {
	//lock onto main thread for glfw and gl calls
	runtime.LockOSThread()

	//Initialize glfw
	if err := glfw.Init(); err != nil {
		panic(err)
	}
}

func main() {

	window := createMainWindow()
	defer glfw.Terminate()

	//Ensure that we can read input
	window.SetInputMode(glfw.StickyKeysMode, glfw.True)

	//render loop
	for !window.ShouldClose() && window.GetKey(glfw.KeyEscape) != glfw.Press {
		draw(window)
	}

}

//Render method
func draw(window *glfw.Window) {


	//Swap buffers and poll for events
	window.SwapBuffers()
	glfw.PollEvents()
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