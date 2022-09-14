package gl

import (
	"encoding/binary"
	"errors"
	"os"
)

type texture struct {
	name string
	width, height uint
	pixels [][]color
}


func NewTexture(fileName string) (*texture, error) {
	// t := texture{width: width, height: height, pixels: [][]color{}}
	// Open the file
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()	// Close the file when the process is done
	headerSizeBits := make([]byte, 4)
  f.ReadAt(headerSizeBits, 10)
	headerSize := fromByteToInt(headerSizeBits)
	
	widthBits := make([]byte, 4)
	f.ReadAt(widthBits, 18)
	width := fromByteToInt(widthBits)

	heightBits := make([]byte, 4)
	f.ReadAt(heightBits, 22)
	height := fromByteToInt(heightBits)

	f.Seek(int64(headerSize), 0)
	pixels := [][]color{}

	for x := 0; x < int(height); x++ {
		col := []color{}
		for y := 0; y < int(width); y++ {
			bByte := make([]byte,1)
			gByte := make([]byte,1)
			rByte := make([]byte,1)
			f.Read(bByte)
			f.Read(gByte)
			f.Read(rByte)
			// r := fromByteToInt(bByte)
			r := (float64(rByte[0]) / 255.0)
			g := (float64(gByte[0]) / 255.0)
			b := (float64(bByte[0]) / 255.0)
			clr, errC := NewColor(r,g,b)
			if errC != nil {
				
			} else {
				col = append(col, *clr)
			}
		}
		pixels = append(pixels, col)
	}
	texture := texture{}
  texture.width = uint(width)
	texture.height = uint(height)
	texture.pixels = pixels
	texture.name = fileName
	return &texture, nil
}

func fromByteToInt(bytes []byte) uint32 {
	return binary.LittleEndian.Uint32(bytes)
}


func (t *texture) GetColor(u,v float32) (*color, error) {

	if 0 <= u && u <= 1 && 0 <= v && v <= 1{
		return &t.pixels[int((v) * float32(t.height))][int((u) * float32(t.width))], nil
	} 
	return nil, errors.New("Cannot get the color from texture.")
}

// Get the texture pixels
func (t *texture) Pixels() [][]color {
	return t.pixels
}