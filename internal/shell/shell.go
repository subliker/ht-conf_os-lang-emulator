package shell

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/subliker/ht-conf_os-lang-emulator/internal/echo"
	"github.com/subliker/ht-conf_os-lang-emulator/internal/fs"
	"github.com/subliker/ht-conf_os-lang-emulator/internal/input"
	"github.com/subliker/ht-conf_os-lang-emulator/internal/output"
	"github.com/subliker/ht-conf_os-lang-emulator/internal/uniq"
	"github.com/subliker/ht-conf_os-lang-emulator/internal/whoami"
	"github.com/urfave/cli/v3"
)

type (
	sh struct {
		i    input.CLIInput
		o    output.CLIOutput
		sf   ShellFlags
		fs   fs.FileSystem
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
	fs, err := fs.NewFS(sf.APath)
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

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigs
		sh.Exit()
		os.Exit(0)
	}()

	for sh.init {
		sh.o.WriteInputPrompt(sh.fs.CurPath())
		cmnd, err := sh.i.ReadCmnd()
		if err != nil {
			sh.o.WriteString("Error reading command: " + err.Error())
			continue
		}

		if err := sh.RunStringCmnd(cmnd); err != nil {
			sh.o.WriteString("Error parsing command: " + err.Error())
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
		sh.Exit()
	case "whoami":
		whoami.Run(sh.sf.Username, sh.o.WriteString)
	case "echo":
		echo.Run(pcmnd, sh.o.WriteString, sh.fs)
	case "mkdir":
		if len(pcmnd) != 2 {
			sh.o.WriteString("incorrect mkdir command")
			break
		}
		if err := sh.fs.MakeDirectory(pcmnd[1]); err != nil {
			sh.o.WriteString(err.Error())
		}
	case "cd":
		if len(pcmnd) != 2 {
			sh.o.WriteString("incorrect cd command")
			break
		}
		if err := sh.fs.ChangeDirectory(pcmnd[1]); err != nil {
			sh.o.WriteString(err.Error())
		}
	case "ls":
		sh.fs.List(sh.o.WriteString, len(pcmnd) > 1 && pcmnd[1] == "-l")
	case "uniq":
		uniq.Run(pcmnd, sh.o.WriteString, sh.fs)
	default:
		sh.o.WriteString("Command " + pcmnd[0] + " wasn't found")
	}

	return nil
}
