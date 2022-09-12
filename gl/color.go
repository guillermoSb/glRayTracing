package gl

import "fmt"

// Struct that defines the properties of a color
// - r: value of the red color
// - g: value of the green color
// - b: value of the blue color
type color struct {
	r,g,b float64
}



// Creates a new color from its rgb components
// - r: red color
// - g: green color
// - b: blue color
func NewColor(r,g,b float64) (*color, error) {
	if r > 1 || g > 1 || b > 1 || r < 0 || g < 0 || b < 0{
			return nil, fmt.Errorf("Can't create color from the received values. Expected r=%.1f, g=%.1f, b=%.1f to be within 0 and 1.", r,g,b)
	}
	clr := color{r,g,b}
	return &clr, nil
}


// Obtain the red value
func (c *color) Red() float64 {
	return c.r
}
// Obtain the green value
func (c *color) Green() float64 {
	return c.g
}
// Obtain the blue value
func (c *color) Blue() float64 {
	return c.g
}