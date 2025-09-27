package main

import (
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	"strings"
)

type Shader struct {
	program  uint32
	vao, vbo uint32
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
    void main()
    {
       gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);
    }`

	fragmentShaderSource := `#version 330 core
    out vec4 FragColor;
    void main()
    {
       FragColor = vec4(1.0f, 0.5f, 0.2f, 1.0f);
    }`

	shader.program = newProgram(vertexShaderSource, fragmentShaderSource)

	// Positions of the triangle
	vertices := []float32{
		-0.5, -0.5, 0.0, // let
		0.5, -0.5, 0.0, // right
		0.0, 0.5, 0.0, // top
	}

	// VAO and VBO
	gl.GenVertexArrays(1, &shader.vao)
	gl.GenBuffers(1, &shader.vbo)

	gl.BindVertexArray(shader.vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, shader.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// sizeof(float) == 4
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(0)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.BindVertexArray(0)

	return shader
}

func (shader *Shader) Delete() {
	gl.DeleteProgram(shader.program)
}
