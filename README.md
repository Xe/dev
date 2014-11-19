dev
===

A small tool for provisioning development Docker containers.

```
usage: dev [<flags>] <command> [<flags>] [<args> ...]

a simple tool to make development quicker

Flags:
  --help  Show help.
  --mainfest=".dev.yaml"
          the manifest to create and spawn images from

Commands:
  help [<command>]
    Show help for a command.

  down
    destroys a development container

  establish
    creates the backing image for a manifest with an overlay

  init
    create a new .dev.yaml file

  purge
    deletes an established image

  up
    bring a development container up
```

A manifest is just a super-simple `.yaml` file that describes the container to
provision.

```yaml
base:     xena/dev-moonscript     # image to launch Docker with
repopath: github.com/Xe/dev/spike # repo path for mounting $CURDIR
golang:   false                   # Go has a more opinionated package store
ssh:      true                    # pass through ssh keys?
user:     xena                    # user in the docker container
projname: spike                   # project name
workdir:  image                   # use the docker image for workdir (default is code)
```

An overlay image may also be defined. If you define an overlay image, you do
not need to specify the `base` image. An overlay does not get any extra files
other than a Dockerfile. Example:

```yaml
# Overlay is the dockerfile to overlay into this for `dev establish`
overlay: |
  FROM xena/dev-moonscript

  RUN echo "Hi mom!"
```

## Usage

Usage is simple:

```console
$ dev up
Starting up container for spike
Container spike-dev is running!
To use this container please attach to it with:
    $ docker attach spike-dev
$ docker attach spike-dev
docker:dev:spike ~
-->
```

```console
$ dev down
Container destroyed
$
```

```console
$ dev establish
Sending build context to Docker daemon  2.56 kB
Step 0 : FROM xena/dev-moonscript
 ---> 87f5e995d998
Step 1 : RUN echo "Hi mom!"
 ---> Using cache
 ---> a024c7c26b61
Successfully built a024c7c26b61
```

## Installation

```console
$ go get github.com/Xe/dev
$ dev help
```

---

This tool Works For Me™ and patches/feature/problem requests are welcome in the
issue tracker. Please keep any patches simple.
