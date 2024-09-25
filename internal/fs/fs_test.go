package fs

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	assert := assert.New(t)

	fs := fileSystem{fd: filepath.Join("C:", "Temp")}
	expected := filepath.Join("C:", "Temp", "main", "a")

	assert.Equal(expected, fs.path("main/a"))
}
