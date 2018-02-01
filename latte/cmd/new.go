package cmd

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

var NewCommand = cli.Command{
	Name:    "new",
	Aliases: []string{"n"},
	Usage:   "Creates a new webapp",
	Flags: []cli.Flag{
		cli.BoolTFlag{
			Name:  "web, w",
			Usage: "Create an interactive web application",
		},
		cli.BoolFlag{
			Name:  "api, a",
			Usage: "Create an API based web application",
		},
	},
	Action: newWebAppCommand,
}

func newWebAppCommand(c *cli.Context) {
	if len(c.Args()) != 1 {
		fmt.Fprintf(os.Stderr, "You must specify a name for your webapp\n")
		os.Exit(1)
	}

	switch {
	case c.Bool("web") && c.Bool("api"):
		fmt.Fprintf(os.Stderr, "You must specify only one of web or api type application\n")
		os.Exit(1)
	case c.Bool("web"):
		createWebApp(c.Args()[0])
	case c.Bool("api"):
		createAPIApp(c.Args()[0])
	}
}

func createWebApp(appName string) {
	createWebAppDirsAndFiles()
}

func createWebAppDirsAndFiles() {
	createTemplates()
	createPublic()
	createAssets()
	createWebpack()
}

func createTemplates() {

}

func createPublic() {

}

func createAssets() {

}

func createWebpack() {

}

func createAPIApp(appName string) {

}
