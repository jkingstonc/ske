package ske

import (
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
		// check what the extension of the file is
		switch filepath.Ext(filepath.Base(path)){
		case ".png":
			fallthrough
		case ".jpg":
			texture := Screen.LoadTexture(AssetsRoot +path)
			f.LoadedFiles[path]=texture
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