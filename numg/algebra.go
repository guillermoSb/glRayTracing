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


// Gets the magnitude of the vector
func (v3 *V3) Magnitude() float64 {
	return math.Sqrt(math.Pow(v3.X, 2) + math.Pow(v3.Y, 2) + math.Pow(v3.Z, 2))
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

// Adds two vectors
func Add(A,B V3) V3{
	newV := V3{0,0,0}
	newV.X = A.X + B.X
	newV.Y = A.Y + B.Y
	newV.Z = A.Z + B.Z
	return newV
}

// Multiply a vector by a scalar
func MultiplyVectorWithConstant(A V3, c float64) V3{
	newV := V3{0,0,0}
	newV.X = A.X * c
	newV.Y = A.Y * c
	newV.Z = A.Z * c
	return newV
}

// Obtain the dot product of two vectors
func V3DotProduct(A V3, B V3) float64 {
	return A.X * B.X + A.Y * B.Y + A.Z * B.Z
}


func ReflectionVector(D, N V3) V3 {
	reflection :=  MultiplyVectorWithConstant(N, 2 * V3DotProduct(N, D))
	reflection = Subtract(reflection, D)
	reflection = NormalizeV3(reflection)
	return reflection
}