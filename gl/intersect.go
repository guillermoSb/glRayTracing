package gl

import "guillermoSb/glRayTracing/numg"

type intersect struct{
	obj struct {
		material
	}
	point numg.V3
	distance float64
	normal numg.V3
}

func NewIntersect(distance float64, point,normal numg.V3, obj struct{material}) *intersect {
	i := intersect{obj: obj, distance: distance, point: point, normal: normal}
	return &i
}