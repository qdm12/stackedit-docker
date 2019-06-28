# StackEdit Docker server

*StackEdit v5.13.3 (November 2018) with a Golang HTTP server on Scratch*

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
  - [Stackedit 5.13.2](https://github.com/benweet/stackedit)
  - [Scratch](https://hub.docker.com/_/scratch)
  - Golang simple HTTP static server
- Running without root
- Built-in Docker healthcheck
- Nice emojis in the logs...

## Setup

Run with:

```bash
docker run -d -p 8000:8000/tcp qmcgaw/stackedit
```

Compose with:

```yml
version: '3'
services:
  stackedit:
    image: qmcgaw/stackedit
    container_name: stackedit
    ports:
      - 8000:8000/tcp
    network_mode: bridge
```

and `docker-compose up -d`

Build with:

```bash
docker build -t qmcgaw/stackedit https://github.com/qdm12/stackedit-docker.git
```

Access with [http://localhost:8000](http://localhost:8000)

## Environment variables

- `LISTENINGPORT` to change the internal HTTP server listening port if you need to

## Acknowledgements

Credits to the [developers](https://github.com/benweet/stackedit/graphs/contributors) 
of [StackEdit](https://stackedit.io/)

## TODOs

- [ ] Configuration of Stackedit with env variables
