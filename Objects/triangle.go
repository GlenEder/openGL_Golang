package Objects

import "github.com/go-gl/gl/v4.1-core/gl"

//Points of triangle that are used for rendering our first triangle
var TrianglePoints = []float32{
	-1.0, -1.0, 0.0,
	1.0, -1.0, 0.0,
	0.0, 1.0, 0.0,
}

func DrawTriangle(vao uint32) {
	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
}