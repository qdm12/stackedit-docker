# StackEdit Docker server

*StackEdit v5.14.5 (January 2020) with a Golang HTTP server on Scratch*

[![Docker StackEdit](https://github.com/qdm12/stackedit-docker/raw/master/title.png)](https://hub.docker.com/r/qmcgaw/stackedit/)

[![Build status](https://github.com/qdm12/stackedit-docker/workflows/Buildx%20latest/badge.svg)](https://github.com/qdm12/stackedit-docker/actions?query=workflow%3A%22Buildx+latest%22)
[![Docker Pulls](https://img.shields.io/docker/pulls/qmcgaw/stackedit.svg)](https://hub.docker.com/r/qmcgaw/stackedit)
[![Docker Stars](https://img.shields.io/docker/stars/qmcgaw/stackedit.svg)](https://hub.docker.com/r/qmcgaw/stackedit)
[![Image size](https://images.microbadger.com/badges/image/qmcgaw/stackedit.svg)](https://microbadger.com/images/qmcgaw/stackedit)
[![Image version](https://images.microbadger.com/badges/version/qmcgaw/stackedit.svg)](https://microbadger.com/images/qmcgaw/stackedit)

[![GitHub last commit](https://img.shields.io/github/last-commit/qdm12/stackedit-docker.svg)](https://github.com/qdm12/stackedit-docker/commits)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/y/qdm12/stackedit-docker.svg)](https://github.com/qdm12/stackedit-docker/commits)
[![GitHub issues](https://img.shields.io/github/issues/qdm12/stackedit-docker.svg)](https://github.com/qdm12/stackedit-docker/issues)
[![Donate PayPal](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://paypal.me/qmcgaw)

## Features

- [Stackedit features](https://github.com/benweet/stackedit/blob/master/README.md#stackedit-can)
- Lightweight image based on:
  - [Stackedit 5.14.5](https://github.com/benweet/stackedit)
  - [Scratch](https://hub.docker.com/_/scratch)
  - Golang simple HTTP static server
- Running without root
- Cross cpu architecture compatible: amd64, 386, arm64v8, arm32v7 and arm32v6 (ask for more)
- Small 36.8MB image size (amd64, uncompressed)
- Built-in Docker healthcheck
- Nice emojis in the logs...

## Setup

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

| Environment variable | Default | Description |
| --- | --- | --- |
| `LISTENING_PORT` | `8000` | Internal server listening port |
| `ROOT_URL` | `/` | Root URL to use, useful when used with a reverse proxy |
| `NODE_ENV` | `production` | Sets production behavior for stackedit  |
| `PANDOC_PATH` | `pandoc` | *Non functional yet* |
| `WKHTMLTOPDF_PATH` | `wkhtmltopdf` | *Non functional yet* |
| `USER_BUCKET_NAME` | `stackedit-users` | ? |
| `PAYPAL_RECEIVER_EMAIL` |  | Receive Paypal donation email address |
| `DROPBOX_APP_KEY` | | |
| `DROPBOX_APP_KEY_FULL` | | |
| `GITHUB_CLIENT_ID` | | |
| `GITHUB_CLIENT_SECRET` | | |
| `GOOGLE_CLIENT_ID` | | |
| `GOOGLE_API_KEY` | | |
| `WORDPRESS_CLIENT_ID` | | |

## Acknowledgements

Credits to the [developers](https://github.com/benweet/stackedit/graphs/contributors) of [StackEdit](https://stackedit.io/)

## TODOs

- [ ] Add static binary programs
    - [ ] pandoc
    - [ ] wkhtmltopdf
- [ ] Travis CI build cross CPU arch
