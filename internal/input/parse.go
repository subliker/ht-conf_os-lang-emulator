package input

import (
	"errors"
)

var (
	ErrIncorrectCmnd = errors.New("incorrect command")
)

func ParseCmnd(s string) ([]string, error) {
	m := map[rune]bool{
		'\'': false,
		'"':  false,
		'`':  false,
	}
	w := ""
	res := []string{}
	for _, b := range s {
		switch b {
		case '\'', '"', '`':
			m[b] = !m[b]
		case ' ':
			if !m['\''] && !m['"'] && !m['`'] {
				res = append(res, w)
				w = ""
				continue
			}
		}
		if b != '\n' && b != '\r' {
			w += string(b)
		}
	}

	if m['\''] || m['"'] || m['`'] {
		return nil, ErrIncorrectCmnd
	}

	if w != "" {
		res = append(res, w)
	}

	return res, nil
}
