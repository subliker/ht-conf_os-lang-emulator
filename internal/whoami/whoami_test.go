package whoami

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhoami(t *testing.T) {
	assert := assert.New(t)

	{
		username := "user"
		expected := "You are \033[32m" + username + "\033[0m. Did you forget?"

		out := ""
		write := func(s string) { out += s }

		assert.Nil(Run(username, write))
		assert.Equal(expected, out)
	}

	{
		username := ""
		expected := "You are... Empty? It's illegal!"

		out := ""
		write := func(s string) { out += s }

		assert.ErrorIs(Run(username, write), ErrEmptyUsername)
		assert.Equal(expected, out)
	}

}
