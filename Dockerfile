FROM alpine:3.8 AS stackedit
RUN echo https://alpine.global.ssl.fastly.net/alpine/v3.8/main > /etc/apk/repositories && \
    apk add -q --progress --update --no-cache git npm
RUN git clone --branch v5.13.2 --single-branch --depth 1 https://github.com/benweet/stackedit.git /stackedit &> /dev/null
WORKDIR /stackedit
RUN npm install
RUN npm run build

FROM golang:alpine AS server
RUN apk --update add git build-base upx
WORKDIR /go/src/app
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-s -w" -installsuffix cgo -o server . && \
    upx -v --best --ultra-brute --overlay=strip server && \
    upx -t server

FROM scratch
LABEL maintainer="quentin.mcgaw@gmail.com" \
      description="StackEdit server in a lightweight Docker container" \
      download="???MB" \
      size="???MB" \
      ram="7MB" \
      cpu_usage="Very low" \
      github="https://github.com/qdm12/stackedit-docker"
EXPOSE 80
COPY --from=stackedit /stackedit/dist /
COPY --from=server /go/src/app/server /server
ENTRYPOINT [ "/server" ]
