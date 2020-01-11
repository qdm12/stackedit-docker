ARG ALPINE_VERSION=3.10
ARG GO_VERSION=1.13.0
ARG STACKEDIT_VERSION=v5.14.5

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS server
RUN apk --update add git g++
WORKDIR /tmp/gobuild
COPY go.mod go.sum ./
RUN go mod download 2>&1
COPY main.go ./
#RUN go test -v -race ./...
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o app main.go

FROM alpine:${ALPINE_VERSION} AS stackedit
ARG STACKEDIT_VERSION
WORKDIR /stackedit
RUN apk add -q --progress --update --no-cache git npm
RUN wget -q https://github.com/benweet/stackedit/archive/${STACKEDIT_VERSION}.tar.gz -O stackedit.tar.gz && \
    tar -xzf stackedit.tar.gz --strip-components=1 && \
    rm stackedit.tar.gz
#ENV NODE_ENV production
RUN npm install
RUN npm run build

FROM scratch AS final
ARG BUILD_DATE
ARG VCS_REF
ARG STACKEDIT_VERSION
LABEL \
    org.opencontainers.image.authors="quentin.mcgaw@gmail.com" \
    org.opencontainers.image.created=$BUILD_DATE \
    org.opencontainers.image.version=$STACKEDIT_VERSION \
    org.opencontainers.image.revision=$VCS_REF \
    org.opencontainers.image.url="https://github.com/qdm12/stackedit-docker" \
    org.opencontainers.image.documentation="https://github.com/qdm12/stackedit-docker/blob/master/README.md" \
    org.opencontainers.image.source="https://github.com/qdm12/stackedit-docker" \
    org.opencontainers.image.title="stackedit-docker" \
    org.opencontainers.image.description="StackEdit server in a lightweight Docker container"
EXPOSE 8000
HEALTHCHECK --start-period=1s --interval=100s --timeout=2s --retries=1 CMD ["/server","healthcheck"]
USER 1000
ENTRYPOINT ["/server"]
COPY --from=stackedit --chown=1000 /stackedit/dist   /html/dist
COPY --from=stackedit --chown=1000 /stackedit/static /html/static
COPY --from=server --chown=1000 /tmp/gobuild/app /server
