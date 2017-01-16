package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "open-food-data"
	app.Usage = `Get open food data`

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "enable debug output for logging",
		},
	}

	app.Before = func(context *cli.Context) error {
		if context.GlobalBool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}

		return nil
	}

	app.Commands = []cli.Command{
		getFoodDataCommand,
		getIndexCommand,
		getImageCommand,
		getCategoryCommand,
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Error(err)
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
