package input

import (
	"bufio"
	"os"
)

type (
	CLIInput interface {
		ReadAndParseCmnd() ([]string, error)
	}

	input struct {
		r *bufio.Reader
	}
)

func NewCLIInput() CLIInput {
	reader := bufio.NewReader(os.Stdin)
	i := input{reader}

	return CLIInput(i)
}

func (i input) ReadAndParseCmnd() ([]string, error) {
	t, err := i.r.ReadString('\n')
	if err != nil {
		return nil, err
	}

	return parseCmnd(t)
}
