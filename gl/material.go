package gl

const OPAQUE = 0
const REFLECTIVE = 1

type material struct {
	diffuse color
	specularity float64
	matType int
}

func  NewMaterial(diffuse color, specularity float64, matType int) *material {
	m := material{diffuse: diffuse, specularity: specularity, matType: matType}
	return &m
}