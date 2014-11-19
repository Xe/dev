package main

import (
	"log"
	"os"

	"github.com/fsouza/go-dockerclient"
)

func newDockerClient() (c *docker.Client, err error) {
	if allowed := os.Getenv("DOCKER_TLS_VERIFY"); allowed != "" {
		path := os.Getenv("DOCKER_CERT_PATH")

		c, err = docker.NewTLSClient(os.Getenv("DOCKER_HOST"),
			path+"/cert.pem",
			path+"/key.pem",
			path+"/ca.pem",
		)
	} else {
		if host := os.Getenv("DOCKER_HOST"); host != "" {
			c, err = docker.NewClient(host)
		} else {
			c, err = docker.NewClient("unix:///var/run/docker.sock")
		}
	}

	if err != nil {
		log.Println("failure in newDockerClient()")
		log.Fatal(err)
	}

	err = c.Ping()
	if err != nil {
		log.Println("failure in test")
		log.Fatal(err)
	}

	return
}
