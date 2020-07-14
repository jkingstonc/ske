package ske

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