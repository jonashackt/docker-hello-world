# docker-hello-world
Example hello world container showing how to use GitHub Container Registry


As Docker Inc introduced a rate-limiting https://www.docker.com/increase-rate-limits I began to bump [into problems like this](https://github.com/jonashackt/molecule-ansible-docker-aws/runs/1968417806?check_suite_focus=true) while running a simple `docker run hello-world` on GitHub Actions:

```
Unable to find image 'hello...se the limit by authenticating and upgrading: https://www.docker.com/increase-rate-limit.\nSee 'docker run --help'.\n"
```

Many [people started to migrate their Docker images](https://medium.com/faun/migrating-my-docker-images-to-the-github-container-registry-9f304ccf0aaa) to the new GitHub Container Registry, which is currently in public beta: https://docs.github.com/en/packages/guides/pushing-and-pulling-docker-images

And there are already many projects that are simply not available anymore on DockerHub - but on GitHub Container Registry (like https://hub.docker.com/r/oracle/graalvm-ce to https://github.com/orgs/graalvm/packages/container/package/graalvm-ce)

Well I thought why not crafting a simple and small `hello-world` image and publish it to GitHub Container Registry?!


### A simple Go executable

The [original hello-world image from Docker](https://github.com/docker-library/hello-world) also uses a small executable to print a text. I decided to go with golang to create a ultra-small executable myself. 

So there's [hello-world.go](hello-world.go):

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello from Docker on GitHub Container Registry!\nThis message shows that your installation appears to be working correctly.\n\nAs Docker Inc introduced rate-limiting in https://www.docker.com/increase-rate-limits\nwe simply need our own hello-world image on GitHub Container Registry.\n\nTo generate this message, Docker took the following steps:\n 1. The Docker client contacted the Docker daemon.\n 2. The Docker daemon pulled this \"hello-world\" image from the GitHub Container Registry.\n    (amd64)\n 3. The Docker daemon created a new container from that image which runs the\n    executable that produces the output you are currently reading.\n 4. The Docker daemon streamed that output to the Docker client, which sent it\n    to your terminal.\n\n")
}
```

Build it (you need to have go installed with like `brew install go`) with:

```shell
go build hello-world.go
```

This produces a `hello-world` executable you can simply run with `./hello world`.


### A Docker multistage Build for GO

As we only need to have Go runtime stuff present to build the binary, we should implement a Docker multi-stage build. Since the GO Docker image https://hub.docker.com/_/golang is quite huge:
```shell
$ docker images
golang                             latest                861b1afd1d13   7 days ago       862MB
```

Therefore let's split our [Dockerfile](Dockerfile) a bit:

```dockerfile
# We need a golang build environment first
FROM golang:1.16.0-alpine3.13

WORKDIR /go/src/app
ADD hello-world.go /go/src/app

RUN go build hello-world.go

# We use a Docker multi-stage build here in order that we only take the compiled go executable
FROM alpine:3.13

COPY --from=0 "/go/src/app/hello-world" hello-world

ENTRYPOINT ./hello-world
```

Now let's build and run our image:

```shell
$ docker build . --tag hello-world
$ docker run hello-world
Hello from Docker on GitHub Container Registry!
This message shows that your installation appears to be working correctly.

As Docker Inc introduced rate-limiting in https://www.docker.com/increase-rate-limits
we simply need our own hello-world image on GitHub Container Registry.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled this "hello-world" image from the GitHub Container Registry.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.
```

The resulting image is around `7.55MB` which should be small enough for our use cases.