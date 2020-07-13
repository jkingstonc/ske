package ske

import "os"

type FileManager struct {
	LoadedFiles map[string]*os.File
}

// load a file into the game kernel
func (f*FileManager) Load(path string){
	file, err := os.Open(AssetsRoot+path)
	Assert(err==nil, "cannot open file")
	f.LoadedFiles[path] = file
}

// load a file into the game kernel
func (f*FileManager) File(name string) *os.File{
	return f.LoadedFiles[name]
}