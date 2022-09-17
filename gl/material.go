package gl

type material struct {
	diffuse color
}

func  NewMaterial(diffuse color) *material {
	m := material{diffuse: diffuse}
	return &m
}