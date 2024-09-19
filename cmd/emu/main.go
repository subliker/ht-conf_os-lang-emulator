package main

import (
	"context"
	"log"
	"os"

	"github.com/subliker/ht-conf_os-lang-emulator/internal/shell"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "emu",
		Usage: "emulator for the OS shell language",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "username",
				Value: "user",
				Usage: "username for shell user",
			},
			&cli.StringFlag{
				Name:  "pcname",
				Value: "pc",
				Usage: "pcname for shell pc",
			},
			&cli.StringFlag{
				Name:  "apath",
				Value: "fs",
				Usage: "path to the archive of the virtual file system",
			},
			&cli.StringFlag{
				Name:  "startpath",
				Value: "",
				Usage: "path to the start script",
			},
		},
		Action: shell.RunShell,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
