package output

import (
	"bufio"
	"os"
	"os/exec"
)

type (
	CLIOutput interface {
		Clear()
	}
	output struct {
		w  *bufio.Writer
		ip InputPromptData
	}
	InputPromptData struct {
		username string
		pcName   string
	}
)

func NewCLIOutput(sd InputPromptData) CLIOutput {
	writer := bufio.NewWriter(os.Stdout)
	o := output{writer, sd}

	return CLIOutput(o)
}

func (output) Clear() {
	exec.Command("clear")
}

func (o output) WriteString(s string) {
	o.w.WriteString(s)
}

func (o output) WriteInputPrompt(curPath string) {
	/*
		{username} at {pcName} in {curPath}
		$
	*/
	o.w.WriteString("/n" + o.ip.username + " at " + o.ip.pcName + " in " + curPath + "\n$")
}
