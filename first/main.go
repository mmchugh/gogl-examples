package main

import (
    glfw "github.com/go-gl/glfw3"
    gl "github.com/go-gl/gl"
	glhelpers "github.com/mmchugh/glhelpers"
)

func main() {
    if !glfw.Init() {
        panic("Can't init glfw!")
    }
    defer glfw.Terminate()

    window, err := glfw.CreateWindow(400, 400, "go-gl 1", nil, nil)
    if err != nil {
        panic(err)
    }

	window.MakeContextCurrent()

    gl.Init()

	vertices := []float32 { 0.0, 0.5, 0.5, -0.5, -0.5, -0.5 }
	vertex_buffer := gl.GenBuffer()
	vertex_buffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices) * 4, vertices, gl.STATIC_DRAW)

    var program gl.Program = glhelpers.CreateProgramFromPaths("2d.vshader", "white.fshader")
	program.Use()
	position := program.GetAttribLocation("position")
	position.EnableArray()
	position.AttribPointer(2, gl.FLOAT, false, 0, nil)

    for !window.ShouldClose() {
        gl.DrawArrays(gl.TRIANGLES, 0, 3)
        window.SwapBuffers()
        glfw.PollEvents()
    }
}
