package output

import (
	"bufio"
	"os"
	"os/exec"
)

type (
	CLIOutput interface {
		Clear()
		WriteInputPrompt(string)
		WriteString(string)
	}
	output struct {
		w  *bufio.Writer
		ip InputPromptData
	}
	InputPromptData struct {
		Username string
		PcName   string
	}
)

func NewCLIOutput(sd InputPromptData) CLIOutput {
	writer := bufio.NewWriter(os.Stdout)
	o := output{writer, sd}

	return CLIOutput(o)
}

func (o output) Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	cmd = exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (o output) WriteString(s string) {
	o.w.WriteString(s)
	o.w.Flush()
}

func (o output) WriteInputPrompt(curPath string) {
	/*
		{username} at {pcName} in {curPath}
		$
	*/
	o.w.WriteString("\n\033[32m" + o.ip.Username + "\033[0m at \033[35m" + o.ip.PcName + "\033[0m in \033[33m" + curPath + "\033[0m\n$ ")
	o.w.Flush()
}
