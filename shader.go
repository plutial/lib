package main

import (
	"fmt"
	"os"
	"strings"

	// OpenGL
	"github.com/go-gl/gl/v3.3-core/gl"

	// Matrix
	glm "github.com/go-gl/mathgl/mgl32"
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
	vertexShaderSource, err := os.ReadFile("shader.vs")
	if err != nil {
		panic(err)
	}

	// Add a terminate character at the end
	vertexShaderSource = append(vertexShaderSource, '\x00')

	fragmentShaderSource, err := os.ReadFile("shader.fs")
	if err != nil {
		panic(err)
	}

	// Add a terminate character at the end
	fragmentShaderSource = append(fragmentShaderSource, '\x00')

	shader.program = newProgram(string(vertexShaderSource), string(fragmentShaderSource))

	// Positions, colors, and the texture co-ordinates
	vertices := []float32{
		// Positions  Texture coords
		1.0, 1.0, 1.0, 1.0, // top right
		1.0, 0.0, 1.0, 0.0, // bottom right
		0.0, 0.0, 0.0, 0.0, // bottom let
		0.0, 1.0, 0.0, 1.0, // top let
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
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(0, 4, gl.FLOAT, false, 4*int32(sizeofFloat), 0)

	// Color attributes
	/*gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, 8*int32(sizeofFloat), uintptr(3*sizeofFloat))
	gl.EnableVertexAttribArray(1)

	// Texture co-ordinate attributes
	gl.VertexAttribPointerWithOffset(2, 2, gl.FLOAT, false, 8*int32(sizeofFloat), uintptr(6*sizeofFloat))
	gl.EnableVertexAttribArray(2)*/

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.BindVertexArray(0)

	// Projection of the window
	projection := Matrix4Ortho(0, 800, 450, 0, -1, 1)
	shader.SetMatrix4(&projection[0], "projection")

	return shader
}

func (shader *Shader) SetInteger(value int32, name string) {
	gl.Uniform1i(gl.GetUniformLocation(shader.program, gl.Str(name+"\x00")), value)
}

func (shader *Shader) SetVector2(vector *float32, name string) {
	gl.Uniform2fv(gl.GetUniformLocation(shader.program, gl.Str(name+"\x00")), 1, vector)
}

func (shader *Shader) SetVector4(vector *float32, name string) {
	gl.Uniform4fv(gl.GetUniformLocation(shader.program, gl.Str(name+"\x00")), 1, vector)
}

func (shader *Shader) SetMatrix4(matrix *float32, name string) {
	gl.UniformMatrix4fv(gl.GetUniformLocation(shader.program, gl.Str(name+"\x00")), 1, false, matrix)
}

var abc float32

func (shader *Shader) Render(x, y float32) {
	// Source rectangle
	abc += 0.01
	vector := glm.Vec4{abc, abc, abc, abc}
	shader.SetVector4(&vector[0], "uvModel")

	// Destination rectangle
	projection := glm.Ortho(0, 800, 450, 0, -1, 1)
	shader.SetMatrix4(&projection[0], "projection")

	// Order of transformations are important
	model := glm.Ident4()

	// Position
	model = model.Mul4(glm.Translate3D(x, y, 0))

	// Rotation
	var angle float32 = 0.0
	model = model.Mul4(glm.Translate3D(0.5*16, 0.5*16, 0.0))
	model = model.Mul4(glm.HomogRotate3DZ(angle))
	model = model.Mul4(glm.Translate3D(-0.5*16, -0.5*16, 0.0))

	// Size
	model = model.Mul4(glm.Scale3D(16, 16, 1))

	fmt.Println(model)

	// Apply the matrix
	shader.SetMatrix4(&model[0], "model")
}

func (shader *Shader) Delete() {
	gl.DeleteVertexArrays(1, &shader.vao)
	gl.DeleteBuffers(1, &shader.vbo)
	gl.DeleteBuffers(1, &shader.ebo)
	gl.DeleteProgram(shader.program)
}
