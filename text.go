package ske

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"reflect"
)

type FontResource struct {
	Font *ttf.Font
	Path string
}

func (f*FontResource) NewFontTexture(size int) {
	font, err := ttf.OpenFont(Engine.options.AssetsRoot+f.Path, size)
	Assert(err==nil, "cannot open font")
	f.Font = font
}

func (*FontResource) Type() uint8 {
	return FONT
}

// text is a drawable resource
type Text struct {
	Texture *Texture
}

func (t*Text) Draw(position *sdl.Rect){
	t.Texture.Draw(position)
}


func (*Text) Type() uint8 {
	return TEXTURE
}


// attach this to game objects to render text
type TextComponent struct {
	Component
	Text string
	Font string
	Size int
	Texture *Texture
}

func (t*TextComponent) OnLoad(){
	// we first create the font
	texture := makeText(t.Text, Loader.File(t.Font).(*FontResource), t.Size, sdl.Color{255,0,0,255})
	t.Texture = texture

	mesh := t.Entity.GetComponent(reflect.TypeOf(&MeshComponent{})).(*MeshComponent)
	mesh.Drawable = t.Texture
}
func (t*TextComponent) Update(){}

// Load a font, and return the rendered font as a texture
func makeText(text string, font *FontResource, size int, color sdl.Color) *Texture {
	// create the font with the correct size
	font.NewFontTexture(size)
	solidSurface, err := font.Font.RenderUTF8Solid(text, color)
	Assert(err==nil, "failed to render font to surface")
	solidTexture, err := Screen.Renderer.CreateTextureFromSurface(solidSurface)
	Assert(err==nil, "failed to create texture from surface")
	r := solidSurface.ClipRect
	v := V2(float64(r.W), float64(r.H))
	solidSurface.Free()
	return &Texture{solidTexture, v}
}