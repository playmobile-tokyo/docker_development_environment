#FROM golang:1.10.3-alpine3.8
FROM golang:latest

WORKDIR /go

ADD . /go


CMD ["go", "get", "-u", "github.com/mercari/widebullet/..."]

CMD ["wbt", "-c", "/go/config/example.toml"]

#CMD ["go", "run", "./src/hello.go"]

