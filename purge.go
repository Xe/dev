package main

import (
	"fmt"
	"log"
)

func purgeCommand() {
	man, err := loadManifest(*manifest)
	if err != nil {
		log.Fatal(err)
	}

	c, err := newDockerClient()
	if err != nil {
		log.Fatal(err)
	}

	err = c.RemoveImage("dev-" + man.Projname)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Image deleted")
}
