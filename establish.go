package main

import (
	"archive/tar"
	"bytes"
	"log"
	"os"
	"time"

	"github.com/fsouza/go-dockerclient"
)

func establishCommand() {
	man, err := loadManifest(*manifest)
	if err != nil {
		log.Fatal(err)
	}

	c, err := newDockerClient()
	if err != nil {
		log.Fatal(err)
	}

	if man.Overlay == "" {
		log.Fatal("You need an overlay to use this command")
	}

	inbuf := bytes.NewBuffer(nil)
	t := time.Now()

	tr := tar.NewWriter(inbuf)
	tr.WriteHeader(&tar.Header{
		Name:       "Dockerfile",
		Size:       int64(len(man.Overlay)),
		ModTime:    t,
		AccessTime: t,
		ChangeTime: t,
	})
	tr.Write([]byte(man.Overlay))
	tr.Close()

	opts := docker.BuildImageOptions{
		Name:         "dev-" + man.Projname,
		InputStream:  inbuf,
		OutputStream: os.Stdout,
	}

	err = c.BuildImage(opts)
	if err != nil {
		log.Fatal(err)
	}
}
