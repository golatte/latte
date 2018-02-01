package main

import (
	"os"

	"github.com/a2sdl/latte/latte/cmd"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "Build a Go web application quickly"
	app.Version = "0.1.0"
	app.Commands = []cli.Command{
		cmd.NewCommand,
	}

	app.Run(os.Args)
}
