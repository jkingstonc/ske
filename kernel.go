package ske

import (
	"bytes"
	"encoding/gob"
	"github.com/boltdb/bolt"
	"github.com/veandco/go-sdl2/mix"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FileManager struct {
	LoadedFiles map[string]Resource
	DB *bolt.DB
}

func (f*FileManager) Setup(){
	// open the db
	db, err := bolt.Open(Engine.options.AssetsRoot+"dat", 0640, nil)
	Assert(err==nil, "couldn't open bolt database")

	// create a bucket
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("data.store"))
		return err
	})

	f.DB = db
}

func (f*FileManager) Close(){
	f.DB.Close()
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


// https://www.opsdash.com/blog/persistent-key-value-store-golang.html

func (f*FileManager) WriteKeyValue(key string, value interface{}){
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(value)
	Assert(err==nil, "could not encode value using gob")
	f.DB.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte("data.store")).Put([]byte(key), buf.Bytes())
	})
}

func (f*FileManager) ReadKeyValue(key string, value interface{}) {
	err := f.DB.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte("data.store")).Cursor()
		if k, v := c.Seek([]byte(key)); k == nil || string(k) != key {
			return nil
		} else if value == nil {
			return nil
		} else {
			d := gob.NewDecoder(bytes.NewReader(v))
			err := d.Decode(value)
			Assert(err==nil, "could not decode value")
			return err
		}
	})
	Assert(err==nil, "could not view db")
}