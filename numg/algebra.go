package numg

import (
	"math"
)

// Two dimensional vector
// - X: x coordinate value
// - Y: y coordinate value
type V2 struct {
    X,Y float64
}

// Three dimensional vector
// - X: x coordinate value
// - Y: y coordinate value
// - Z: z coordinate value
type V3 struct {
    X,Y,Z float64
}


// Normalizes a 3D vector
// - A: the vector to normalize
func NormalizeV3(A V3) V3 {
	m := math.Sqrt(math.Pow(float64(A.X), 2) + math.Pow(float64(A.Y), 2) + math.Pow(float64(A.Z), 2))
	newA := V3{A.X/m,A.Y/m,A.Z/m}
	return newA
}

// Subtracts two vectors
func Subtract(A,B V3) V3{
	newV := V3{0,0,0}
	newV.X = A.X - B.X
	newV.Y = A.Y - B.Y
	newV.Z = A.Z - B.Z
	return newV
}