package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	// Matrix
	glm "github.com/go-gl/mathgl/mgl32"

	"log"
	"runtime"
)

func main() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()

	window := Init()
	defer glfw.Terminate()

	// Shaders
	shader := NewShader()
	defer shader.Delete()

	texture1, err := NewTexture("wall.jpg")
	if err != nil {
		log.Fatal(err)
	}

	texture2, err := NewTexture("container.jpg")
	if err != nil {
		log.Fatal(err)
	}

	source := glm.Vec4{0, 0, 1600, 1600}
	destination1 := glm.Vec4{0, 0, 16, 16}
	destination2 := glm.Vec4{16, 0, 16, 16}

	for !window.ShouldClose() {
		if glfw.GetCurrentContext().GetKey(glfw.KeyEscape) == glfw.Press {
			window.SetShouldClose(true)
		}

		// Clear the screen
		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// Render the textures
		shader.Render(texture1, source, destination1, 0)
		shader.Render(texture2, source, destination2, 0)

		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
