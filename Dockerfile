FROM golang:1.17-alpine3.15 as builder

WORKDIR /workspace
RUN apk add --no-cache gcc libc-dev
COPY . .

RUN go test -v -cover ./... && \
  go build -a

##########################

FROM alpine:3.15.0 as final

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY --from=builder /workspace/topokube .

CMD ["./topokube"]
