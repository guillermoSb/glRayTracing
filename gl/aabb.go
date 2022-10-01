package gl

import (
	"guillermoSb/glRayTracing/numg"
	"math"
)

// Axis Aligned Bounding Box
type aabb struct {
	material  material
	position  numg.V3
	size      numg.V3
	planes    []plane
	boundsMin numg.V3
	boundsMax numg.V3
}

func NewAABB(material material, position numg.V3, size numg.V3) *aabb {

	halfSizes := numg.V3{0, 0, 0}
	halfSizes.X = size.X / 2
	halfSizes.Y = size.Y / 2
	halfSizes.Z = size.Z / 2
	planes := []plane{}
	// Sides
	planes = append(planes, *NewPlane(material, numg.Add(position, numg.V3{halfSizes.X, 0, 0}), numg.V3{1, 0, 0}))
	planes = append(planes, *NewPlane(material, numg.Add(position, numg.V3{-halfSizes.X, 0, 0}), numg.V3{-1, 0, 0}))
	// Up an down
	planes = append(planes, *NewPlane(material, numg.Add(position, numg.V3{0, halfSizes.Y, 0}), numg.V3{0, 1, 0}))
	planes = append(planes, *NewPlane(material, numg.Add(position, numg.V3{0, -halfSizes.Y, 0}), numg.V3{0, -1, 0}))
	// Front and back
	planes = append(planes, *NewPlane(material, numg.Add(position, numg.V3{0, 0, halfSizes.Z}), numg.V3{0, 0, 1}))
	planes = append(planes, *NewPlane(material, numg.Add(position, numg.V3{0, 0, -halfSizes.Z}), numg.V3{0, 0, -1}))
	// Create the bounds for the box
	boundsMin := numg.V3{}
	boundsMax := numg.V3{}
	epsilon := 0.001

	boundsMin.X = position.X - (epsilon + halfSizes.X)
	boundsMax.X = position.X + (epsilon + halfSizes.X)

	boundsMin.Y = position.Y - (epsilon + halfSizes.Y)
	boundsMax.Y = position.Y + (epsilon + halfSizes.Y)

	boundsMin.Z = position.Z - (epsilon + halfSizes.Z)
	boundsMax.Z = position.Z + (epsilon + halfSizes.Z)

	aabb := aabb{material, position, size, planes, boundsMin, boundsMax}

	return &aabb
}

func (aabb *aabb) rayIntersect(origin, dir numg.V3) *intersect {
	t := math.Inf(1)
	var intersect *intersect = nil

	for _, plane := range aabb.planes {
		planeInter := plane.rayIntersect(origin, dir)

		if planeInter != nil {
			// Check if the intersection is inside the bound
			planePoint := planeInter.point
			if aabb.boundsMin.X <= planePoint.X && aabb.boundsMax.X >= planePoint.X {
				if aabb.boundsMin.Y <= planePoint.Y && aabb.boundsMax.Y >= planePoint.Y {
					if aabb.boundsMin.Z <= planePoint.Z && aabb.boundsMax.Z >= planePoint.Z {
						if planeInter.distance < t {
							t = planeInter.distance
							intersect = planeInter
						}
					}
				}

			}
		}
	}
	if intersect != nil {
		return NewIntersect(t, intersect.point, intersect.normal, struct{ material }{
			material: aabb.material,
		})
	}

	return nil
}
