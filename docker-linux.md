# Linux in Docker

## Start an interactive shell using the `ubuntu` image

Start a container using the image `ubuntu` and name it `try_ubuntu`, start an interactive shell in it:

    docker run -it --name try_ubuntu ubuntu /bin/bash

`-i`, `--interactive`: Keep STDIN open even if not attached
`-t`, `--tty`: Allocate a pseudo-TTY
`--name string`: Assign a name to the container

Before running `apt-get install`:

    apt-get -qq update

## Copy file from container to host

    docker cp <containerIdOrName>:/file/path/within/container /host/path/target
