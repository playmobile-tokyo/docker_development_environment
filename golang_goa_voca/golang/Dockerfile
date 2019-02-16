#FROM golang:1.10.3-alpine3.8
FROM golang:latest

#USE root

WORKDIR /go

ADD . /go

CMD ["go", "get", "-u", "github.com/goadesign/goa/..."]
CMD ["go", "get", "-u", "github.com/golang/dep/cmd/dep"]
# CMD ["dep", "init"]

# CMD ["go", "run", "./src/hello.go"]
