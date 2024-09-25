package whoami

import "errors"

var (
	ErrEmptyUsername = errors.New("empty username")
)

func Run(username string, write func(string)) error {
	if username == "" {
		write("You are... Empty? It's illegal!")
		return ErrEmptyUsername
	}
	write("You are \033[32m" + username + "\033[0m. Did you forget?")
	return nil
}
