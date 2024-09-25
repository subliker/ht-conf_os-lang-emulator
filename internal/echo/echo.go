package echo

import (
	"errors"

	"github.com/subliker/ht-conf_os-lang-emulator/internal/fs"
)

var (
	ErrNotEnoughArgs = errors.New("arguments count less 2")
	ErrWritingFile   = errors.New("error writing file")
)

func Run(pcmnd []string, write func(string), fs fs.FileSystem) error {
	if len(pcmnd) <= 1 {
		return ErrNotEnoughArgs
	}

	outf := ""
	rw := false

	if len(pcmnd) > 1 && (pcmnd[len(pcmnd)-2] == ">" || pcmnd[len(pcmnd)-2] == ">>") {
		outf = pcmnd[len(pcmnd)-1]
		rw = pcmnd[len(pcmnd)-2] == ">"
		pcmnd = pcmnd[:len(pcmnd)-2]
	} else if len(pcmnd[len(pcmnd)-1]) > 1 && (pcmnd[len(pcmnd)-1][0] == '>' || string(pcmnd[len(pcmnd)-1][0])+string(pcmnd[len(pcmnd)-1][1]) == ">>") {
		outf = ""
		rw = pcmnd[len(pcmnd)-1][:2] != ">>"
		if rw {
			outf = pcmnd[len(pcmnd)-1][1:]
		} else {
			outf = pcmnd[len(pcmnd)-1][2:]
		}
		pcmnd = pcmnd[:len(pcmnd)-1]
	}

	str := ""

	for i, s := range pcmnd {
		if i == 0 {
			continue
		}
		if i != 1 {
			str += " "
		}
		str += s
	}

	if outf == "" {
		write(str)
		return nil
	}

	if err := fs.WriteFile(outf, rw, str); err != nil {
		write("Error writing to file(" + outf + "): " + err.Error())
		return ErrWritingFile
	}
	return nil
}
