FROM golang:1.9-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . $GOPATH/src/github.com/bailantaotao/misc
RUN cd $GOPATH/src/github.com/bailantaotao/misc/cmd/log && go build && mv log /log

# Pull mapi into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /log /usr/local/bin/

ENTRYPOINT ["log"]
