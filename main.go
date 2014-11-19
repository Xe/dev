package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v1"
)

var (
	app = kingpin.New("dev", "a simple tool to make development quicker")

	manifest = app.Flag("mainfest", "the manifest to create and spawn images from").Default(".dev.yaml").String()

	downcommand      = app.Command("down", "destroys a development container")
	establishcommand = app.Command("establish", "creates the backing image for a manifest with an overlay")
	initcommand      = app.Command("init", "create a new .dev.yaml file")
	purgecommand     = app.Command("purge", "deletes an established image")
	upcommand        = app.Command("up", "bring a development container up")
)

func main() {
	command := kingpin.MustParse(app.Parse(os.Args[1:]))

	switch command {
	case downcommand.FullCommand():
		downCommand()

	case establishcommand.FullCommand():
		establishCommand()

	case initcommand.FullCommand():
		initCommand()

	case purgecommand.FullCommand():
		purgeCommand()

	case upcommand.FullCommand():
		upCommand()

	default:
		app.Usage(os.Stderr)

		os.Exit(1)
	}
}
