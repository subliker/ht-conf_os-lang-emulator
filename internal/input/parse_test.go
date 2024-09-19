package input

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCmd(t *testing.T) {
	assert := assert.New(t)

	{
		r, err := parseCmnd("echo Hello World!")
		assert.Equal([]string{"echo", "Hello", "World!"}, r)
		assert.Nil(err)
	}

	{
		r, err := parseCmnd("echo 'Hello World!'")
		assert.Equal([]string{"echo", "'Hello World!'"}, r)
		assert.Nil(err)
	}

	{
		r, err := parseCmnd("echo 'Hello \"World!'")
		assert.Equal([]string(nil), r)
		assert.EqualError(err, ErrIncorrectCmnd.Error())
	}
}
