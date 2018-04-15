package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kaneta1992/gl/renderer/gl33"
)

const windowWidth = 1200
const windowHeight = 900

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	win, err := gl33.NewWindow(windowWidth, windowHeight, "ora")

	vs, err := gl33.NewVertexShaderFromFile("shader/vertex.glsl")
	if err != nil {
		panic(err)
	}
	fs, err := gl33.NewFragmentShaderFromFile("shader/fragment.glsl")
	if err != nil {
		panic(err)
	}
	material, err := gl33.NewMaterial(vs, fs)
	if err != nil {
		panic(err)
	}

	projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(windowWidth)/windowHeight, 0.1, 10.0)
	material.SetMatrix4ByName("projection", projection)

	camera := mgl32.LookAtV(mgl32.Vec3{0, 0, 3.0}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	material.SetMatrix4ByName("camera", camera)

	material.SetInt1ByName("tex", 0)
	material.SetInt1ByName("cherryTex", 1)
	material.SetInt1ByName("kiraTex", 2)
	material.SetInt1ByName("mask1", 3)
	material.SetInt1ByName("mask3", 4)
	material.SetInt1ByName("sunTex", 5)
	material.SetInt1ByName("lensTex", 6)

	// Load the texture
	card, err := gl33.NewTexture("card.png")
	if err != nil {
		log.Fatalln(err)
	}
	cherry, err := gl33.NewTexture("cherry2.png")
	if err != nil {
		log.Fatalln(err)
	}
	kira, err := gl33.NewTexture("kira.png")
	if err != nil {
		log.Fatalln(err)
	}
	mask1, err := gl33.NewTexture("mask2.png")
	if err != nil {
		log.Fatalln(err)
	}
	mask3, err := gl33.NewTexture("mask4.png")
	if err != nil {
		log.Fatalln(err)
	}
	sun, err := gl33.NewTexture("sun.png")
	if err != nil {
		log.Fatalln(err)
	}
	lens, err := gl33.NewTexture("lens.png")
	if err != nil {
		log.Fatalln(err)
	}

	mesh, err := gl33.NewMesh(cubeVertices, material)

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0.2, 0.2, 0.2, 1.0)

	t := float64(0.0)
	win.GameLoop(func(dt float64) {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Update
		t += dt
		material.SetFloatByName("time", float32(t))

		// Render
		scale := mgl32.Scale3D(0.75, 0.75, 1.0)
		model := mgl32.Translate3D(0, 0, 0).Mul4(scale)
		material.SetMatrix4ByName("model", model)
		card.Set(0)
		cherry.Set(1)
		kira.Set(2)
		mask1.Set(3)
		mask3.Set(4)
		sun.Set(5)
		lens.Set(6)

		mesh.Draw()
	})
	win.Delete()
	mesh.Delete()
	material.Delete()
	card.Delete()
}

var cubeVertices = []float32{
	//  X, Y, Z, U, V
	// Front
	-1.0, -1.4, 0.0, 0.0, 1.0,
	1.0, -1.4, 0.0, 1.0, 1.0,
	-1.0, 1.4, 0.0, 0.0, 0.0,
	1.0, -1.4, 0.0, 1.0, 1.0,
	1.0, 1.4, 0.0, 1.0, 0.0,
	-1.0, 1.4, 0.0, 0.0, 0.0,
}
