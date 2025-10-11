package main

import (
	"image/color"
	"log"

	// GLFW
	"github.com/go-gl/glfw/v3.3/glfw"

	// OpenGL
	"github.com/go-gl/gl/v3.3-core/gl"
)

// Global window
var window *glfw.Window

// Global frame counter
var frameCount int

func Init() {
	// Initialize GLFW
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	var err error
	window, err = glfw.CreateWindow(800, 450, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	// Frame buffer callback
	glfw.GetCurrentContext().SetFramebufferSizeCallback(FramebufferSizeCallback)

	// Key callback
	glfw.GetCurrentContext().SetKeyCallback(KeyCallback)

	// Initialize OpenGL
	if err := gl.Init(); err != nil {
		log.Fatalln(err)
	}
}

func ColorWindow(color color.RGBA) {
	// Convert it to decimals before applying
	gl.ClearColor(float32(color.R)/255, float32(color.G)/255, float32(color.B)/255, float32(color.A)/255)
}

func ClearWindow() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func RenderWindow() {
	window.SwapBuffers()

	// Window updates
	UpdateWindow()
}

func UpdateWindow() {
	// Update the frame count
	frameCount++

	glfw.PollEvents()
}

func WindowShouldClose() bool {
	return window.ShouldClose()
}

func FramebufferSizeCallback(window *glfw.Window, width, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}
