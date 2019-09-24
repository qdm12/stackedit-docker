# StackEdit Docker server

*StackEdit v5.14.0 (July 2019) with a Golang HTTP server on Scratch*

[![Docker StackEdit](https://github.com/qdm12/stackedit-docker/raw/master/readme/title.png)](https://hub.docker.com/r/qmcgaw/stackedit/)

[![Docker Build Status](https://img.shields.io/docker/build/qmcgaw/stackedit.svg)](https://hub.docker.com/r/qmcgaw/stackedit)

[![GitHub last commit](https://img.shields.io/github/last-commit/qdm12/stackedit-docker.svg)](https://github.com/qdm12/stackedit-docker/commits)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/y/qdm12/stackedit-docker.svg)](https://github.com/qdm12/stackedit-docker/commits)
[![GitHub issues](https://img.shields.io/github/issues/qdm12/stackedit-docker.svg)](https://github.com/qdm12/stackedit-docker/issues)

[![Docker Pulls](https://img.shields.io/docker/pulls/qmcgaw/stackedit.svg)](https://hub.docker.com/r/qmcgaw/stackedit)
[![Docker Stars](https://img.shields.io/docker/stars/qmcgaw/stackedit.svg)](https://hub.docker.com/r/qmcgaw/stackedit)
[![Docker Automated](https://img.shields.io/docker/automated/qmcgaw/stackedit.svg)](https://hub.docker.com/r/qmcgaw/stackedit)

[![Image size](https://images.microbadger.com/badges/image/qmcgaw/stackedit.svg)](https://microbadger.com/images/qmcgaw/stackedit)
[![Image version](https://images.microbadger.com/badges/version/qmcgaw/stackedit.svg)](https://microbadger.com/images/qmcgaw/stackedit)

[![Donate PayPal](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://paypal.me/qdm12)

| Image size | RAM usage | CPU usage |
| --- | --- | --- |
| 29.3MB | 7MB | Very low |

## Features

- [Stackedit features](https://github.com/benweet/stackedit/blob/master/README.md#stackedit-can)
- Lightweight image based on:
  - [Stackedit 5.14.0](https://github.com/benweet/stackedit)
  - [Scratch](https://hub.docker.com/_/scratch)
  - Golang simple HTTP static server
- Running without root
- Built-in Docker healthcheck
- Nice emojis in the logs...

## Setup

1. <details><summary>CLICK IF YOU HAVE AN ARM DEVICE</summary><p>

    - If you have a ARM 32 bit v6 architecture

        ```sh
        docker build -t qmcgaw/REPONAME_DOCKER \
        --build-arg BASE_IMAGE_BUILDER_GO=arm32v6/golang \
        --build-arg BASE_IMAGE_BUILDER_NODE=arm32v6/alpine \
        --build-arg GOARCH=arm \
        --build-arg GOARM=6 \
        https://github.com/qdm12/stackedit-docker.git
        ```

    - If you have a ARM 32 bit v7 architecture

        ```sh
        docker build -t qmcgaw/REPONAME_DOCKER \
        --build-arg BASE_IMAGE_BUILDER_GO=arm32v7/golang \
        --build-arg BASE_IMAGE_BUILDER_NODE=arm32v7/alpine \
        --build-arg GOARCH=arm \
        --build-arg GOARM=7 \
        https://github.com/qdm12/stackedit-docker.git
        ```

    - If you have a ARM 64 bit v8 architecture

        ```sh
        docker build -t qmcgaw/REPONAME_DOCKER \
        --build-arg BASE_IMAGE_BUILDER_GO=arm64v8/golang \
        --build-arg BASE_IMAGE_BUILDER_NODE=arm64v8/alpine \
        --build-arg GOARCH=arm64 \
        https://github.com/qdm12/stackedit-docker.git
        ```

    </p></details>

1. Use the following command:

    ```sh
    docker run -d -p 8000:8000/tcp qmcgaw/stackedit
    ```

    You can also use [docker-compose.yml](https://github.com/qdm12/stackedit-docker/blob/master/docker-compose.yml) with:

    ```sh
    docker-compose up -d
    ```

1. Access at [http://localhost:8000](http://localhost:8000)

## Environment variables

- `LISTENINGPORT` to change the internal HTTP server listening port if you need to

## Acknowledgements

Credits to the [developers](https://github.com/benweet/stackedit/graphs/contributors) 
of [StackEdit](https://stackedit.io/)

## TODOs

- [ ] Configuration of Stackedit with env variables
