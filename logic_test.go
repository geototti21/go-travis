package travis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	app := New("my-name", 3)
	assert.Equal(t, &App{Name: "my-name", Version: 3}, app)
}


func TestNewV1(t *testing.T) {
	app := NewV1("my-name")
	assert.Equal(t, &App{Name: "my-name", Version: 1}, app)
}
