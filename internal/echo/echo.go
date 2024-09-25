package echo

import "github.com/subliker/ht-conf_os-lang-emulator/internal/fs"

func Run(pcmnd []string, write func(string), fs fs.FileSystem) {
	if len(pcmnd) <= 1 {
		return
	}

	outf := ""
	rw := false

	if len(pcmnd) > 1 && (pcmnd[len(pcmnd)-2] == ">" || pcmnd[len(pcmnd)-2] == ">>") {
		outf = pcmnd[len(pcmnd)-1]
		rw = pcmnd[len(pcmnd)-2] == ">"
		pcmnd = pcmnd[:len(pcmnd)-2]
	} else if len(pcmnd[len(pcmnd)-1]) > 1 && (pcmnd[len(pcmnd)-1][0] == '>' || string(pcmnd[len(pcmnd)-1][0])+string(pcmnd[len(pcmnd)-1][1]) == ">>") {
		outf = pcmnd[len(pcmnd)-1][1:]
		rw = pcmnd[len(pcmnd)-1][0] == '>'
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
		return
	}

	if err := fs.WriteFile(outf, rw, str); err != nil {
		write("Error writing to file(" + outf + "): " + err.Error())
		return
	}
}
