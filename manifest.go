package main

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Manifest struct {
	Base     string
	Repopath string
	Golang   bool
	Ssh      bool
	User     string
	Projname string
	Overlay  string
}

func loadManifest(finname string) (m *Manifest, err error) {
	fin, err := os.Open(finname)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(fin)

	m = &Manifest{}

	err = yaml.Unmarshal(data, m)

	return
}
