package gl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)



func Test_can_create_color(t *testing.T) {
	// Arrange, Act
	c, err := NewColor(1,0,0)
	// Assert
	assert.NotNil(t,c)
	assert.Nil(t,err)
}

func Test_color_only_accepts_values_between_1_and_0(t *testing.T) {
	// Arrange, Act
	c,err := NewColor(10,10,10)

	assert.Nil(t,c, "Expected color to not exist")
	// Assert
	assert.NotNil(t,err)
}
