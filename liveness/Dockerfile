FROM golang
MAINTAINER liveness
ADD . /go/src/
EXPOSE 8080
CMD ["/usr/local/go/bin/go", "run", "/go/src/liveness.go"]
