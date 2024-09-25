package shell

func (sh *sh) Exit() {
	sh.o.WriteString("\033[36mGoodbye! Comeback soon!\033[0m")
	if err := sh.fs.WriteZip(); err != nil {
		sh.o.WriteString("error writing zip: " + err.Error())
	}
	sh.fs.Clear()
	sh.init = false
}
