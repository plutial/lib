package main

// A 2D Matrix represented as a 1D array
type Matrix4 [16]float32

// An identity matrix
// An identity matrix, when multiplied, returns what it was multiplied by
func Matrix4Identity() Matrix4 {
	return Matrix4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// Returns the value at (x, y) of the matrix
func (m *Matrix4) Get(x, y int) float32 {
	return m[y*4+x]
}

// Set the value at (x, y)
func (m *Matrix4) Set(value float32, x, y int) {
	m[y*4+x] = value
}

// Returns the output of the multiplication of two matrices
func Matrix4Multiply(a, b Matrix4) Matrix4 {
	m := Matrix4{}

	for y := range 4 {
		for x := range 4 {
			var value float32 = 0
			value += a.Get(0, y) * b.Get(x, 0)
			value += a.Get(1, y) * b.Get(x, 1)
			value += a.Get(2, y) * b.Get(x, 2)
			value += a.Get(3, y) * b.Get(x, 3)

			m.Set(value, x, y)
		}
	}

	return m
}

// Translate the matrix
func Matrix4Translate(m Matrix4, x, y, z float32) Matrix4 {
	// x, y, z are set at the bottom of the matrix
	translationMatrix := Matrix4{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	}

	return Matrix4Multiply(m, translationMatrix)
}

// Scale the matrix
func Matrix4Scale(m Matrix4, x, y, z float32) Matrix4 {
	// x, y, z are the scales
	scaleMatrix := Matrix4{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	}

	return Matrix4Multiply(m, scaleMatrix)
}

// Ortho
func Matrix4Ortho(left, right, bottom, top, near, far float32) Matrix4 {
	matrix := Matrix4Identity()
	matrix = Matrix4Scale(matrix, 2/(right-left), 2/(top-bottom), -2/(far-near))
	matrix.Set(-(right+left)/(right-left), 0, 3)
	matrix.Set(-(top+bottom)/(top-bottom), 1, 3)
	matrix.Set(-(far+near)/(far-near), 2, 3)

	return matrix
}
