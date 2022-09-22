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

func Fresnel(normal, dir V3, ior float64) float64 {
	cosi := Clamp(V3DotProduct(normal, dir),-1,1)
	etai := 1.0
	etat := ior
	if cosi > 0 {
		etai, etat = etat, etai
	}
	sint := etai / etat * math.Sqrt(math.Max(0, 1 - cosi * cosi))
	kr := 1.0; 
	if (sint < 1) { 
		cost := math.Sqrt(math.Max(0.0,1 - sint * sint))
		cosi = math.Abs(cosi)
		rs := ((etat * cosi) - (etai * cost)) / ((etat * cosi) + (etai * cost))
		rp := ((etai * cosi) - (etat * cost)) / ((etai * cosi) + (etat * cost))
		kr = (rs * rs + rp * rp) / 2.0
  } 
	return kr
}
// Vec3f refract(const Vec3f &I, const Vec3f &N, const float &ior) 
// { 
//     float cosi = clamp(-1, 1, dotProduct(I, N)); 
//     float etai = 1, etat = ior; 
//     Vec3f n = N; 
//     if (cosi < 0) { cosi = -cosi; } else { std::swap(etai, etat); n= -N; } 
//     float eta = etai / etat; 
//     float k = 1 - eta * eta * (1 - cosi * cosi); 
//     return k < 0 ? 0 : eta * I + (eta * cosi - sqrtf(k)) * n; 
// } 

func Refract(normal, dir V3, ior float64) V3 {
	cosi := Clamp(V3DotProduct(normal, dir), -1, 1)
	etai := 1.0
	etat := ior
	n := normal
	if cosi < 0 {cosi = -cosi } else { 
		etat, etai = etai, etat
		n =  MultiplyVectorWithConstant(n, -1)
	}
	eta := etai / etat;
	k := 1 - eta * eta * (1 - cosi * cosi);
	if k < 0 {
		return V3{0,0,0}
	} else {
		return Add(MultiplyVectorWithConstant(dir, eta),  MultiplyVectorWithConstant(normal, eta * cosi - math.Sqrt(k)))
	}
}

func Clamp(x,min,max float64) float64 {
	if x < min {
		x = min
	} else if x > max {
		x = max
	}
	return x
}


func ReflectionVector(D, N V3) V3 {
	reflection :=  MultiplyVectorWithConstant(N, 2 * V3DotProduct(N, D))
	reflection = Subtract(reflection, D)
	reflection = NormalizeV3(reflection)
	return reflection
}