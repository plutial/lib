package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
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

	for !window.ShouldClose() {
		if glfw.GetCurrentContext().GetKey(glfw.KeyEscape) == glfw.Press {
			window.SetShouldClose(true)
		}

		// Clear the screen
		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// Draw the triangle
		gl.UseProgram(shader.program)
		gl.BindVertexArray(shader.vao)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
