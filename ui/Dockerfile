FROM node:lts-bullseye-slim as builder

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
ENV CI=true

COPY . .

RUN npm ci --omit=optional --legacy-peer-deps \
    && npm test \
    && npm run lint \
    && npm run build

#######################################

FROM nginx:1.21.4-alpine

RUN mkdir -p /var/www/ui
WORKDIR /var/www/ui

COPY --from=builder /usr/src/app/out .

COPY --from=builder /usr/src/app/nginx.conf /etc/nginx/conf.d/default.conf
