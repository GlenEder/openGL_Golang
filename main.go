package main

import "runtime"

func init() {
	//lock onto main thread for glfw and gl calls
	runtime.LockOSThread()
}

func main() {



}