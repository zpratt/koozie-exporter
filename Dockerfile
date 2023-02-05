FROM golang:1.17-alpine3.15 as builder

WORKDIR /workspace
RUN apk update && apk add --no-cache gcc=10.3.1_git20211027-r0 libc-dev=0.7.2-r3
COPY . .

RUN go build -a

##########################

FROM alpine:3.17.1 as final

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY --from=builder /workspace/topokube .

CMD ["./topokube"]
