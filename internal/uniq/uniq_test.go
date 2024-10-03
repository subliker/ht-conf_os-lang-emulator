package uniq

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subliker/ht-conf_os-lang-emulator/internal/fs/mocks"
)

var fsmock = &mocks.FileSystem{}

func TestUniq(t *testing.T) {
	assert := assert.New(t)

	tmpFile, err := os.Create("test1.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	err = os.WriteFile(tmpFile.Name(),
		[]byte(`Hello World!
Hello World!
Hello World!
Hello World!!!
Hello World!
Hello World!
Why?..`),
		os.ModePerm,
	)
	if err != nil {
		t.Fatal(err)
	}

	expectedOut :=
		`Hello World!
Hello World!!!
Hello World!
Why?..`

	out := ""
	pcmnd := []string{"uniq", tmpFile.Name()}
	fsmock.On("OpenFile", tmpFile.Name()).Return(tmpFile, nil)
	write := func(s string) { out += s }

	assert.Nil(Run(pcmnd, write, fsmock))
	assert.Equal(expectedOut, out)
	fsmock.AssertCalled(t, "OpenFile", tmpFile.Name())
}

func TestUniqWithCounter(t *testing.T) {
	assert := assert.New(t)

	tmpFile, err := os.Create("test2.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	err = os.WriteFile(tmpFile.Name(),
		[]byte(`Hello World!
Hello World!
Hello World!
Hello World!!!
Hello World!
Hello World!
Why?..`),
		os.ModePerm,
	)
	if err != nil {
		t.Fatal(err)
	}

	expectedOut :=
		`3 Hello World!
1 Hello World!!!
2 Hello World!
1 Why?..`

	out := ""
	pcmnd := []string{"uniq", "-c", tmpFile.Name()}
	fsmock.On("OpenFile", tmpFile.Name()).Return(tmpFile, nil)
	write := func(s string) { out += s }

	assert.Nil(Run(pcmnd, write, fsmock))
	assert.Equal(expectedOut, out)
	fsmock.AssertCalled(t, "OpenFile", tmpFile.Name())
}

func TestUniqEmptyFile(t *testing.T) {
	assert := assert.New(t)

	tmpFile, err := os.Create("test3.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	err = os.WriteFile(tmpFile.Name(),
		[]byte(``),
		os.ModePerm,
	)
	if err != nil {
		t.Fatal(err)
	}

	expectedOut := ``

	out := ""
	pcmnd := []string{"uniq", "-c", tmpFile.Name()}
	fsmock.On("OpenFile", tmpFile.Name()).Return(tmpFile, nil)
	write := func(s string) { out += s }

	assert.Nil(Run(pcmnd, write, fsmock))
	assert.Equal(expectedOut, out)
	fsmock.AssertCalled(t, "OpenFile", tmpFile.Name())
}

func TestUniqIncorrectSyntax(t *testing.T) {
	assert := assert.New(t)

	out := ""
	pcmnd := []string{"uniq"}
	write := func(s string) { out += s }

	assert.ErrorIs(Run(pcmnd, write, fsmock), ErrIncorrectSyntax)
	assert.Equal("incorrect syntax of command uniq", out)

	out = ""
	pcmnd = []string{"uniq", "-c"}

	assert.ErrorIs(Run(pcmnd, write, fsmock), ErrIncorrectSyntax)
	assert.Equal("filepath is not set", out)
}

func TestUniqErrorOpenFile(t *testing.T) {
	assert := assert.New(t)

	out := ""
	opErr := errors.New("directory a not found")
	pcmnd := []string{"uniq", "-c", "a/a.txt"}
	fsmock.On("OpenFile", "a/a.txt").Return(nil, opErr)
	write := func(s string) { out += s }

	assert.ErrorIs(Run(pcmnd, write, fsmock), ErrOpenFile)
	assert.NotEqual(out, "")
}
