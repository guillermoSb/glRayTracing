package gl

import (
	"guillermoSb/glRayTracing/numg"
	"log"
	"os"
)

type renderer struct {
	width, height uint
	pixels [][]color
	background texture
}

// Creates a new renderer and sends the reference
// - width: width of the renderer
// - height: height of the renderer
// - background: name of the background
func NewRenderer(width, height uint, background string)(*renderer, error) {
	r := renderer{width: width, height: height, pixels: [][]color{}} 
	// Create the pixel array
	for x := 0; x < int(width); x++ {
		col := []color{}
		for y := 0; y < int(height); y++ {
			clearClr, _ := NewColor(0.2,0.2,0.2)	// Use black as the background color
			col = append(col, *clearClr)	// append the color to the column
		}
		r.pixels = append(r.pixels, col)	// append the column to the pixels
	}
	// Load the background
	if background != "" {
		t, errT := NewTexture(background)
		if errT != nil {
			return nil, errT
		}
		r.background = *t
	}
	return &r, nil
}

// Get the pixel array
func (r *renderer) Pixels() [][]color {
	return r.pixels
}

func (r *renderer) GLDrawBackground() {
	for x := 0; x < int(r.width); x++ {
		for y := 0; y < int(r.height); y++ {
			ux := (float32(x)/float32(r.width))	// value for x 0 -1
			uy := (float32(y)/float32(r.height))	// value for y 0 -1
			texX := int(ux * float32(r.background.width))	
			texY := int(uy * float32(r.background.height))
			r.GLPoint(numg.V2{X:float64(x),Y:float64(y)}, r.background.pixels[texY][texX])
		}	
	}
}

// Draw a point on the screen
func (r *renderer) GLPoint(point numg.V2, clr color) {
	// Check tht the point is within the screen bounds
	
	if int(point.X) < 0 || 
	int(point.X) >= int(r.width) || 
	int(point.Y) < 0 || 
	int(point.Y) >= int(r.height) {
		return
	} 
	// Set the color on the pixel
	r.pixels[int(point.X)][int(point.Y)] = clr
}

// Create the renderer with the pixel array
// filename: The name of the file
func (r *renderer) GlFinish(fileName string) {
	// Attempt to open the file
	f, err := os.Create(fileName)

	// Check if the file was successfully created
	if err != nil {
		log.Fatal(err)
	}
	// Example, writing 5 x 5 image
	defer f.Close()	// Close the file when the process is done
	f.Write([]byte("B"))
	f.Write([]byte("M"))
	f.Write(numg.Dword(uint32(r.width) * uint32(r.height * 3)))	// File Size
	f.Write([]byte{0, 0})	// Reserved
	f.Write([]byte{0, 0})	// Reserved
	f.Write([]byte{54, 0, 0, 0 })	// ?
	f.Write([]byte{40, 0, 0, 0})	// Header Size
	f.Write(numg.Dword(uint32(r.width)))		// Width
	f.Write(numg.Dword(uint32(r.height)))		// Height
	f.Write([]byte{1, 0})	// Plane
	f.Write([]byte{24, 0})	// BPP
	f.Write([]byte{0,0,0,0})
	f.Write([]byte{0,0,0,0})
	f.Write([]byte{0,0,0,0})
	f.Write([]byte{0,0,0,0})
	f.Write([]byte{0,0,0,0})
	f.Write([]byte{0,0,0,0})
	// Pixel Data
	for i := 0; i < int(r.height); i++ {
		for j := 0; j < int(r.width); j++ {
			f.Write(r.pixels[j][i].Bytes())
		}
	}
}
// TODO: Be able to load images and use any aspect ratio
