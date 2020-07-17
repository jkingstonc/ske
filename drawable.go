package ske

import (
	"github.com/veandco/go-sdl2/sdl"
)

// TODO we need to call .Draw(position) in screen.go when we draw a drawable

// drawable is a Resource interface, that can be drawn to the screen
type Drawable interface {
	Draw(position *sdl.Rect)
	Size() Vec
}

// image resource
type Texture struct {
	// sdl texture
	Data *sdl.Texture
	// size of the texture in pixels
	TextureSize Vec
}

func (*Texture) Type() uint8{
	return TEXTURE
}

func (t*Texture) Draw(position *sdl.Rect){
	Screen.Renderer.CopyEx(t.Data, nil, position, 0, nil, sdl.FLIP_NONE)
}

func (t*Texture) Size() Vec{
	return t.TextureSize
}

// sprite resource
type Sprite struct {
	Animation Animation
	// scale relative to the pixel size
	Scale     Vec
	Rot		  float64
}

func (*Sprite) Type() uint8{
	return SPRITE
}

// as a sprite has an animation, we draw the correct animation
func (s*Sprite) Draw(position *sdl.Rect){
	s.Animation.Atlas.Draw(position)
}

func (s*Sprite) Size() Vec{
	return s.Animation.Atlas.Size()
}

// an atlas is simply a texture, that we can pick out other textures from
type Atlas struct {
	Texture *Texture
	// store the texture atlas as an array of positional data for each cell
	Positions []sdl.Rect
	// the coordinate of the active texture
	Position Vec
	// the size of the atlas
	GridSize    Vec
}

func NewAtlas(texture *Texture, w, h int) *Atlas{
	atlas := &Atlas{
		Texture:   texture,
	}
	atlas.splice(w, h)
	return atlas
}

// set the atlas position for the texture to be drawn in the atlas
func (a *Atlas) SetAtlasPosition(pos uint){
	// calculate the y value by dividing by the width
	y := int(pos)/ int(a.Size().X)
	// calculate the x value by getting the remainder
	x := int(pos) % int(a.Size().X)
	a.Position = V2(float64(x), float64(y))
}

// splice out the atlas into the correct texture sizes
// w, h refers to the number of rows & columns
func (a*Atlas) splice(w, h int){
	a.GridSize = V2(float64(w), float64(h))

	// calculate the size of each cell
	cellWidth := int(a.Texture.TextureSize.X)/w
	cellHeight := int(a.Texture.TextureSize.Y)/h

	// setup the positional data
	for y:=0;y<h;y++{
		for x:=0;x<w;x++{
			a.Positions = append(a.Positions, sdl.Rect{
				X: int32(x*cellWidth),
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

// draw the correct texture in the atlas
func (a*Atlas) Draw(position *sdl.Rect){
	Screen.Renderer.Copy(a.Texture.Data, &a.Positions[int(a.Position.X + a.Position.Y*a.GridSize.Y)], position)
}

// return the size of the texture's in the atlas NOT the atlas as a whole
func (a*Atlas) Size() Vec{
	return V2((a.Texture.TextureSize.X)/a.GridSize.X, (a.Texture.TextureSize.Y)/a.GridSize.Y)
}