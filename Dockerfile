# Sets linux/amd64 in case it's not injected by older Docker versions
ARG BUILDPLATFORM=linux/amd64

ARG ALPINE_VERSION=3.14
ARG STACKEDIT_VERSION=v5.14.10
ARG GO_VERSION=1.17
ARG XCPUTRANSLATE_VERSION=v0.6.0
ARG GOLANGCI_LINT_VERSION=v1.42.1

FROM --platform=${BUILDPLATFORM} qmcgaw/xcputranslate:${XCPUTRANSLATE_VERSION} AS xcputranslate
FROM --platform=${BUILDPLATFORM} qmcgaw/binpot:golangci-lint-${GOLANGCI_LINT_VERSION} AS golangci-lint

FROM --platform=${BUILDPLATFORM} golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS base
ENV CGO_ENABLED=0
WORKDIR /tmp/gobuild
RUN apk --update add git g++
COPY --from=xcputranslate /xcputranslate /usr/local/bin/xcputranslate
COPY --from=golangci-lint /bin /go/bin/golangci-lint
COPY go.mod go.sum ./
RUN go mod download
COPY main.go .

FROM base AS lint
COPY .golangci.yml ./
RUN golangci-lint run --timeout=10m

FROM base AS server
ARG TARGETPLATFORM
ARG VERSION=unknown
ARG CREATED="an unknown date"
ARG COMMIT=unknown
RUN GOARCH="$(xcputranslate translate -targetplatform=${TARGETPLATFORM} -field arch)" \
    GOARM="$(xcputranslate translate -targetplatform=${TARGETPLATFORM} -field arm)" \
    go build -trimpath -ldflags="-s -w \
    -X 'main.version=$VERSION' \
    -X 'main.buildDate=$CREATED' \
    -X 'main.commit=$COMMIT' \
    " -o app main.go

FROM --platform=amd64 alpine:${ALPINE_VERSION} AS stackedit
ARG STACKEDIT_VERSION
WORKDIR /stackedit
RUN apk add -q --progress --update --no-cache git npm python3 make g++
RUN git clone --branch ${STACKEDIT_VERSION} --single-branch --depth 1 https://github.com/benweet/stackedit.git . &> /dev/null
RUN npm install
ENV NODE_ENV=production
RUN sed -i "s/assetsPublicPath: '\/',/assetsPublicPath: '.\/',/g" config/index.js
RUN npm run build

FROM scratch AS final
ARG CREATED
ARG COMMIT
ARG STACKEDIT_VERSION
LABEL \
    org.opencontainers.image.authors="quentin.mcgaw@gmail.com" \
    org.opencontainers.image.created=$CREATED \
    org.opencontainers.image.version=$STACKEDIT_VERSION \
    org.opencontainers.image.revision=$COMMIT \
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
