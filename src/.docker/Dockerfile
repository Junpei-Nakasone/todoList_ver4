
FROM golang:1.12.5-alpine3.9

WORKDIR /go/src

RUN apk update && \
    apk add --update git && \
    apk add openssh && \
    rm -rf /var/cache/apk/*

ENV GO111MODULE=on 
ENV CGO_ENABLED=0
ENV TZ = Asia/Tokyo

# EXPOSE 8888
# COPY ../server/go.mod /go/src/go.mod
# COPY ../server/go.sum /go/src/go.sum

RUN go mod download