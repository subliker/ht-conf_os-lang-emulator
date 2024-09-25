package echo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subliker/ht-conf_os-lang-emulator/internal/fs/mocks"
)

var fsmock = &mocks.FileSystem{}

func TestEchoOutput(t *testing.T) {
	assert := assert.New(t)

	out := ""
	pcmnd := []string{"echo", "Hello", "World!"}
	write := func(s string) { out += s }

	assert.Nil(Run(pcmnd, write, fsmock))
	assert.Equal(out, "Hello World!")
	fsmock.AssertNotCalled(t, "WriteFile")

}

func TestEchoAppendFile(t *testing.T) {
	assert := assert.New(t)

	out := ""
	fsmock.On("WriteFile", "a.txt", false, "Hello World!").Return(error(nil)).Once()
	pcmnd := []string{"echo", "Hello", "World!", ">>", "a.txt"}
	write := func(s string) { out += s }

	assert.Nil(Run(pcmnd, write, fsmock))
	assert.Equal(out, "")
	fsmock.AssertCalled(t, "WriteFile", "a.txt", false, "Hello World!")

}

func TestEchoAppendFileWithoutSpace(t *testing.T) {
	assert := assert.New(t)

	out := ""
	fsmock.On("WriteFile", "a.txt", false, "Hello World!").Return(error(nil)).Once()
	pcmnd := []string{"echo", "Hello", "World!", ">>a.txt"}
	write := func(s string) { out += s }

	assert.Nil(Run(pcmnd, write, fsmock))
	assert.Equal(out, "")
	fsmock.AssertCalled(t, "WriteFile", "a.txt", false, "Hello World!")
}

func TestEchoWriteFile(t *testing.T) {
	assert := assert.New(t)

	out := ""
	fsmock.On("WriteFile", "a.txt", true, "Hello World!").Return(error(nil)).Once()
	pcmnd := []string{"echo", "Hello", "World!", ">", "a.txt"}
	write := func(s string) { out += s }

	assert.Nil(Run(pcmnd, write, fsmock))
	assert.Equal(out, "")
	fsmock.AssertCalled(t, "WriteFile", "a.txt", true, "Hello World!")
}
func TestEchoWriteFileWithoutSpace(t *testing.T) {
	assert := assert.New(t)

	out := ""
	fsmock.On("WriteFile", "a.txt", true, "Hello World!").Return(error(nil)).Once()
	pcmnd := []string{"echo", "Hello", "World!", ">a.txt"}
	write := func(s string) { out += s }

	assert.Nil(Run(pcmnd, write, fsmock))
	assert.Equal(out, "")
	fsmock.AssertCalled(t, "WriteFile", "a.txt", true, "Hello World!")
}

func TestEchoWithLessArguments(t *testing.T) {
	assert := assert.New(t)

	out := ""
	pcmnd := []string{"echo"}
	write := func(s string) { out += s }

	assert.ErrorIs(Run(pcmnd, write, fsmock), ErrNotEnoughArgs)
	assert.Equal(out, "")
	fsmock.AssertNotCalled(t, "WriteFile")
}
func TestEchoWithWriteError(t *testing.T) {
	assert := assert.New(t)

	out := ""
	wrErr := errors.New("directory a not found")
	fsmock.On("WriteFile", "a/a.txt", true, "Hello World!").Return(wrErr).Once()
	pcmnd := []string{"echo", "Hello", "World!", ">", "a/a.txt"}
	write := func(s string) { out += s }

	assert.ErrorIs(Run(pcmnd, write, fsmock), ErrWriteFile)
	assert.Equal(out, "Error writing to file(a/a.txt): "+wrErr.Error())
	fsmock.AssertCalled(t, "WriteFile", "a/a.txt", true, "Hello World!")
}
