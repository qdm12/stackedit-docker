ARG ALPINE_VERSION=3.12
ARG GO_VERSION=1.14
ARG STACKEDIT_VERSION=v5.14.5

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS server
RUN apk --update add git
ENV CGO_ENABLED=0
WORKDIR /tmp/gobuild
COPY go.mod go.sum ./
RUN go mod download 2>&1
COPY main.go ./
#RUN go test -v -race ./...
RUN go build -ldflags="-s -w" -o app main.go

FROM --platform=amd64 alpine:${ALPINE_VERSION} AS stackedit
ARG STACKEDIT_VERSION
WORKDIR /stackedit
RUN apk add -q --progress --update --no-cache git npm python2
RUN wget -q https://github.com/benweet/stackedit/archive/${STACKEDIT_VERSION}.tar.gz -O stackedit.tar.gz && \
    tar -xzf stackedit.tar.gz --strip-components=1 && \
    rm stackedit.tar.gz
RUN npm install
RUN npm audit fix
ENV NODE_ENV=production
RUN sed -i "s/assetsPublicPath: '\/',/assetsPublicPath: '.\/',/g" config/index.js && cat config/index.js
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
ENV \
    LISTENING_PORT=8000 \
    ROOT_URL=/ \
    NODE_ENV=production \
    PANDOC_PATH=pandoc \
    WKHTMLTOPDF_PATH=wkhtmltopdf \
    USER_BUCKET_NAME=stackedit-users \
    PAYPAL_RECEIVER_EMAIL= \
    DROPBOX_APP_KEY= \
    DROPBOX_APP_KEY_FULL= \
    GITHUB_CLIENT_ID= \
    GITHUB_CLIENT_SECRET= \
    GOOGLE_CLIENT_ID= \
    GOOGLE_API_KEY= \
    WORDPRESS_CLIENT_ID=
COPY --from=stackedit --chown=1000 /stackedit/dist   /html/dist
COPY --from=stackedit --chown=1000 /stackedit/static /html/static
COPY --from=server --chown=1000 /tmp/gobuild/app /server
