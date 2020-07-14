package ske
//package ske
//
//import (
//	"fmt"
//	"os"
//
//	"github.com/veandco/go-sdl2/sdl"
//	"github.com/veandco/go-sdl2/ttf"
//)
//
//// SDL implementation of a renderer
//type SDLScreen struct {
//	Window   *sdl.Window
//	Renderer *sdl.Renderer
//	ZBuf     ZBuffer
//}
//
//var (
//	Screen   *SDLScreen
//	Width    int32 = 800
//	Height   int32 = 600
//)
//
//const (
//	UNIT_SIZE float64 = 50
//
//	NOFILL = 0x1 << 0
//	FILL   = 0x1 << 1
//
//	FLIP = 0x1 << 0
//
//	// max of 100 drawables per layer
//	MAX_LAYER_DRAWS = 100
//
//	D_RECT    = 0x0
//	D_LINE    = 0x1
//	D_TEXTURE = 0x2
//	D_TEXT    = 0x3
//)
//
//type Layer struct {
//	Drawables []Drawable
//}
//
//type Drawable struct {
//	Type   uint8
//	Data   interface{}
//	Colour Vec
//	Flags  uint8
//	V1     Vec
//	V2     Vec
//	Angle  float64
//}
//
//// store sorted layers
//type ZBuffer struct {
//	Layers []Layer
//}
//
//// iteratively add layers to the z-buffer until it matches the desired z-value
//func matchLayers(z int) {
//	// add more layers until we get the desired z-depth
//	l := len(Renderer.ZBuf.Layers)
//	for z >= l {
//		Renderer.ZBuf.Layers = append(Renderer.ZBuf.Layers, Layer{})
//		l = len(Renderer.ZBuf.Layers)
//	}
//}
//
//func worldToScreen(v Vec) Vec {
//	// first take the position
//	pos1 := v.Add(Cam.Pos)
//	// then place it in screen space
//	pos1 = pos1.Mul(UNIT_SIZE)
//	// zoom with camera
//	pos1 = pos1.Div(Cam.Pos.Z)
//	// finally center on screen
//	pos1 = pos1.Add(V2(float64(Width/2), float64(Height/2)))
//	return pos1
//}
//
//func screenToWorld(v Vec) Vec {
//	pos1 := v
//	// uncenter from screen
//	pos1 = pos1.Sub(V2(float64(Width/2), float64(Height/2)))
//	if Cam.Pos.Z != 0 {
//		// unzoom from camera
//		pos1 = pos1.Mul(Cam.Pos.Z)
//	}
//	// place in world space
//	pos1 = pos1.Div(UNIT_SIZE)
//	return pos1
//}
//
//// get a world-to-screen projection (2 coordinates) from a world position and a scale
//func project(v, size Vec) (Vec, Vec) {
//	// first take the position
//	pos1 := v.Add(Cam.Pos)
//	// then place it in world space
//	pos1 = pos1.Mul(UNIT_SIZE)
//	// get the corner positions using the unit size
//	// top left corner
//	pos1 = pos1.Sub(size.Mul(UNIT_SIZE))
//	// bottom right corner
//	pos2 := pos1.Add(size.Mul(UNIT_SIZE))
//	// zoom with the camera
//	pos1 = pos1.Div(Cam.Pos.Z)
//	pos2 = pos2.Div(Cam.Pos.Z)
//	// finally center on screen
//	pos1 = pos1.Add(V2(float64(Width/2), float64(Height/2)))
//	pos2 = pos2.Add(V2(float64(Width/2), float64(Height/2)))
//	return pos1, pos2
//}
//
//// draw a single instance of a drawable
//// NOTE the drawing convention is that pos stores the coordinates, NOT the x & y and w & h
//// TODO implement a z buffer
//func DrawRectScreen(v1, v2 Vec, col Vec, flags uint8) {
//	drawable := Drawable{
//		Type:   D_RECT,
//		Colour: col,
//		Flags:  flags,
//		V1:     v1,
//		V2:     v2,
//	}
//	// add more layers until we get the desired z-depth
//	matchLayers(int(v1.Z))
//	// add the drawable to the layer
//	Renderer.ZBuf.Layers[int(v1.Z)].Drawables = append(Renderer.ZBuf.Layers[int(v1.Z)].Drawables, drawable)
//}
//
//func DrawRectWorld(v1, size, col Vec, flags uint8) {
//	pos1, pos2 := project(v1, size)
//	drawable := Drawable{
//		Type:   D_RECT,
//		Colour: col,
//		Flags:  flags,
//		V1:     pos1,
//		V2:     pos2,
//	}
//	// add more layers until we get the desired z-depth
//	matchLayers(int(v1.Z))
//	// add the drawable to the layer
//	Renderer.ZBuf.Layers[int(v1.Z)].Drawables = append(Renderer.ZBuf.Layers[int(v1.Z)].Drawables, drawable)
//}
//
//// draw a single instance of a drawable
//// NOTE the drawing convention is that pos stores the coordinates, NOT the x & y and w & h
//func DrawTextureScreen(v1, v2 Vec, resource *Resource, angle float64) {
//	drawable := Drawable{
//		Type:  D_TEXTURE,
//		Data:  resource.Data.(*sdl.Texture),
//		V1:    v1,
//		V2:    v2,
//		Angle: angle,
//	}
//	// add more layers until we get the desired z-depth
//	matchLayers(int(v1.Z))
//	// add the drawable to the layer
//	Renderer.ZBuf.Layers[int(v1.Z)].Drawables = append(Renderer.ZBuf.Layers[int(v1.Z)].Drawables, drawable)
//}
//
//func DrawTextureWorld(v1, size Vec, resource *Resource, angle float64) {
//	// get the world projection
//	pos1, pos2 := project(v1, size)
//	drawable := Drawable{
//		Type:  D_TEXTURE,
//		Data:  resource.Data.(*sdl.Texture),
//		V1:    pos1,
//		V2:    pos2,
//		Angle: angle,
//	}
//	// add more layers until we get the desired z-depth
//	matchLayers(int(v1.Z))
//	// add the drawable to the layer
//	Renderer.ZBuf.Layers[int(v1.Z)].Drawables = append(Renderer.ZBuf.Layers[int(v1.Z)].Drawables, drawable)
//}
//
//// draw a line between points v1 and v2 with a colour
//func DrawLineScreen(v1, v2 Vec, col Vec) {
//	drawable := Drawable{
//		Type:   D_LINE,
//		Colour: col,
//		V1:     v1,
//		V2:     v2,
//	}
//	// add more layers until we get the desired z-depth
//	matchLayers(int(v1.Z))
//	// add the drawable to the layer
//	Renderer.ZBuf.Layers[int(v1.Z)].Drawables = append(Renderer.ZBuf.Layers[int(v1.Z)].Drawables, drawable)
//}
//
//// draw a line between points v1 and v2 with a colour
//func DrawLineWorld(v1, v2 Vec, col Vec) {
//	pos1 := worldToScreen(v1)
//	pos2 := worldToScreen(v2)
//	drawable := Drawable{
//		Type:   D_LINE,
//		Colour: col,
//		V1:     pos1,
//		V2:     pos2,
//	}
//	// add more layers until we get the desired z-depth
//	matchLayers(int(v1.Z))
//	// add the drawable to the layer
//	Renderer.ZBuf.Layers[int(v1.Z)].Drawables = append(Renderer.ZBuf.Layers[int(v1.Z)].Drawables, drawable)
//}
//
//// Get the mouse position in the screen
//func MousePosScreen() Vec {
//	x, y, _ := sdl.GetMouseState()
//	return V2(float64(x), float64(y))
//}
//
//// Get the mouse position in the world
//func MousePosWorld() Vec {
//	x, y, _ := sdl.GetMouseState()
//	return screenToWorld(V2(float64(x), float64(y)))
//}
//
//// poll window events (i.e. window moves and inputs)
//func PollEvents() {
//	var event sdl.Event
//	// handle events, in this case escape key and close window
//	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
//		switch t := event.(type) {
//		case *sdl.WindowEvent:
//			if t.Event == sdl.WINDOWEVENT_RESIZED {
//				Width, Height = Renderer.Window.GetSize()
//				Renderer.Window.SetSize(Width, Height)
//			}
//		case *sdl.QuitEvent:
//			ForceQuit()
//		case *sdl.KeyboardEvent:
//			if t.Repeat == 1 {
//				InputBuffer.KeyHeld = t.Keysym.Sym
//			} else {
//				if t.Type == sdl.KEYDOWN {
//					InputBuffer.KeyPressed = t.Keysym.Sym
//				} else if t.Type == sdl.KEYUP {
//					InputBuffer.KeyReleased = t.Keysym.Sym
//				}
//			}
//		case *sdl.MouseButtonEvent:
//			if t.Type == sdl.MOUSEBUTTONDOWN {
//				InputBuffer.MousePressed = int8(t.Button)
//			} else if t.Type == sdl.MOUSEBUTTONUP {
//				InputBuffer.MouseReleased = int8(t.Button)
//			}
//
//		case *sdl.MouseWheelEvent:
//			if t.Type == sdl.MOUSEWHEEL {
//				InputBuffer.MouseScrolled = int8(t.Y)
//			}
//		}
//	}
//}
//
//func RendererPrepare() {
//	Renderer.Renderer.SetDrawColor(0, 0, 0, 0)
//	Renderer.Renderer.Clear()
//}
//
//// method to actually draw to the screen. called once per frame
//func RendererFlush() {
//	var previousColour Vec
//	// iterate over each layer and the drawable in that layer
//	for _, layer := range Renderer.ZBuf.Layers {
//		for _, drawable := range layer.Drawables {
//			// set the colour
//			if !drawable.Colour.Equals(previousColour) {
//				previousColour = drawable.Colour
//			}
//			Renderer.Renderer.SetDrawColor(uint8(previousColour.X), uint8(previousColour.Y), uint8(previousColour.Z), uint8(previousColour.W))
//			// actually call SDL draw function
//			switch drawable.Type {
//			case D_RECT:
//				if drawable.Flags&NOFILL > 0 {
//					Renderer.Renderer.DrawRect(&sdl.Rect{int32(drawable.V1.X), int32(drawable.V1.Y), int32(drawable.V2.X - drawable.V1.X), int32(drawable.V2.Y - drawable.V1.Y)})
//				} else {
//					Renderer.Renderer.FillRect(&sdl.Rect{int32(drawable.V1.X), int32(drawable.V1.Y), int32(drawable.V2.X - drawable.V1.X), int32(drawable.V2.Y - drawable.V1.Y)})
//				}
//				break
//			case D_LINE:
//				Renderer.Renderer.DrawLine(int32(drawable.V1.X), int32(drawable.V1.Y), int32(drawable.V2.X), int32(drawable.V2.Y))
//				break
//			case D_TEXTURE:
//				Renderer.Renderer.CopyEx(drawable.Data.(*sdl.Texture), nil, &sdl.Rect{int32(drawable.V1.X), int32(drawable.V1.Y), int32(drawable.V2.X - drawable.V1.X), int32(drawable.V2.Y - drawable.V1.Y)}, drawable.Angle, nil, sdl.FLIP_NONE)
//				break
//			}
//		}
//	}
//	// clear the z-buffer and draw to the screen
//	Renderer.ZBuf.Layers = nil
//	Renderer.Renderer.Present()
//}
//
//// initilise SDL and the global renderer
//func RendererSetup() {
//	Renderer = &RendererSDL{
//		ZBuf: ZBuffer{},
//	}
//	// initialise SDL
//	err := sdl.Init(sdl.INIT_EVERYTHING)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Failed to initialize sdl: %s\n", err)
//		os.Exit(1)
//	}
//	// Using the SDL_ttf library so need to initialize it before using it
//	if err = ttf.Init(); err != nil {
//		fmt.Printf("Failed to initialize TTF: %s\n", err)
//		os.Exit(2)
//	}
//	window, err := sdl.CreateWindow("eri", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
//		Width, Height, sdl.WINDOW_RESIZABLE)
//	if err != nil {
//		fmt.Fprint(os.Stderr, "Failed to create graphics: %s\n", err)
//		os.Exit(2)
//	}
//	ren, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
//	if err != nil {
//		ren.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
//		fmt.Fprint(os.Stderr, "Failed to create graphics: %s\n", err)
//		os.Exit(2)
//	}
//	// disable anti-aliasting
//	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")
//	sdl.SetHint(sdl.HINT_RENDER_BATCHING, "1")
//	// assign the members to the renderer
//	Renderer.Renderer = ren
//	Renderer.Window = window
//}
//
//// close the renderer and close SDL
//func RendererClose() {
//	Renderer.Renderer.Destroy()
//	Renderer.Window.Destroy()
//	sdl.Quit()
//}
