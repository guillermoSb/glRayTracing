package gl

import "guillermoSb/glRayTracing/numg"

type figure interface {
	rayIntersect(origin, dir numg.V3) *intersect
}
