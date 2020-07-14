package ske

import "github.com/veandco/go-sdl2/sdl"

// drawable is a Resource interface, that can be drawn to the screen
type Drawable interface {
	Draw(position *sdl.Rect)
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

func (t*Texture) Draw(position *sdl.Rect){
	t.
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

// an atlas is simply a texture, that we can pick out other textures from
type Atlas struct {
	Texture *Texture
	// store the texture atlas as an array of positional data for each cell
	Positions []sdl.Rect
	// the coordinate of the active texture
	Position Vec
	// the size of the atlas
	Size    Vec
}

// set the atlas position for the position in the atlas
func (a *Atlas) SetAtlasPosition(x, y int){
	a.Position = V2(float64(x), float64(y))
}

// splice out a specific texture from the atlas.
// w, h refers to the number of rows & columns
func (a*Atlas) Splice(w, h int){
	a.Size = V2(float64(w), float64(h))

	// calculate the size of each cell
	cellWidth := int(a.Texture.Size.X)/w
	cellHeight := int(a.Texture.Size.Y)/h

	// setup the positional data
	for y:=0;y<w;y++{
		for x:=0;x<h;x++{
			a.Positions = append(a.Positions, sdl.Rect{
				X: int32(x*cellHeight),
				Y: int32(y*cellHeight),
				W: int32(cellWidth),
				H: int32(cellHeight),
			})
		}
	}
}

func (*Atlas) Type() uint8 {
	return ATLAS
}

func (a*Atlas) Draw(position *sdl.Rect){
	Screen.Renderer.Copy(a.Texture.Data, &a.Positions[int(a.Position.X + a.Position.Y*a.Size.Y)], position)
}