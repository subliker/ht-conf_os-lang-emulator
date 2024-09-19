package input

import (
	"bufio"
	"os"
)

type (
	Reader interface {
		ReadAndParseCmnd() ([]string, error)
	}

	bufioReader struct {
		s *bufio.Reader
	}
)

func NewInput() Reader {
	reader := bufio.NewReader(os.Stdin)
	br := bufioReader{reader}

	return Reader(br)
}

func (br bufioReader) ReadAndParseCmnd() ([]string, error) {
	t, err := br.s.ReadString('\n')
	if err != nil {
		return nil, err
	}

	return parseCmnd(t)
}
