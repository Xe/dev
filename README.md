dev
===

A small tool for provisioning development Docker containers.

```
dev version 0.2

Usage: dev [command] <manifest>

  if manifest is undefined the default value
  .dev.yaml will be used.

Available commands:
        down   Destroys a development container
   establish   Create a Docker image from the manifest
          up   Brings up a development container
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
spike-dev (ctid: 63377a81e6b6) running!
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

```console
$ dev establish
Sending build context to Docker daemon  2.56 kB
Step 0 : FROM xena/dev-moonscript
 ---> 87f5e995d998
Step 1 : RUN echo "Hi mom!"
 ---> Using cache
 ---> a024c7c26b61
Successfully built a024c7c26b61
Docker image dev-spike created.
```

## Installation

### Install moonscript

Ubuntu:

```console
$ sudo apt-get install luarocks
$ sudo luarocks install moonrocks --server=http://rocks.moonscript.org
$ sudo moonrocks install yaml
$ sudo moonrocks install moonscript
```

Or use `make deps`

```console
$ make deps
```

No other dependencies are required other than the `docker` client binary.

### Tool Installation

Copy `dev.moon` to a place that is in your `PATH` environment variable. 
I personally get away with the following:

```console
$ cp dev.moon ~/bin/dev
```

Or if you are lazy:

```console
$ make install
```

---

This tool Works For Me™ and patches/feature/problem requests are welcome in the 
issue tracker. Please keep any patches simple.
