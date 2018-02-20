package gl33

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Shader struct {
	id uint32
}

func NewVertexShader(source string) (*Shader, error) {
	vertexShader := &Shader{}
	var err error
	vertexShader.id, err = compileShader(source, gl.VERTEX_SHADER)
	return vertexShader, err
}

func NewFragmentShader(source string) (*Shader, error) {
	fragmentShader := &Shader{}
	var err error
	fragmentShader.id, err = compileShader(source, gl.FRAGMENT_SHADER)
	return fragmentShader, err
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
