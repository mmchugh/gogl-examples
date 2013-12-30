package main

import (
	gl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	glhelpers "github.com/mmchugh/glhelpers"
	mathgl "github.com/Jragonmiris/mathgl"
)

func main() {
	if !glfw.Init() {
		panic("Can't init glfw!")
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(400, 400, "go-gl 2", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	gl.Init()

	vertices := []float32{0.0, 0.5, 0.5, -0.5, -0.5, -0.5}
	vertex_buffer := gl.GenBuffer()
	vertex_buffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, vertices, gl.STATIC_DRAW)

	var program gl.Program = glhelpers.CreateProgramFromPaths("2d.vshader", "white.fshader")
	program.Use()

	mvp_uniform := program.GetUniformLocation("MVP")
	projection_matrix := mathgl.Perspective(45.0, 1.0, 1.0, 500.0)
	view_matrix := mathgl.LookAt(
		0.0, 0.0, -5.0,
		0.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
	)
	model_matrix := mathgl.Ident4f()
	rotation := float32(0.0)

	position := program.GetAttribLocation("position")
	position.EnableArray()
	position.AttribPointer(2, gl.FLOAT, false, 0, nil)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		rotation += 1.0
		model_matrix = mathgl.Ident4f().Mul4(mathgl.HomogRotate3DZ(rotation))
		mvp := projection_matrix.Mul4(view_matrix).Mul4(model_matrix)
		mvp_uniform.UniformMatrix4fv(false, mvp)

		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
