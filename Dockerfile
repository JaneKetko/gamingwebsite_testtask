FROM golang:latest

RUN mkdir -p /go/src/github.com/Ragnar-BY/gamingwebsite_testtask
WORKDIR /go/src/github.com/Ragnar-BY/gamingwebsite_testtask
COPY . /go/src/github.com/Ragnar-BY/gamingwebsite_testtask

RUN go get ./...
RUN go build github.com/Ragnar-BY/gamingwebsite_testtask

ENTRYPOINT /go/bin/gamingwebsite_testtask