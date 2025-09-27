package main

import (
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	"strings"
)

type Shader struct {
	program       uint32
	vao, vbo, ebo uint32
}

func checkCompileErrors(shader uint32, shaderType string) error {
	// Check that the shader has compiled correctly
	var status int32 = 1

	if shaderType != "PROGRAM" {
		gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	} else {
		gl.GetShaderiv(shader, gl.LINK_STATUS, &status)
	}

	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return fmt.Errorf("failed to compile %v: %v", shaderType, log)
	}

	return nil
}

func compileShader(source string, shaderType uint32) uint32 {
	// Create and compiler the shader of the given type
	shader := gl.CreateShader(shaderType)
	cSource, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, cSource, nil)
	free()
	gl.CompileShader(shader)

	var shaderTypeString string
	switch shaderType {
	case gl.VERTEX_SHADER:
		shaderTypeString = "VERTEX"
	case gl.FRAGMENT_SHADER:
		shaderTypeString = "FRAGMENT"
	default:
		shaderTypeString = "UNKNOWN"
	}

	if err := checkCompileErrors(shader, shaderTypeString); err != nil {
		panic(err)
	}

	return shader
}

func newProgram(vertexShaderSource, fragmentShaderSource string) uint32 {
	vertexShader := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	fragmentShader := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)

	program := gl.CreateProgram()

	// Attach both shaders
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	// Verify that the program has successfully been created
	if err := checkCompileErrors(program, "PROGRAM"); err != nil {
		panic(err)
	}

	// The shaders are no longer required
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program
}

func NewShader() Shader {
	shader := Shader{}

	// GLSL Code
	vertexShaderSource := `#version 330 core
	layout (location = 0) in vec3 aPos;
	layout (location = 1) in vec3 aColor;
	layout (location = 2) in vec2 aTexCoord;

	out vec3 ourColor;
	out vec2 TexCoord;

	void main()
	{
		gl_Position = vec4(aPos, 1.0);
		ourColor = aColor;
		TexCoord = vec2(aTexCoord.x, aTexCoord.y);
	}` + "\x00"

	fragmentShaderSource := `#version 330 core
	out vec4 FragColor;

	in vec3 ourColor;
	in vec2 TexCoord;

	// texture sampler
	uniform sampler2D texture1;

	void main()
	{
		FragColor = texture(texture1, TexCoord);
	}` + "\x00"

	shader.program = newProgram(vertexShaderSource, fragmentShaderSource)

	// Positions, colors, and the texture co-ordinates
	vertices := []float32{
		// positions          // colors           // texture coords
		0.5, 0.5, 0.0, 1.0, 0.0, 0.0, 1.0, 1.0, // top right
		0.5, -0.5, 0.0, 0.0, 1.0, 0.0, 1.0, 0.0, // bottom right
		-0.5, -0.5, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, // bottom let
		-0.5, 0.5, 0.0, 1.0, 1.0, 0.0, 0.0, 1.0, // top let
	}
	indices := []uint32{
		0, 1, 3, // first triangle
		1, 2, 3, // second triangle
	}

	// VAO and VBO
	gl.GenVertexArrays(1, &shader.vao)
	gl.GenBuffers(1, &shader.vbo)
	gl.GenBuffers(1, &shader.ebo)

	gl.BindVertexArray(shader.vao)

	// sizeof(float) == 4
	sizeofFloat := 4
	gl.BindBuffer(gl.ARRAY_BUFFER, shader.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*sizeofFloat, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, shader.ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*sizeofFloat, gl.Ptr(indices), gl.STATIC_DRAW)

	// Position attributes
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 8*int32(sizeofFloat), 0)
	gl.EnableVertexAttribArray(0)

	// Color attributes
	gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, 8*int32(sizeofFloat), uintptr(3*sizeofFloat))
	gl.EnableVertexAttribArray(1)

	// Texture co-ordinate attributes
	gl.VertexAttribPointerWithOffset(2, 2, gl.FLOAT, false, 8*int32(sizeofFloat), uintptr(6*sizeofFloat))
	gl.EnableVertexAttribArray(2)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.BindVertexArray(0)

	return shader
}

func (shader *Shader) Delete() {
	gl.DeleteVertexArrays(1, &shader.vao)
	gl.DeleteBuffers(1, &shader.vbo)
	gl.DeleteBuffers(1, &shader.ebo)
	gl.DeleteProgram(shader.program)
}
