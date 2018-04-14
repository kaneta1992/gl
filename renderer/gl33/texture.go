package gl33

import (
	"fmt"
	"image"
	"image/draw"
	"os"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Texture struct {
	id   uint32
	unit uint32
}

func NewTexture(path string) (*Texture, error) {
	texture, err := newTexture(path)
	if err != nil {
		panic(err)
	}
	return &Texture{id: texture, unit: 99999}, nil
}

var glTextureUnits = []uint32{
	gl.TEXTURE0,
	gl.TEXTURE1,
	gl.TEXTURE2,
	gl.TEXTURE3,
	gl.TEXTURE4,
	gl.TEXTURE5,
	gl.TEXTURE6,
	gl.TEXTURE7,
	gl.TEXTURE8,
}

var activeTextureTable = map[uint32]*Texture{}

func (t *Texture) Set(unit uint32) {
	if activeTextureTable[unit] == t {
		return
	}
	// t.unit = unit
	activeTextureTable[unit] = t
	gl.ActiveTexture(glTextureUnits[unit])
	gl.BindTexture(gl.TEXTURE_2D, t.id)
}

func (t *Texture) Delete() {
	gl.DeleteTextures(1, &t.id)
}

func newTexture(file string) (uint32, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return 0, fmt.Errorf("texture %q not found on disk: %v", file, err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return 0, err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return 0, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture, nil
}
