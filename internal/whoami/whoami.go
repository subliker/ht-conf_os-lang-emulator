package whoami

func Whoami(username string, write func(string)) {
	write("You are \033[32m" + username + "\033[0m. Did you forget?")
}
