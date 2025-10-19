package main

import (
	"image/color"

	"github.com/go-gl/glfw/v3.3/glfw"

	"fmt"
	"math/rand"

	// Matrix
	glm "github.com/go-gl/mathgl/mgl32"

	"log"
	"runtime"
)

/*
#undef _GNU_SOURCE
*/
import "C"

func main() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()

	Init()
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

	var blank Texture

	source := glm.Vec4{0, 0, 1600, 1600}
	destination1 := glm.Vec4{0, 0, 16, 16}
	destination2 := glm.Vec4{16, 0, 16, 16}
	destination3 := glm.Vec4{32, 0, 16, 16}

	for !WindowShouldClose() {
		if IsKeyPressed(KeyA) {
			fmt.Println("p", rand.Int())
		}
		if IsKeyPressedRepeat(KeyA) {
			fmt.Println("t", rand.Int())
		}
		if IsKeyReleased(KeyA) {
			fmt.Println("r", rand.Int())
		}

		// Clear the screen
		color := color.RGBA{54, 78, 79, 255}
		ColorWindow(color)
		ClearWindow()

		// Render the textures
		shader.Render(texture1, source, destination1, 0)
		shader.Render(texture2, source, destination2, 0)
		shader.Render(blank, source, destination3, 0)

		RenderWindow()
	}
}
