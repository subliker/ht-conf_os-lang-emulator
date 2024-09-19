package shell

import (
	"context"

	"github.com/subliker/ht-conf_os-lang-emulator/internal/input"
	"github.com/subliker/ht-conf_os-lang-emulator/internal/output"
	"github.com/urfave/cli/v3"
)

type (
	sh struct {
		i       input.CLIInput
		o       output.CLIOutput
		curPath string
		sf      ShellFlags
	}
	ShellFlags struct {
		Username  string
		PcName    string
		APath     string
		StartPath string
	}
)

func newShell(sf ShellFlags) sh {
	return sh{input.NewCLIInput(), output.NewCLIOutput(output.InputPromptData{Username: sf.Username, PcName: sf.PcName}), "/", sf}
}

func RunShell(ctx context.Context, c *cli.Command, sf ShellFlags) error {
	sh := newShell(sf)

	sh.o.Clear()

	for {
		sh.o.WriteInputPrompt(sh.curPath)
		cmnd, err := sh.i.ReadAndParseCmnd()
		if err != nil {
			sh.o.WriteString("Error reading command: " + err.Error())
			continue
		}

		if len(cmnd) == 0 {
			continue
		}

		if cmnd[0] == "exit" {
			sh.o.WriteString("\033[36mGoodbye! Comeback soon!\033[0m")
			break
		} else {
			sh.o.WriteString("Command " + cmnd[0] + " wasn't found")
		}
	}

	return nil
}
