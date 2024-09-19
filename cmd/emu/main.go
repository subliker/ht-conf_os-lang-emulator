package main

import (
	"context"
	"log"
	"os"

	"github.com/subliker/ht-conf_os-lang-emulator/internal/shell"
	"github.com/urfave/cli/v3"
)

type flags struct {
}

func main() {
	sf := shell.ShellFlags{}

	cmd := &cli.Command{
		Name:  "emu",
		Usage: "emulator for the OS shell language",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "username",
				Value:       "user",
				Usage:       "username for shell user",
				Destination: &sf.Username,
			},
			&cli.StringFlag{
				Name:        "pcname",
				Value:       "pc",
				Usage:       "pcname for shell pc",
				Destination: &sf.PcName,
			},
			&cli.StringFlag{
				Name:        "apath",
				Value:       "fs",
				Usage:       "path to the archive of the virtual file system",
				Destination: &sf.APath,
			},
			&cli.StringFlag{
				Name:        "startpath",
				Value:       "",
				Usage:       "path to the start script",
				Destination: &sf.StartPath,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			shell.RunShell(ctx, c, sf)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
