package gl33

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Material struct {
	id uint32
}

var latestUseProgramId uint32

func (m *Material) ShouldUseMaterial() {
	if latestUseProgramId == m.id {
		return
	}
	gl.UseProgram(m.id)
}

func NewMaterial(vertex, fragment *Shader) (*Material, error) {
	material := &Material{}
	var err error
	material.id, err = newProgram(vertex, fragment)
	return material, err
}

func (m *Material) SetMatrix4ByName(name string, matrix mgl32.Mat4) {
	m.ShouldUseMaterial()
	location := gl.GetUniformLocation(m.id, gl.Str(name+"\x00"))
	gl.UniformMatrix4fv(location, 1, false, &matrix[0])
}

func (m *Material) SetInt1ByName(name string, integer int32) {
	m.ShouldUseMaterial()
	location := gl.GetUniformLocation(m.id, gl.Str(name+"\x00"))
	gl.Uniform1i(location, integer)
}

func newProgram(vertex, fragment *Shader) (uint32, error) {
	program := gl.CreateProgram()

	gl.AttachShader(program, vertex.id)
	gl.AttachShader(program, fragment.id)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertex.id)
	gl.DeleteShader(fragment.id)

	gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))

	return program, nil
}
