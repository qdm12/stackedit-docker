FROM alpine:3.7
LABEL maintainer="quentin.mcgaw@gmail.com" \
      description="Run the latest StackEdit server in a Docker container" \
      size="581MB" \
      ram="340MB-500MB" \
      github="https://github.com/qdm12/stackedit-docker"
EXPOSE 8080
RUN apk add -q --progress --update --no-cache git nodejs && \
    git clone https://github.com/benweet/stackedit.git && \
    cd stackedit && \
    rm -rf .git .dockerignore .gitignore .travis.yml CHANGELOG.md \
            Dockerfile LICENSE README.md && \
    npm --silent install && \
    npm --silent run build && \
    apk del -q --progress --purge git
WORKDIR /stackedit
ENTRYPOINT npm start