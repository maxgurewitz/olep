FROM golang:1.13.5

LABEL maintainer="Max Gurewitz <maxgurewitz@gmail.com>"

ADD . /go/src/github.com/maxgurewitz/olep

EXPOSE 8080

WORKDIR /go/src/github.com/maxgurewitz/olep

CMD ["go", "run", "cmd/request_handler/main.go"]
