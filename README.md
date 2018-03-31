# StackEdit Docker server

Run a StackEdit v5.10.4 (April 2018) server in a Docker container.

[![Docker StackEdit](https://github.com/qdm12/stackedit-docker/raw/master/readme/title.png)](https://hub.docker.com/r/qmcgaw/stackedit/)

Docker build:
[![Build Status](https://travis-ci.org/qdm12/stackedit-docker.svg?branch=master)](https://travis-ci.org/qdm12/stackedit-docker)

Stackedit build:
[![Build Status](https://img.shields.io/travis/benweet/stackedit.svg?style=flat)](https://travis-ci.org/benweet/stackedit)

This image is **581 MB** and consumes **340MB-500MB** of RAM

It is based on:
- [Stackedit](https://github.com/benweet/stackedit)
- Alpine Linux
- Nodejs

## Features

- [Stackedit features](https://github.com/benweet/stackedit/blob/master/README.md#stackedit-can)

## Installation

### Option 1 of 2: Using Docker Compose

1. Download [**docker-compose.yml**](https://raw.githubusercontent.com/qdm12/stackedit-docker/master/docker-compose.yml)
1. Optionally edit *docker-compose.yml* to fit you better
1. With a terminal, go to the directory containing the file and launch 
the container in the background with:

    ```bash   
    docker-compose up -d
    ```

### Option 2 of 2: Using Docker only

In a terminal, enter:

```bash   
docker run -d --name=stackedit --restart=always -p 8080:8080/tcp qmcgaw/stackedit
```

The container TCP port 8080 is forwarded to the host TCP port 8080

## Testing

Go to [http://localhost:8080](http://localhost:8080)

## Acknowledgements

Credits to the [developers](https://github.com/benweet/stackedit/graphs/contributors) 
of [StackEdit](https://stackedit.io/)

