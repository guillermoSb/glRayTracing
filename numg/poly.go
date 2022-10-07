package numg

import (
	"math"
)

const EQN_EPS = 1e-9

func IsZero(num float64) bool {
	return (num) > -EQN_EPS && num < EQN_EPS
}

func Solve2(c []float64) []float64 {
	p := c[1] / (2 * c[2])
	q := c[0] / c[2]

	D := p*p - q

	if IsZero(D) {
		return []float64{-p}
	} else if D < 0 {
		return []float64{}
	} else {
		sqrt_D := math.Sqrt(D)
		return []float64{sqrt_D - p, -sqrt_D - p}
	}
}

func Solve3(c []float64) []float64 {
	/* normal form: x^3 + Ax^2 + Bx + C = 0 */
	A := c[2] / c[3]
	B := c[1] / c[3]
	C := c[0] / c[3]
	sq_A := math.Pow(A, 2)
	p := 1.0 / 3.0 * (-1.0/3.0*sq_A + B)
	q := 1.0 / 2.0 * (2.0/27.0*A*sq_A - 1.0/3*A*B + C)
	cb_p := p * p * p
	D := q*q + cb_p
	s := []float64{}
	if IsZero(D) {
		if IsZero(q) {
			s = append(s, 0)
		} else {
			u := math.Cbrt(-q)
			s = append(s, 2.0*u)
			s = append(s, -u)
		}
	} else if D < 0 {
		phi := 1.0 / 3.0 * math.Acos(-q/math.Sqrt(-cb_p))
		t := 2 * math.Sqrt(-p)
		s = append(s, t*math.Cos(phi))
		s = append(s, -t*math.Cos(phi+math.Pi/3.0))
		s = append(s, -t*math.Cos(phi-math.Pi/3.0))

	} else {
		sqrt_D := math.Sqrt(D)
		u := math.Cbrt(sqrt_D - q)
		v := -math.Cbrt(sqrt_D + q)

		s = append(s, u+v)
	}
	sub := 1.0 / 3 * A
	for i := 0; i < len(s); i++ {
		s[i] -= sub
	}
	return s
}

func Solve4(c []float64) []float64 {
	/* normal form: x^4 + Ax^3 + Bx^2 + Cx + D = 0 */
	A := (c[3] / c[4])
	B := (c[2] / c[4])
	C := (c[1] / c[4])
	D := (c[0] / c[4])
	sq_A := math.Pow(A, 2)
	p := -3.0/8.0*sq_A + B
	q := 1.0/8.0*sq_A*A - 1.0/2.0*A*B + C
	r := -3.0/256.0*math.Pow(sq_A, 2) + 1.0/16.0*sq_A*B - 1.0/4.0*A*C + D
	s := []float64{}
	if IsZero(r) {
		coeffs := []float64{q, p, 0, 1}
		s = Solve3(coeffs)
		s = append(s, 0)
	} else {
		coeffs := []float64{
			1.0/2.0*r*p - 1.0/8.0*math.Pow(q, 2),
			-r,
			(-1.0 / 2.0) * p,
			1}
		s = Solve3(coeffs)
		z := s[0]
		u := z*z - r
		v := 2.0*z - p
		if IsZero(u) {
			u = 0
		} else if u > 0 {
			u = math.Sqrt(u)
		} else {
			s = []float64{}
			return s
		}

		if IsZero(v) {
			v = 0
		} else if v > 0 {
			v = math.Sqrt(v)
		} else {
			s = []float64{}
			return s
		}
		g := 0.0
		if q < 0 {
			g = -v
		} else {
			g = v
		}
		coeffs2 := []float64{}
		coeffs2 = append(coeffs2, z-u)
		coeffs2 = append(coeffs2, g)
		coeffs2 = append(coeffs2, 1)
		s = Solve2(coeffs2)

		coeffs3 := []float64{}
		g2 := 0.0
		if q < 0 {
			g2 = v
		} else {
			g = -v
		}
		coeffs3 = append(coeffs3, z+u)
		coeffs3 = append(coeffs3, g2)
		coeffs3 = append(coeffs3, 1)
		s = append(s, Solve2(coeffs3)...)

	}
	sub := 1.0 / 4.0 * A
	for i := 0; i < len(s); i++ {
		s[i] -= sub
	}
	return s
}
