FROM alpine:3.8 AS stackedit
RUN echo https://alpine.global.ssl.fastly.net/alpine/v3.8/main > /etc/apk/repositories && \
    apk add -q --progress --update --no-cache git npm
RUN git clone --branch v5.13.2 --single-branch --depth 1 https://github.com/benweet/stackedit.git /stackedit &> /dev/null
WORKDIR /stackedit
RUN npm install
RUN npm run build

FROM nginx:1.15-alpine
LABEL maintainer="quentin.mcgaw@gmail.com" \
      description="StackEdit server in a lightweight Docker container" \
      download="15.6MB" \
      size="45.5MB" \
      ram="7MB" \
      cpu_usage="Very low" \
      github="https://github.com/qdm12/stackedit-docker"
COPY --from=stackedit /stackedit/dist /usr/share/nginx/html/
ENTRYPOINT nginx -g "daemon off;"
