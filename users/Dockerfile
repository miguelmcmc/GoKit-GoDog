FROM golang:alpine

RUN apk update ; apk add -U --no-cache tzdata bash wget curl git upx ca-certificates

WORKDIR /go/src/ironchip.net/kit

ADD . .

RUN GO111MODULE=on go mod download
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run ironchip.net/kit/users/cmd  -watch ironchip.net/kit/users
