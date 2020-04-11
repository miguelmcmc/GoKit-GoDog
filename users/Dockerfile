FROM golang

RUN mkdir -p /go/src/ironchip.net/kit

ADD . /go/src/ironchip.net/kit

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run ironchip.net/kit/users/cmd  -watch ironchip.net/kit/users
