package gl33

import "github.com/go-gl/gl/v3.3-core/gl"

type Mesh struct {
	vboId    uint32
	vaoId    uint32
	material *Material
}

func NewMesh(vertices []float32, material *Material) (*Mesh, error) {
	mesh := &Mesh{material: material}
	// Configure the vertex data
	gl.GenVertexArrays(1, &mesh.vaoId)
	gl.BindVertexArray(mesh.vaoId)

	gl.GenBuffers(1, &mesh.vboId)
	gl.BindBuffer(gl.ARRAY_BUFFER, mesh.vboId)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	vertAttrib := uint32(gl.GetAttribLocation(material.id, gl.Str("position\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(material.id, gl.Str("uv\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

	gl.BindVertexArray(0)

	return mesh, nil
}

func (m *Mesh) Draw() {
	m.material.ShouldUseMaterial()
	gl.BindVertexArray(m.vaoId)
	gl.DrawArrays(gl.TRIANGLES, 0, 2*3)
	gl.BindVertexArray(0)
}

func (m *Mesh) Delete() {
	gl.DeleteBuffers(1, &m.vboId)
	gl.DeleteVertexArrays(1, &m.vaoId)
}
