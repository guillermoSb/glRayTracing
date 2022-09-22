package gl

const OPAQUE = 0
const REFLECTIVE = 1
const TRANSPARENT = 2

type material struct {
	diffuse color
	specularity float64
	matType int
	ior float64
}

func  NewMaterial(diffuse color, specularity float64, matType int, ior float64) *material {
	m := material{diffuse: diffuse, specularity: specularity, matType: matType, ior: ior}
	return &m
}