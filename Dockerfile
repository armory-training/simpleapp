FROM golang

ADD . /go/src/github.com/armory-training/simpleapp
WORKDIR /go/src/github.com/armory-training/simpleapp
RUN go build
EXPOSE 8080
ENTRYPOINT ./simpleapp
