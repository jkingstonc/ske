package ske

const (
	TEXTURE = 0x0
	SPRITE  = 0x1
	ATLAS   = 0x2
	AUDIO   = 0x3
	TILEMAP = 0x4
	FONT    = 0x5
	TEXT    = 0x6
)

type Resource interface {
	Type() uint8
}

// tile resource (used in tile-maps)
type Tile struct {
	Pos     Vec
	Texture *Texture
}

// tile-map resource.
// a tile-map is essentially an array of tiles, the tilemap is used for easy texture loading.
// the programmer must then create the entities that go along with the tilemap
type TileMap struct {
	Tiles []*Tile
}

func (*TileMap) Type() uint8{
	return TILEMAP
}