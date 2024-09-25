package uniq

import (
	"bufio"
	"strconv"

	"github.com/subliker/ht-conf_os-lang-emulator/internal/fs"
)

func Run(pcmnd []string, write func(string), fs fs.FileSystem) {
	if len(pcmnd) <= 1 {
		write("incorrect syntax of command uniq")
		return
	}

	ctr := pcmnd[1] == "-c"
	fpath := ""
	if ctr {
		if len(pcmnd) < 3 {
			write("filepath is not set")
			return
		}
		fpath = pcmnd[2]
	} else {
		fpath = pcmnd[1]
	}

	f, err := fs.OpenFile(fpath)
	if err != nil {
		write("error opening file " + fpath)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	if !scanner.Scan() {
		print(111)
		return
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
}
