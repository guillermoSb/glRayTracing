package gl

import (
	"guillermoSb/glRayTracing/numg"
	"math"
)

type plane struct {
	material material
	position numg.V3
	normal   numg.V3
}

func NewPlane(material material, position numg.V3, normal numg.V3) *plane {
	p := plane{material: material, position: position, normal: numg.NormalizeV3(normal)}
	return &p
}

func (plane *plane) rayIntersect(origin, dir numg.V3) *intersect {
	denom := numg.V3DotProduct(numg.NormalizeV3(dir), plane.normal)
	if math.Abs(denom) > 0.0001 {
		// t (distance) = ((plane.position - origin) o plane.normal) / (dir o plane.normal))
		num := numg.V3DotProduct(numg.Subtract(plane.position, origin), plane.normal)
		t := (num / denom)
		if t > 0 {
			p := numg.Add(origin, numg.MultiplyVectorWithConstant(numg.NormalizeV3(dir), t))
			return NewIntersect(
				t,
				p,
				plane.normal,
				struct{ material }{
					material: plane.material,
				})
		}
	}
	return nil
}
