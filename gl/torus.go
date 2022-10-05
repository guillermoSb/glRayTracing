package gl

import (
	"guillermoSb/glRayTracing/numg"
	"math"
)

type torus struct {
	r  float64 // r
	r2 float64 // R
}

// Returns a new torus
func NewTorus(r float64, r2 float64) *torus {
	t := torus{r: r, r2: r2}
	return &t
}

// Check if there is an intersection on the torus
func (torus *torus) rayIntersect(origin, dir numg.V3) {
	// r = origin
	sumDirQuad := math.Pow(dir.X, 2) + math.Pow(dir.Y, 2) + math.Pow(dir.Z, 2)
	sumOriginQuad := math.Pow(origin.X, 2) + math.Pow(origin.Y, 2) + math.Pow(origin.Z, 2)
	rQuad := math.Pow(torus.r, 2)
	r2Quad := math.Pow(torus.r2, 2)
	c4 := sumDirQuad
	c3 := 4 * sumDirQuad * (origin.X*dir.X + origin.Y*dir.Y + origin.Z*dir.Z)
	c2 := 2*sumDirQuad*(sumOriginQuad-(r2Quad+rQuad)) + 4*(origin.X*dir.X+origin.Y*dir.Y+origin.Z*dir.Z)*(origin.X*dir.X+origin.Y*dir.Y+origin.Z*dir.Z) + 4*r2Quad*rQuad
	c1 := 4*(sumOriginQuad-(r2Quad+rQuad))*(origin.X*dir.X+origin.Y*dir.Y+origin.Z*dir.Z) + 8*r2Quad*origin.Y*dir.Y
	c0 := (sumOriginQuad-(r2Quad+rQuad))*(sumOriginQuad-(r2Quad+rQuad)) - 4*r2Quad*(rQuad-math.Pow(origin.Y, 2))
}
