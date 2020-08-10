FROM golang:1.14.7-alpine3.12 as builder

WORKDIR /workspace
RUN apk add --no-cache gcc libc-dev
COPY . .

RUN go build -a

##########################

FROM alpine:3.12 as final

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY --from=builder /workspace/topokube .

CMD ["./topokube"]
