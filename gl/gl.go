package gl

type renderer struct {
	width, height float64
}

func NewRenderer(width, height float64)(*renderer, error) {
	r := renderer{width: width, height: height} 
	return &r, nil
}


func (r *renderer) String() string {
	return "Hello"
}