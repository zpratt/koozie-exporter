FROM node:14-alpine3.14 as builder

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
ENV CI=true

COPY . .

RUN apk update && apk --no-cache --virtual build-dependencies add \
    python3=3.9.5-r2 \
    make=4.3-r0 \
    g++=10.3.1_git20210424-r2 \
    && npm ci --no-optional --no-cache \
    && npm run build \
    && apk del build-dependencies

#######################################

FROM nginx:1.21.4-alpine

RUN mkdir -p /var/www/ui
WORKDIR /var/www/ui

COPY --from=builder /usr/src/app/out .

COPY --from=builder /usr/src/app/nginx.conf /etc/nginx/conf.d/default.conf
