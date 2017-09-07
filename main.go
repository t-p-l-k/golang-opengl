package main

import (
	"flag"

	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
	var bender = flag.Bool("bender", false, "Render bender.")
	flag.Parse()
	var fragShader = "default.frag"
	fmt.Printf("%+v", *bender)
	if *bender {
		fragShader = "bender.frag"
	}
	config := ReadConfig()
	window := InitOpenGL(config)
	defer glfw.Terminate()

	// Configure global gl settings
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.ALPHA)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)

	// Configure vertex and fragment shaders
	program, err := NewProgram(
		ImportShader("default.vert"),
		ImportShader(fragShader),
	)
	if err != nil {
		panic(err)
	}

	// Generate vertex array
	vertexArray := MakeVertexArray(square)

	previousTime := float32(glfw.GetTime())

	// Main loop
	for !window.ShouldClose() {
		// Set uniforms

		w, h := window.GetSize()
		ru := gl.GetUniformLocation(program, gl.Str("iResolution\x00"))
		gl.Uniform2f(ru, float32(w), float32(h))

		x, y := window.GetCursorPos()
		l := window.GetMouseButton(glfw.MouseButton1)
		r := window.GetMouseButton(glfw.MouseButton2)
		mu := gl.GetUniformLocation(program, gl.Str("iMouse\x00"))
		gl.Uniform4f(mu, float32(x), float32(y), float32(l), float32(r))

		tu := gl.GetUniformLocation(program, gl.Str("iTime\x00"))
		gl.Uniform1f(tu, float32(glfw.GetTime()))

		draw(vertexArray, window, program, &previousTime)
	}
}

func draw(
	vertexArray uint32,
	window *glfw.Window,
	program uint32,
	previousTime *float32,
) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	// Update
	time := float32(glfw.GetTime())
	elapsed := time - *previousTime
	*previousTime = time
	// FIXME[Dmitry Teplov] Not using for now. Delete?
	_ = elapsed

	// Render
	gl.UseProgram(program)

	gl.BindVertexArray(vertexArray)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(square)))

	// Maintenance
	glfw.PollEvents()
	window.SwapBuffers()
}

var square = []float32{
	-1, 1, 0,
	-1, -1, 0,
	1, -1, 0,

	-1, 1, 0,
	1, 1, 0,
	1, -1, 0,
}
