package uniq

import (
	"bufio"
	"errors"
	"strconv"

	"github.com/subliker/ht-conf_os-lang-emulator/internal/fs"
)

var (
	ErrIncorrectSyntax = errors.New("arguments count less 2")
	ErrOpenFile        = errors.New("error opening file")
)

func Run(pcmnd []string, write func(string), fs fs.FileSystem) error {
	if len(pcmnd) <= 1 {
		write("incorrect syntax of command uniq")
		return ErrIncorrectSyntax
	}

	ctr := pcmnd[1] == "-c"
	fpath := ""
	if ctr {
		if len(pcmnd) < 3 {
			write("filepath is not set")
			return ErrIncorrectSyntax
		}
		fpath = pcmnd[2]
	} else {
		fpath = pcmnd[1]
	}

	f, err := fs.OpenFile(fpath)
	if err != nil {
		write("error opening file " + fpath)
		return ErrOpenFile
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	if !scanner.Scan() {
		return nil
	}
	s := scanner.Text()
	c := 1

	for scanner.Scan() {
		if scanner.Text() == s {
			c++
		} else {
			if ctr {
				write(strconv.Itoa(c) + " ")
			}
			write(s + "\n")
			c = 1
			s = scanner.Text()
		}
	}
	if ctr {
		write(strconv.Itoa(c) + " ")
	}
	write(s + "\n")
	return nil
}
