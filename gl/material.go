package gl

type material struct {
	diffuse color
	specularity float64
}

func  NewMaterial(diffuse color, specularity float64) *material {
	m := material{diffuse: diffuse, specularity: specularity}
	return &m
}