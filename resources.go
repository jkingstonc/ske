package ske

import "github.com/veandco/go-sdl2/sdl"

const (
	TEXTURE = 0x0
	SPRITE  = 0x1
	ATLAS   = 0x2
	AUDIO   = 0x3
	TILEMAP = 0x4
)

type Resource interface {
	Type() uint8
}

// image resource
type Texture struct {
	// sdl texture
	Data *sdl.Texture
	// size of the texture in pixels
	Size Vec
}

func (*Texture) Type() uint8{
	return TEXTURE
}

// sprite resource (essentially an animated texture)
type Sprite struct {
	Animation Animation
	// scale relative to the pixel size
	Scale     Vec
}

func (*Sprite) Type() uint8{
	return SPRITE
}

// an animation has a series of textures
type Animation struct {
	Textures []*Texture
	Speed    float64
}

// an atlas is simply a texture, that we can pick out other textures from
type Atlas struct {
	Texture *Texture
	Size    Vec
}

// splice out a specific texture from the atlas
func (*Atlas) Splice(p1, p2 Vec) *Texture{
	return nil
}

func (*Atlas) Type() uint8 {
	return ATLAS
}

// image resource
type Audio struct {
}

func (*Audio) Type() uint8{
	return AUDIO
}

// tile resource (used in tile-maps)
type Tile struct {
	Pos     Vec
	Texture *Texture
}

// tile-map resource
type TileMap struct {
	Tiles [][]Tile
}

func (*TileMap) Type() uint8{
	return TILEMAP
}