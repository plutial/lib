package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
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

	for !window.ShouldClose() {
		if glfw.GetCurrentContext().GetKey(glfw.KeyEscape) == glfw.Press {
			window.SetShouldClose(true)
		}

		// Clear the screen
		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		shader.SetInteger(0, "image")
		shader.Render(0, 0)

		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, texture1)

		// Draw the triangle
		gl.UseProgram(shader.program)
		gl.BindVertexArray(shader.vao)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
		gl.BindVertexArray(0)

		shader.Render(16, 0)
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, texture2)

		// Draw the triangle
		gl.UseProgram(shader.program)
		gl.BindVertexArray(shader.vao)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
		gl.BindVertexArray(0)

		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
