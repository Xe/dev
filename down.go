package main

import (
	"fmt"
	"log"

	"github.com/fsouza/go-dockerclient"
)

func downCommand() {
	man, err := loadManifest(*manifest)
	if err != nil {
		log.Fatal(err)
	}

	c, err := newDockerClient()
	if err != nil {
		log.Fatal(err)
	}

	err = c.RemoveContainer(docker.RemoveContainerOptions{
		ID:    man.Projname + "-dev",
		Force: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Container destroyed")
}
