package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	local = flag.Bool("local", true, "if set, dev will use the DOCKER_HOST envvar to determine where to open the Docker client")
)

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	subcommand := args[0]

	fmt.Println("Command: " + subcommand)
}
