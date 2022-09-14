package numg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test_can_normalize_v3_vector(t *testing.T) {
	// Arrange
	sut := V3 {2.0,2.0,1.0}
	// Act
	sut = NormalizeV3(sut)	
	// Assert
	assert.Equal(t, 2.0/3.0, sut.X)
	assert.Equal(t, 2.0/3.0, sut.Y)
	assert.Equal(t, 1.0/3.0, sut.Z)
}

func Test_can_subtract_v3_vector(t *testing.T) {
	// Arrange
	a := V3{2,3,1}
	b := V3{2,3,2}
	// Act
	sut := Subtract(a,b)
	// Assert
	assert.Equal(t, 0.0, sut.X)
	assert.Equal(t, 0.0, sut.Y)
	assert.Equal(t, -1.0, sut.Z)
}

func Test_can_compute_dot_product(t *testing.T) {
	// Arrange
	a := V3{2,3,1}
	b := V3{2,3,2}
	// Act
	sut := V3DotProduct(a,b)
	// Assert
	assert.Equal(t, 15.0, sut)
	
}


func Test_can_get_v3_magnitude(t *testing.T) {
	// Arrange
	a := V3 {2,2,1}
	// Act
	sut := a.Magnitude()
	// Assert
	assert.Equal(t,3.0,sut)

}