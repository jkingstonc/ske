package ske

import (
	"github.com/veandco/go-sdl2/mix"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FileManager struct {
	LoadedFiles map[string]Resource
}

// load a file into the game kernel
func (f*FileManager) Load(paths... string){
	for _, path := range paths{
		fullPath := Engine.options.AssetsRoot + path
		// check what the extension of the file is
		switch filepath.Ext(filepath.Base(path)) {
		case ".ogg":
			music, err := mix.LoadMUS(fullPath)
			Assert(err==nil, "could not load music file in the kernel")
			f.LoadedFiles[path] = &Audio{Music: music,}
		case ".ttf":
			f.LoadedFiles[path] = &FontResource{Font: nil, Path: path}
		case ".png":
			fallthrough
		case ".jpg":
			texture := Screen.LoadTexture(fullPath)
			f.LoadedFiles[path] = texture
		case ".tmx": // tilemaps
			break
		}
	}
}

// load a file into the game kernel
func WriteRaw(path string, data []byte){
	file, err := os.Create(Engine.options.AssetsRoot +path)
	Assert(err==nil, "cannot open file")
	file.Write(data)
}

// load a file into the game kernel
func ReadRaw(path string) []byte{
	bytes, err := ioutil.ReadFile(Engine.options.AssetsRoot +path)
	Assert(err==nil, "cannot open file")
	return bytes
}

// load a file into the game kernel
func (f*FileManager) File(name string) Resource{
	return f.LoadedFiles[name]
}