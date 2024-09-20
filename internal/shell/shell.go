package shell

import (
	"context"
	"errors"

	"github.com/subliker/ht-conf_os-lang-emulator/internal/input"
	"github.com/subliker/ht-conf_os-lang-emulator/internal/output"
	"github.com/urfave/cli/v3"
)

type (
	sh struct {
		i    input.CLIInput
		o    output.CLIOutput
		sf   ShellFlags
		fs   fileSystem
		init bool
	}
	ShellFlags struct {
		Username  string
		PcName    string
		APath     string
		StartPath string
	}
)

var (
	ErrShellIsNotInit = errors.New("shell wasn't initialized")
)

func newShell(sf ShellFlags) sh {
	fs, err := newFS(sf.APath)
	if err != nil {
		panic("error creating filesystem: " + err.Error())
	}

	return sh{
		i:    input.NewCLIInput(),
		o:    output.NewCLIOutput(output.InputPromptData{Username: sf.Username, PcName: sf.PcName}),
		sf:   sf,
		fs:   fs,
		init: true,
	}
}

func RunShell(ctx context.Context, c *cli.Command, sf ShellFlags) error {
	sh := newShell(sf)

	sh.o.Clear()

	if sh.sf.StartPath != "" {
		if err := sh.RunScriptFile(sh.sf.StartPath); err != nil {
			print(err.Error())
		}
	}

	for sh.init {
		sh.o.WriteInputPrompt(sh.fs.GetCurPathString())
		cmnd, err := sh.i.ReadCmnd()
		if err != nil {
			sh.o.WriteString("Error reading command: " + err.Error())
			continue
		}

		if err := sh.RunStringCmnd(cmnd); err != nil {
			continue
		}

	}

	return nil
}

func (sh *sh) RunStringCmnd(cmnd string) error {
	if !sh.init {
		return ErrShellIsNotInit
	}

	pcmnd, err := input.ParseCmnd(cmnd)
	if err != nil {
		return err
	}

	if len(cmnd) == 0 {
		return nil
	}

	switch pcmnd[0] {
	case "exit":
		sh.o.WriteString("\033[36mGoodbye! Comeback soon!\033[0m")
		sh.init = false
		return nil
	default:
		sh.o.WriteString("Command " + pcmnd[0] + " wasn't found")
		return nil
	}
}
