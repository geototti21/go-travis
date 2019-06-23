package travis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	app := New("my-name", 3)
	assert.Equal(t, &App{Name: "my-name", Version: 3}, app)
}
