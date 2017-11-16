# Linux in Docker

## Start an interactive shell using the `ubuntu` image

Start a container using the image `ubuntu` and name it `try_ubuntu`, start an interactive shell in it:

    docker run -it --name try_ubuntu ubuntu /bin/bash

`-i`, `--interactive`: Keep STDIN open even if not attached
`-t`, `--tty`: Allocate a pseudo-TTY
`--name string`: Assign a name to the container

Before running `apt-get install`:

    apt-get -qq update

## Copy a file from container to host

    docker cp <containerIdOrName>:/file/path/within/container /host/path/target

## Running Docker inside a Docker container

Running Docker inside Docker is not suggested. Doing so will cause problems, as mentioned in this post:

https://stackoverflow.com/q/27879713/142239

Instead, the easiest way to archive similar functionality is to map the UNIX socket of the host to the container. Then, when Docker is executed inside the container, sibling containers will be created instead of child containers.

    docker run -it --name try_ubuntu -v /var/run/docker.sock:/var/run/docker.sock ubuntu /bin/bash

`-v`, `--volume list`: Bind mount a volume (`<hostPath>:<containerPath>`)
