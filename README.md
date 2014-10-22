dev
===

A small tool for provisioning development Docker containers.

```
dev version 0.1

Usage: dev [command] <manifest>

  if manifest is undefined the default value
  .dev.yaml will be used.

Available commands:
  DOWN  Destroys a development container
  UP    Brings up a development container
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
```

## Usage

Usage is simple:

```console
$ dev up
Starting up container for spike
spike-dev (43c5c1) running!
To use this container please attach to it with:
    $ docker attach spike-dev
$ docker attach spike-dev
docker:dev:spike ~
-->
```

```console
$ dev down
Container destroyed.
$
```

## Installation

This is a tool in Moonscript and thus requires you have the moonscript binary 
installed. No other dependencies are required other than the `docker` client 
binary.

Copy `dev.moon` to a place that is in your `PATH` environment variable. 
I personally get away with the following:

```console
$ cp dev.moon ~/bin/dev
```

---

This tool Works For Me and patches/feature/problem requests are welcome in the 
issue tracker. Please keep any patches simple.
