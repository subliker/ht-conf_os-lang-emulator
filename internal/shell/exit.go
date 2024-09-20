package shell

func (sh *sh) Exit() {
	sh.o.WriteString("\033[36mGoodbye! Comeback soon!\033[0m")
	sh.init = false
}
