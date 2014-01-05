package main

import (
	"time"
	"math"
	"unsafe"
	gl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	glhelpers "github.com/mmchugh/glhelpers"
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

	vertices := []glhelpers.Vec2{{0.0, 0.5}, {0.5, -0.5}, {-0.5, -0.5}}
	vertex_buffer := gl.GenBuffer()
	vertex_buffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, int(unsafe.Sizeof(vertices)), vertices, gl.STATIC_DRAW)

	var program gl.Program = glhelpers.CreateProgramFromPaths("2d.vshader", "white.fshader")
	program.Use()

	mvp_uniform := program.GetUniformLocation("MVP")
	projection_matrix := glhelpers.Perspective(45.0, 1.0, 1.0, 500.0)
	view_matrix := glhelpers.LookAt(
		0.0, 0.0, -5.0,
		0.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
	)
	model_matrix := glhelpers.Ident4()
	rotation := float32(0.0)
	speed := float32(math.Pi)
	position := program.GetAttribLocation("position")
	position.EnableArray()
	position.AttribPointer(2, gl.FLOAT, false, 0, nil)
	last_time := time.Now()

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		current_time := time.Now()
		rotation += speed * float32(current_time.Sub(last_time).Seconds())
		last_time = current_time
		model_matrix = glhelpers.Ident4().RotateZ(rotation)
		mvp := projection_matrix.Mult(view_matrix).Mult(model_matrix)
		mvp_uniform.UniformMatrix4fv(false, mvp)

		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
