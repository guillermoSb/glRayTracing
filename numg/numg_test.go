package numg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_can_solve_quadratic_poly(t *testing.T) {
	// Arrange
	c := []float64{-16.0, 0, 1.0}
	// Act
	sut := Solve2(c)
	// Assert
	assert.Equal(t, sut[0], 4.0)
	assert.Equal(t, sut[1], -4.0)
}

func Test_can_solve_cubic_poly(t *testing.T) {
	// Arrange
	c := []float64{2, -3, 0, 1}
	// Act
	sut := Solve3(c)
	fmt.Println(sut)
	// Assert
	assert.Equal(t, sut[0], -2.0)
	assert.Equal(t, sut[1], 1.0)
}

func Test_can_solve_dquad_poly(t *testing.T) {
	// Arrange
	c := []float64{-1, 1, 0, 1, 1}
	// Act
	sut := Solve4(c)
	fmt.Print("sut", sut)
	// Assert
	// assert.Equal(t, sut[0], -2.0)
	// assert.Equal(t, sut[1], 1.0)
}
