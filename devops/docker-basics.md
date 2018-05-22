# Docker Basics

Table of Contents

- [Containers](#containers)
- [Images](#images)

## Containers

### List running containers

    docker ps

### List all containers

    docker ps -a

### Run a container (the image will be downloaded if not yet so)

Start a container using the image `nginx` with host port `8080` mapped to container port `80` and name it `try_nginx`:

    docker run --name try_nginx -p 8080:80 nginx

Start a container using the image `ubuntu` and name it `try_ubuntu` (nothing happens since no command is given):

    docker run --name try_ubuntu ubuntu

### Run a container in background

Start a container in the background using the image `nginx` with host port `8080` mapped to container port `80` and name it `try_nginx`:

    docker run -d --name try_nginx -p 8080:80 nginx

> `-d`, `--detach`: Run container in background and print container ID

### Run a one-time command

Start a container using the image `ubuntu`, and delete it after the command `echo` finishes:

    docker run --rm ubuntu echo "Hello World"

> `--rm`: Automatically remove the container when it exits

### Run a container, override the entry point command

Start a container using the image `nginx`, start an interactive shell instead of running the HTTP server:

    docker run --rm -it --entrypoint /bin/bash nginx

> `--rm`: Automatically remove the container when it exits
>
> `-i`, `--interactive`: Keep STDIN open even if not attached
>
> `-t`, `--tty`: Allocate a pseudo-TTY
>
> `--entrypoint string`: Overwrite the default ENTRYPOINT of the image

### Start an interactive shell in a running container

Start an interactive shell in the running container `try_nginx`:

    docker exec -it try_nginx /bin/bash

### Run a command in a running container

Run the `ps` command in the running container `try_nginx`:

    docker exec try_nginx ps

### Start a stopped (an existing) container

Start the container `try_nginx`:

    docker start try_nginx

### Stop a running container

Stop the container `ac6aa090558c`:

    docker stop ac6aa090558c

### Remove a container

Remove the container `try_ubuntu`:

    docker rm try_ubuntu

## Images

## List downloaded images

    docker images

Alternative spelling:

    docker image ls

## Download (or update) an image

Download (or update) the image `ubuntu`:

    docker pull ubuntu

## Download (or update) a specific tag of an image

Download (or update) the `trusty` tag of the image `ubuntu`:

    docker pull ubuntu:trusty

## Get a list of supported tags

HTML result:

https://hub.docker.com/_/ubuntu/

JSON result:

https://registry.hub.docker.com/v1/repositories/ubuntu/tags

## Remove an image

Remove the image `ubuntu`:

    docker rmi ubuntu

Alternative spelling:

    docker image rm ubuntu
