package echo

import (
	"errors"
	"strings"

	"github.com/subliker/ht-conf_os-lang-emulator/internal/fs"
)

var (
	ErrNotEnoughArgs = errors.New("arguments count less 2")
	ErrWriteFile     = errors.New("error writing file")
)

func Run(pcmnd []string, write func(string), fs fs.FileSystem) error {
	if len(pcmnd) <= 1 {
		return ErrNotEnoughArgs
	}
	offset := 0

	e := len(pcmnd) > 2 && pcmnd[1] == "-e"
	if e {
		offset++
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
		if i <= offset {
			continue
		}
		if i != offset+1 {
			str += " "
		}
		str += s
	}

	if len(str) > 1 && str[0] == '"' && str[len(str)-1] == '"' {
		str = strings.Trim(str, "\"")
	} else if len(str) > 1 && str[0] == '\'' && str[len(str)-1] == '\'' {
		str = strings.Trim(str, "'")
	}

	if e {
		str = formattedString(str)
	}

	if outf == "" {
		write(str)
		return nil
	}

	if err := fs.WriteFile(outf, rw, str); err != nil {
		write("Error writing to file(" + outf + "): " + err.Error())
		return ErrWriteFile
	}
	return nil
}

func formattedString(old string) string {
	formattingCommands := map[string]string{
		"\\n": "\n",
		"\\t": "\t",
		"\\b": " ",
		"\\r": "\r",
		"\\f": "\f",
	}
	for k, v := range formattingCommands {
		old = strings.ReplaceAll(old, k, v)
	}
	return old
}
