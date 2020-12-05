package read_test

import (
	"testing"

	"github.com/joanlopez/golang-snippets/files/read"
	"github.com/stretchr/testify/assert"
)

func TestLines(t *testing.T) {
	lines, err := read.Lines("./sample")
	assert.NoError(t, err)
	assert.Equal(t, []string{"This is an example", "to show how to read", "files in Go."}, lines)
}

func TestWords(t *testing.T) {
	words, err := read.Words("./sample")
	assert.NoError(t, err)
	assert.Equal(t, []string{"This", "is", "an", "example", "to", "show", "how", "to", "read", "files", "in", "Go."}, words)
}

func TestRunes(t *testing.T) {
	runes, err := read.Runes("./sample")
	assert.NoError(t, err)
	assert.Equal(t, []rune{84, 104, 105, 115, 32, 105, 115, 32, 97, 110, 32, 101, 120, 97, 109, 112, 108, 101, 10, 116, 111, 32, 115, 104, 111, 119, 32, 104, 111, 119, 32, 116, 111, 32, 114, 101, 97, 100, 10, 102, 105, 108, 101, 115, 32, 105, 110, 32, 71, 111, 46}, runes)
}

func TestBytes(t *testing.T) {
	bytes, err := read.Bytes("./sample")
	assert.NoError(t, err)
	assert.Equal(t, []byte{0x54, 0x68, 0x69, 0x73, 0x20, 0x69, 0x73, 0x20, 0x61, 0x6e, 0x20, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0xa, 0x74, 0x6f, 0x20, 0x73, 0x68, 0x6f, 0x77, 0x20, 0x68, 0x6f, 0x77, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x61, 0x64, 0xa, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x20, 0x69, 0x6e, 0x20, 0x47, 0x6f, 0x2e}, bytes)
}
