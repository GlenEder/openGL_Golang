package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"runtime"
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



}