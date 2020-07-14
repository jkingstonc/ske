package ske

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Camera struct {
	Pos  Vec
	Zoom float64
}

//// SDL implementation of a renderer
type SDLScreen struct {
	Cam      Camera
	Window   *sdl.Window
	Renderer *sdl.Renderer
	ZBuf     ZBuffer
}

var (
	Width    int32 = 800
	Height   int32 = 600
)

const (
	UNIT_SIZE float64 = 50

	NOFILL = 0x1 << 0
	FILL   = 0x1 << 1

	FLIP = 0x1 << 0

	// max of 100 drawables per layer
	MAX_LAYER_DRAWS = 100

	D_RECT    = 0x0
	D_LINE    = 0x1
	D_TEXTURE = 0x2
	D_TEXT    = 0x3
)

type Layer struct {
	Drawables []Drawable
}

type Drawable struct {
	Type   uint8
	Data   interface{}
	Colour Vec
	Flags  uint8
	V1     Vec
	V2     Vec
	Angle  float64
}
// store sorted layers
type ZBuffer struct {
	Layers []Layer
}

// iteratively add layers to the z-buffer until it matches the desired z-value
func (s*SDLScreen) matchLayers(z int) {
	// add more layers until we get the desired z-depth
	l := len(s.ZBuf.Layers)
	for z >= l {
		s.ZBuf.Layers = append(s.ZBuf.Layers, Layer{})
		l = len(s.ZBuf.Layers)
	}
}

func (s*SDLScreen) WorldToScreen(v Vec) Vec {
	// first take the position
	pos1 := v.Add(s.Cam.Pos)
	// then place it in screen space
	pos1 = pos1.Mul(UNIT_SIZE)
	// zoom with camera
	pos1 = pos1.Div(s.Cam.Pos.Z)
	// finally center on screen
	pos1 = pos1.Add(V2(float64(Width/2), float64(Height/2)))
	return pos1
}

func (s*SDLScreen) ScreenToWorld(v Vec) Vec {
	pos1 := v
	// uncenter from screen
	pos1 = pos1.Sub(V2(float64(Width/2), float64(Height/2)))
	if s.Cam.Pos.Z != 0 {
		// unzoom from camera
		pos1 = pos1.Mul(s.Cam.Pos.Z)
	}
	// place in world space
	pos1 = pos1.Div(UNIT_SIZE)
	return pos1
}

// get a world-to-screen projection (2 coordinates) from a world position and a scale
func (s*SDLScreen) project(v, size Vec) (Vec, Vec) {
	// first take the position
	pos1 := v.Add(s.Cam.Pos)
	// then place it in world space
	pos1 = pos1.Mul(UNIT_SIZE)
	// get the corner positions using the unit size
	// top left corner
	pos1 = pos1.Sub(size.Mul(UNIT_SIZE))
	// bottom right corner
	pos2 := pos1.Add(size.Mul(UNIT_SIZE))
	// zoom with the camera
	pos1 = pos1.Div(s.Cam.Pos.Z)
	pos2 = pos2.Div(s.Cam.Pos.Z)
	// finally center on screen
	pos1 = pos1.Add(V2(float64(Width/2), float64(Height/2)))
	pos2 = pos2.Add(V2(float64(Width/2), float64(Height/2)))
	return pos1, pos2
}

// draw a single instance of a drawable
// NOTE the drawing convention is that pos stores the coordinates, NOT the x & y and w & h
// TODO implement a z buffer
func (s *SDLScreen) DrawRectScreen(v1, v2 Vec, col Vec, flags uint8) {
	drawable := Drawable{
		Type:   D_RECT,
		Colour: col,
		Flags:  flags,
		V1:     v1,
		V2:     v2,
	}
	// add more layers until we get the desired z-depth
	s.matchLayers(int(v1.Z))
	// add the drawable to the layer
	s.ZBuf.Layers[int(v1.Z)].Drawables = append(s.ZBuf.Layers[int(v1.Z)].Drawables, drawable)
}

func (s *SDLScreen) DrawRectWorld(v1, size, col Vec, flags uint8) {
	pos1, pos2 := s.project(v1, size)
	drawable := Drawable{
		Type:   D_RECT,
		Colour: col,
		Flags:  flags,
		V1:     pos1,
		V2:     pos2,
	}
	// add more layers until we get the desired z-depth
	s.matchLayers(int(v1.Z))
	// add the drawable to the layer
	s.ZBuf.Layers[int(v1.Z)].Drawables = append(s.ZBuf.Layers[int(v1.Z)].Drawables, drawable)
}

// draw a single instance of a drawable
// NOTE the drawing convention is that pos stores the coordinates, NOT the x & y and w & h
func (s *SDLScreen) DrawTextureScreen(v1, v2 Vec, texture *sdl.Texture, angle float64) {
	drawable := Drawable{
		Type:  D_TEXTURE,
		Data:  texture,
		V1:    v1,
		V2:    v2,
		Angle: angle,
	}
	// add more layers until we get the desired z-depth
	s.matchLayers(int(v1.Z))
	// add the drawable to the layer
	s.ZBuf.Layers[int(v1.Z)].Drawables = append(s.ZBuf.Layers[int(v1.Z)].Drawables, drawable)
}

func (s *SDLScreen) DrawTextureWorld(v1, size Vec, texture *sdl.Texture, angle float64) {
	// get the world projection
	pos1, pos2 := s.project(v1, size)
	drawable := Drawable{
		Type:  D_TEXTURE,
		Data:  texture,
		V1:    pos1,
		V2:    pos2,
		Angle: angle,
	}
	// add more layers until we get the desired z-depth
	s.matchLayers(int(v1.Z))
	// add the drawable to the layer
	s.ZBuf.Layers[int(v1.Z)].Drawables = append(s.ZBuf.Layers[int(v1.Z)].Drawables, drawable)
}

// draw a line between points v1 and v2 with a colour
func (s *SDLScreen) DrawLineScreen(v1, v2 Vec, col Vec) {
	drawable := Drawable{
		Type:   D_LINE,
		Colour: col,
		V1:     v1,
		V2:     v2,
	}
	// add more layers until we get the desired z-depth
	s.matchLayers(int(v1.Z))
	// add the drawable to the layer
	s.ZBuf.Layers[int(v1.Z)].Drawables = append(s.ZBuf.Layers[int(v1.Z)].Drawables, drawable)
}

// draw a line between points v1 and v2 with a colour
func (s *SDLScreen) DrawLineWorld(v1, v2 Vec, col Vec) {
	pos1 := s.WorldToScreen(v1)
	pos2 := s.WorldToScreen(v2)
	drawable := Drawable{
		Type:   D_LINE,
		Colour: col,
		V1:     pos1,
		V2:     pos2,
	}
	// add more layers until we get the desired z-depth
	s.matchLayers(int(v1.Z))
	// add the drawable to the layer
	s.ZBuf.Layers[int(v1.Z)].Drawables = append(s.ZBuf.Layers[int(v1.Z)].Drawables, drawable)
}

// Get the mouse position in the screen
func (s*SDLScreen) MousePosScreen() Vec {
	x, y, _ := sdl.GetMouseState()
	return V2(float64(x), float64(y))
}

// Get the mouse position in the world
func (s*SDLScreen) MousePosWorld() Vec {
	x, y, _ := sdl.GetMouseState()
	return s.ScreenToWorld(V2(float64(x), float64(y)))
}

// poll window events (i.e. window moves and inputs)
func (s *SDLScreen) PollEvents() {
	var event sdl.Event
	// handle events, in this case escape key and close window
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.WindowEvent:
			if t.Event == sdl.WINDOWEVENT_RESIZED {
				if Engine.Options().Resizable {
					Width, Height = s.Window.GetSize()
					s.Window.SetSize(Width, Height)
				}
			}
		case *sdl.QuitEvent:
			// running = false
			Engine.ForceStop()
		case *sdl.KeyboardEvent:
			if t.Repeat == 1 {
				Inputs.SetActive(SDLKeyToString(t.Keysym.Sym), HELD, 0)
			} else {
				if t.Type == sdl.KEYDOWN {
					Inputs.SetActive(SDLKeyToString(t.Keysym.Sym), PRESSED, 0)
				} else if t.Type == sdl.KEYUP {
					Inputs.SetActive(SDLKeyToString(t.Keysym.Sym), RELEASED, 0)
				}
			}
		case *sdl.MouseButtonEvent:
			if t.Type == sdl.MOUSEBUTTONDOWN {
				Inputs.SetActive(SDLButtonToString(t.Button), PRESSED, 0)
			} else if t.Type == sdl.MOUSEBUTTONUP {
				Inputs.SetActive(SDLButtonToString(t.Button), RELEASED, 0)
			}
		case *sdl.MouseWheelEvent:
			if t.Type == sdl.MOUSEWHEEL {
				Inputs.SetActive("mousewheel", SCROLLED, float64(t.Y))
			}
		}
	}
}

func (s *SDLScreen) RendererPrepare() {
	s.Renderer.SetDrawColor(0, 0, 0, 0)
	s.Renderer.Clear()
}

func (s *SDLScreen) FetchMeshComponents(){
	for _, entity := range ECS.Entities{
		if entity.Active{
			// fetch the transform and the mesh of the entity
			var transform, mesh IComponent
			for _, component := range entity.Components{
				switch c:=component.(type){
				case *MeshComponent:
					mesh = c
				case *TransformComponent:
					transform = c
				}
			}
			// if the entity has a mesh, and a transform, then draw it
			if transform != nil && mesh != nil{
				t := transform.(*TransformComponent)
				m := mesh.(*MeshComponent)
				s.DrawTextureScreen(t.Pos.Sub(m.Texture.Size.Div(2)), t.Pos.Add(m.Texture.Size.Div(2)), m.Texture.Data, 0)
			}
		}
	}
}

// method to actually draw to the screen. called once per frame
func (s *SDLScreen) RendererFlush() {
	var previousColour Vec
	// iterate over each layer and the drawable in that layer
	for _, layer := range s.ZBuf.Layers {
		for _, drawable := range layer.Drawables {
			// set the colour
			if !drawable.Colour.Equals(previousColour) {
				previousColour = drawable.Colour
			}
			s.Renderer.SetDrawColor(uint8(previousColour.X), uint8(previousColour.Y), uint8(previousColour.Z), uint8(previousColour.W))
			// actually call SDL draw function
			switch drawable.Type {
			case D_RECT:
				if drawable.Flags&NOFILL > 0 {
					s.Renderer.DrawRect(&sdl.Rect{int32(drawable.V1.X), int32(drawable.V1.Y), int32(drawable.V2.X - drawable.V1.X), int32(drawable.V2.Y - drawable.V1.Y)})
				} else {
					s.Renderer.FillRect(&sdl.Rect{int32(drawable.V1.X), int32(drawable.V1.Y), int32(drawable.V2.X - drawable.V1.X), int32(drawable.V2.Y - drawable.V1.Y)})
				}
				break
			case D_LINE:
				s.Renderer.DrawLine(int32(drawable.V1.X), int32(drawable.V1.Y), int32(drawable.V2.X), int32(drawable.V2.Y))
				break
			case D_TEXTURE:
				s.Renderer.CopyEx(drawable.Data.(*sdl.Texture), nil, &sdl.Rect{int32(drawable.V1.X), int32(drawable.V1.Y), int32(drawable.V2.X - drawable.V1.X), int32(drawable.V2.Y - drawable.V1.Y)}, drawable.Angle, nil, sdl.FLIP_NONE)
				break
			}
		}
	}
	// clear the z-buffer and draw to the screen
	s.ZBuf.Layers = nil
	s.Renderer.Present()
}

// initilise SDL and the global renderer
func (s *SDLScreen) Setup() {
	// initialise SDL
	err := sdl.Init(sdl.INIT_EVERYTHING)
	Assert(err==nil, "failed to initialize sdl")
	// Using the SDL_ttf library so need to initialize it before using it
	err = ttf.Init()
	Assert(err==nil, "failed to initialize TTF")
	window, err := sdl.CreateWindow(Engine.Options().Title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		Engine.Options().Width, Engine.Options().Width, sdl.WINDOW_RESIZABLE)
	Assert(err==nil, "failed to create graphics")
	ren, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	Assert(err==nil, "failed to create renderer")
	// enable alpha blending
	ren.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	// disable anti-aliasting
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")
	// enable batching
	sdl.SetHint(sdl.HINT_RENDER_BATCHING, "1")
	// assign the members to the renderer
	s.Renderer = ren
	s.Window = window
}

// close the renderer and close SDL
func (s *SDLScreen) Close() {
	s.Renderer.Destroy()
	s.Window.Destroy()
	sdl.Quit()
}

// Load a texture and return the texture data
func (s *SDLScreen) LoadTexture(path string) *Texture {
	// Load the image into memory
	surfaceImg, err := img.Load(path)
	Assert(err==nil, "cannot load image into surface")
	width := surfaceImg.W
	height := surfaceImg.H
	// Put the image on the GPU
	texture, err := s.Renderer.CreateTextureFromSurface(surfaceImg)
	Assert(err==nil, "cannot create texture from surface")
	// Free the surface in RAM, texture remains in GPU
	surfaceImg.Free()
	return &Texture{
		Data: texture,
		Size: V2(float64(width), float64(height)),
	}
}