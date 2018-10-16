FROM golang:alpine AS caddy
WORKDIR /go/src/github.com/mholt/caddy
RUN echo https://alpine.global.ssl.fastly.net/alpine/v3.8/main > /etc/apk/repositories && \
    apk add --progress --update git gcc musl-dev ca-certificates
RUN git clone --branch v0.11.0 --single-branch --depth 1 https://github.com/mholt/caddy /go/src/github.com/mholt/caddy &> /dev/null && \
	sed -i 's/var EnableTelemetry = true/var EnableTelemetry = false/' /go/src/github.com/mholt/caddy/caddy/caddymain/run.go && \
    sed -i 's/const enableTelemetry = true/const enableTelemetry = false/' /go/src/github.com/mholt/caddy/caddy/caddymain/run.go
RUN git clone https://github.com/caddyserver/builds /go/src/github.com/caddyserver/builds
RUN cd caddy && go run build.go -goos=linux -goarch=amd64

FROM alpine:3.8 AS stackedit
RUN echo https://alpine.global.ssl.fastly.net/alpine/v3.8/main > /etc/apk/repositories && \
    apk add -q --progress --update --no-cache git npm
RUN git clone --branch v5.13.2 --single-branch --depth 1 https://github.com/benweet/stackedit.git /stackedit &> /dev/null
WORKDIR /stackedit
RUN npm install
RUN npm run build

FROM alpine:3.8
LABEL maintainer="quentin.mcgaw@gmail.com" \
      description="StackEdit server in a lightweight Docker container" \
      download="???MB" \
      size="51.2MB" \
      ram="2MB-3MB" \
      cpu_usage="Very low" \
      github="https://github.com/qdm12/stackedit-docker"
EXPOSE 80
COPY --from=stackedit /stackedit/dist /stackedit
COPY --from=caddy /go/src/github.com/mholt/caddy/caddy/caddy /usr/bin/caddy
COPY Caddyfile /Caddyfile
ENTRYPOINT caddy -conf /Caddyfile -log stdout
