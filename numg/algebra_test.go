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