# StackEdit Docker server

Run a StackEdit v5.13.2 (October 2018) server with Nginx in a lightweight Docker container

[![Docker StackEdit](https://github.com/qdm12/stackedit-docker/raw/master/readme/title.png)](https://hub.docker.com/r/qmcgaw/stackedit/)

Docker build:
[![Build Status](https://travis-ci.org/qdm12/stackedit-docker.svg?branch=master)](https://travis-ci.org/qdm12/stackedit-docker)
[![Docker Build Status](https://img.shields.io/docker/build/qmcgaw/stackedit.svg)](https://hub.docker.com/r/qmcgaw/stackedit)

Stackedit build:
[![Build Status](https://img.shields.io/travis/benweet/stackedit.svg?style=flat)](https://travis-ci.org/benweet/stackedit)

[![GitHub last commit](https://img.shields.io/github/last-commit/qdm12/stackedit-docker.svg)](https://github.com/qdm12/stackedit-docker/commits)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/y/qdm12/stackedit-docker.svg)](https://github.com/qdm12/stackedit-docker/commits)
[![GitHub issues](https://img.shields.io/github/issues/qdm12/stackedit-docker.svg)](https://github.com/qdm12/stackedit-docker/issues)

[![Docker Pulls](https://img.shields.io/docker/pulls/qmcgaw/stackedit.svg)](https://hub.docker.com/r/qmcgaw/stackedit)
[![Docker Stars](https://img.shields.io/docker/stars/qmcgaw/stackedit.svg)](https://hub.docker.com/r/qmcgaw/stackedit)
[![Docker Automated](https://img.shields.io/docker/automated/qmcgaw/stackedit.svg)](https://hub.docker.com/r/qmcgaw/stackedit)

[![](https://images.microbadger.com/badges/image/qmcgaw/stackedit.svg)](https://microbadger.com/images/qmcgaw/stackedit)
[![](https://images.microbadger.com/badges/version/qmcgaw/stackedit.svg)](https://microbadger.com/images/qmcgaw/stackedit)

| Download size | Image size | RAM usage | CPU usage |
| --- | --- | --- | --- |
| ???MB | 45.5MB | 7MB | Very low |

## Features

- [Stackedit features](https://github.com/benweet/stackedit/blob/master/README.md#stackedit-can)
- Lightweight image based on:
  - [Stackedit 5.13.2](https://github.com/benweet/stackedit)
  - Alpine 3.8
  - Nginx HTTP server

## Setup

Using plain Docker with:

```bash
docker run -d --name=stackedit --restart=always -p 8000:80/tcp qmcgaw/stackedit
```


Or use Docker Compose:

```yml
version: '3'
services:
  stackedit:
    image: qmcgaw/stackedit
    container_name: stackedit
    ports:
      - 8000:80/tcp
    network_mode: bridge
    restart: always
```


with the command


```bash
docker-compose up -d
```

## Testing

Go to [http://localhost:8000](http://localhost:8000)

## Acknowledgements

Credits to the [developers](https://github.com/benweet/stackedit/graphs/contributors) 
of [StackEdit](https://stackedit.io/)

## TODOs

- [ ] Configuration of Stackedit
