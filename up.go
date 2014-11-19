package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fsouza/go-dockerclient"
)

func upCommand() {
	man, err := loadManifest(*manifest)
	if err != nil {
		log.Fatal(err)
	}

	c, err := newDockerClient()
	if err != nil {
		log.Fatal(err)
	}

	image := man.Base
	if image == "" {
		if man.Overlay != "" {
			image = "dev-" + man.Projname
		} else {
			log.Fatal("no image to run")
		}
	}

	path := "/home/" + man.User + "/dev/" + man.Projname
	if man.Golang {
		path = "/home/" + man.User + "/go/src/" + man.Repopath
	}

	volumes := make(map[string]struct{})
	binds := []string{}

	volumes[path] = struct{}{}
	binds = append(binds, os.Getenv("PWD")+":"+path)

	if man.Ssh {
		volumes["/home/"+man.User+"/.ssh"] = struct{}{}
		binds = append(binds, os.Getenv("HOME")+"/.ssh/:/home/"+man.User+"/.ssh")
	}

	opts := docker.CreateContainerOptions{
		Name: man.Projname + "-dev",
		Config: &docker.Config{
			Hostname:     "dev:" + man.Projname,
			Tty:          true,
			Image:        image,
			AttachStdin:  true,
			AttachStdout: true,
			AttachStderr: true,
			OpenStdin:    true,
			Volumes:      volumes,
		},
	}

	container, err := c.CreateContainer(opts)
	if err != nil {
		log.Fatal(err)
	}

	err = c.StartContainer(container.ID, &docker.HostConfig{
		Binds: binds,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Container %s-dev is running!\n", man.Projname)
	fmt.Printf("To use this container please attach to it with:\n")
	fmt.Printf("    $ docker attach %s-dev\n", man.Projname)
}
