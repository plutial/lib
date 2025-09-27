package main

import (
	"fmt"
	"image"
	"image/draw"
	"os"

	// Support both image types
	_ "image/jpeg"
	_ "image/png"

	// OpenGL
	"github.com/go-gl/gl/v3.3-core/gl"
)

func NewTexture(path string) (uint32, error) {
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	// Texture wrapping paramaters
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

	// Texture filtering parameters: used for filters such as blurring
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST_MIPMAP_NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	// Open the image file
	imageFile, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("texture %q not found on disk: %v", path, err)
	}
	defer imageFile.Close()

	// Load the contents of the file
	contents, _, err := image.Decode(imageFile)
	if err != nil {
		panic(err)
	}

	// Create a RGBA image
	rgbaImage := image.NewRGBA(contents.Bounds())
	if rgbaImage.Stride != rgbaImage.Rect.Size().X*4 {
		return 0, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgbaImage, rgbaImage.Bounds(), contents, image.Point{0, 0}, draw.Src)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgbaImage.Rect.Size().X),
		int32(rgbaImage.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgbaImage.Pix))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	return texture, nil
}
