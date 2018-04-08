FROM golang:1.9-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . $GOPATH/src/github.com/bailantaotao/misc
RUN cd $GOPATH/src/github.com/bailantaotao/misc/cmd/hello && go build && mv hello /hello

# Pull mapi into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /hello /usr/local/bin/

ENTRYPOINT ["hello"]
