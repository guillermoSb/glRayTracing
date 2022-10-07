package gl

import (
	"guillermoSb/glRayTracing/numg"
	"math"
)

type torus struct {
	r        float64 // r
	r2       float64 // R
	material material
}

// Returns a new torus
func NewTorus(r float64, r2 float64, material material) *torus {
	t := torus{r: r, r2: r2, material: material}
	return &t
}

// Check if there is an intersection on the torus
func (torus *torus) rayIntersect(origin, dir numg.V3) *intersect {

	objectMatrix, _ := numg.Identity(4)
	rx, _ := numg.Identity(4)

	// angle
	angle := math.Pi / 2

	rx[1][1] = float32(math.Cos(angle))
	rx[1][2] = -float32(math.Sin(angle))
	rx[2][1] = float32(math.Sin(angle))
	rx[2][2] = float32(math.Cos(angle))

	// rx[0][0] = float32(math.Cos(angle))
	// rx[0][1] = float32(-math.Sin(angle))
	// rx[1][0] = float32(math.Sin(angle))
	// rx[1][1] = float32(math.Cos(angle))

	// ry, _ := numg.Identity(4)
	// rz, _ := numg.Identity(4)

	mr := rx

	objectMatrix = mr
	v := numg.M{{float32(dir.X)}, {float32(dir.Y)}, {float32(dir.Z)}, {1}}
	vt, _ := numg.MultiplyMatrices(objectMatrix, v)

	dirF := numg.V3{X: float64(vt[0][0] / vt[3][0]), Y: float64(vt[1][0] / vt[3][0]), Z: float64(vt[2][0] / vt[3][0])}
	dir = dirF
	v = numg.M{{float32(origin.X)}, {float32(origin.Y)}, {float32(origin.Z)}, {1}}
	vt, _ = numg.MultiplyMatrices(objectMatrix, v)

	origF := numg.V3{X: float64(vt[0][0] / vt[3][0]), Y: float64(vt[1][0] / vt[3][0]), Z: float64(vt[2][0] / vt[3][0])}
	origin = origF
	// r = origin
	sumDirQuad := math.Pow(dir.X, 2) + math.Pow(dir.Y, 2) + math.Pow(dir.Z, 2)
	sumOriginQuad := math.Pow(origin.X, 2) + math.Pow(origin.Y, 2) + math.Pow(origin.Z, 2)
	rQuad := math.Pow(torus.r, 2)
	r2Quad := math.Pow(torus.r2, 2)
	originDirDotProduct := numg.V3DotProduct(origin, dir)
	c4 := math.Pow(sumDirQuad, 2)
	c3 := 4.0 * sumDirQuad * originDirDotProduct
	c2 := 2.0*sumDirQuad*(sumOriginQuad-(r2Quad+rQuad)) + 4.0*math.Pow(originDirDotProduct, 2) + 4.0*r2Quad*math.Pow(dir.Y, 2)
	c1 := 4.0*(sumOriginQuad-(r2Quad+rQuad))*originDirDotProduct + 8.0*r2Quad*origin.Y*dir.Y
	c0 := ((sumOriginQuad - (r2Quad + rQuad)) * (sumOriginQuad - (r2Quad + rQuad))) - 4.0*r2Quad*(rQuad-math.Pow(origin.Y, 2))

	solution := numg.Solve4([]float64{c0, c1, c2, c3, c4})
	if len(solution) == 0 {
		return nil
	}

	shortestT := math.Inf(1)

	for i := 0; i < len(solution); i++ {
		t := solution[i]
		if t > 0 {
			if t < shortestT {
				shortestT = t
			}
		}
	}
	point := numg.V3{X: origin.X + dir.X*shortestT, Y: origin.Y + dir.Y*shortestT, Z: origin.Z + dir.Z*shortestT}

	paramSquared := math.Pow(torus.r2, 2) + math.Pow(torus.r, 2)
	x := point.X
	y := point.Y
	z := point.Z
	sumSquared := math.Pow(x, 2) + math.Pow(y, 2) + math.Pow(z, 2)
	tmp := numg.V3{
		X: 4.0 * x * (sumSquared - paramSquared),
		Y: 4.0 * y * (sumSquared - paramSquared + 2.0*math.Pow(torus.r2, 2)),
		Z: 4.0 * z * (sumSquared - paramSquared),
	}
	return NewIntersect(shortestT, point, numg.NormalizeV3(tmp), struct{ material }{
		material: torus.material,
	})

}
