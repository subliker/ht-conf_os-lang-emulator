package shell

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var (
	ErrDuringScript = errors.New("error during running script")
)

func (sh *sh) RunScriptFile(fpath string) error {
	if !sh.init {
		return ErrShellIsNotInit
	}

	sh.o.WriteString("Running script " + fpath + "...\n")

	f, err := os.Open(fpath)
	if err != nil {
		return fmt.Errorf("error opening script file(%s): %w", fpath, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if err := sh.RunStringCmnd(scanner.Text()); err != nil {
			sh.o.WriteString("Script error! " + err.Error())
			return ErrDuringScript
		}
		sh.o.WriteString("\n")
	}

	return nil
}
