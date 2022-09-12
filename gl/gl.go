package gl

type renderer struct {
	width, height int
}

func NewRenderer(width, height int)(*renderer, error) {
	r := renderer{width: width, height: height} 
	return &r, nil
}

// TODO: Create color struct
// TODO: Create a renderer with height, width
// TODO: Be able to export a black image
// TODO: Be able to draw a point on the image
// TODO: Be able to load images and use any aspect ratio