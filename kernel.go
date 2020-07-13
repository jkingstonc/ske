package ske

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	TEXTURE   = 0x0
	AUDIO   = 0x1
	TILEMAP = 0x2
)

type Resource interface {
	Type() uint8
}

// image resource
type Texture struct {
	// sdl image here
}

func (*Texture) Type() uint8{
	return TEXTURE
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
	Texture Texture
}

// tile-map resource
type TileMap struct {
	Tiles [][]Tile
}

func (*TileMap) Type() uint8{
	return TILEMAP
}


type FileManager struct {
	LoadedFiles map[string]Resource
}

// load a file into the game kernel
func (f*FileManager) Load(paths... string){
	for _, path := range paths{
		// check what the extension of the file is
		switch filepath.Ext(filepath.Base(path)){
		case ".png":
			fallthrough
		case ".jpg":
			_, err := os.Open(AssetsRoot +path)
			Assert(err==nil, "cannot open file")
		}
	}
}

// load a file into the game kernel
func WriteRaw(path string, data []byte){
	file, err := os.Create(AssetsRoot +path)
	Assert(err==nil, "cannot open file")
	file.Write(data)
}

// load a file into the game kernel
func ReadRaw(path string) []byte{
	bytes, err := ioutil.ReadFile(AssetsRoot +path)
	Assert(err==nil, "cannot open file")
	return bytes
}

// load a file into the game kernel
func (f*FileManager) File(name string) Resource{
	return f.LoadedFiles[name]
}