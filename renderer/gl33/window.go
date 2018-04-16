package gl33

import (
	"fmt"
	_ "image/png"
	"log"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Window struct {
	window *glfw.Window
}

func initGraphics() {
	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)
}

func NewWindow(w, h int, title string) (*Window, error) {
	win := &Window{}
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	if runtime.GOOS == "darwin" {
		glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	}
	var err error
	win.window, err = glfw.CreateWindow(w, h, title, nil, nil)
	if err != nil {
		log.Println(err)
	}
	win.window.MakeContextCurrent()
	initGraphics()
	return win, nil
}

func (w *Window) GameLoop(proc func(float64)) {
	previousTime := glfw.GetTime()
	for !w.window.ShouldClose() {
		time := glfw.GetTime()
		elapsed := time - previousTime
		previousTime = time
		proc(elapsed)
		// Maintenance
		w.window.SwapBuffers()
		glfw.PollEvents()
	}
}

func (w *Window) Delete() {
	glfw.Terminate()
}
