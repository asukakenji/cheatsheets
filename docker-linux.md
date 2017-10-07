# Linux in Docker

## Start an interactive shell using the `ubuntu` image

`<containerName>` is the name of the container:

    docker run -it --name <containerName> ubuntu /bin/bash

`-i`, `--interactive`: Keep STDIN open even if not attached
`-t`, `--tty`: Allocate a pseudo-TTY
`--name string`: Assign a name to the container

Before running `apt-get install`:

    apt-get -qq update

## Copy file from container to host

    docker cp <containerIdOrName>:/file/path/within/container /host/path/target
