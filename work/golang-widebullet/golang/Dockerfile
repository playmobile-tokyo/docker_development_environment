# FROM golang:1.10.3-alpine3.8
FROM golang:latest

WORKDIR /go

ADD . /go

CMD ["go", "get", "-u", "github.com/mercari/widebullet/..."]

# CMD ["go", "run", "./src/github.com/mercari/widebullet/cmd/wbt/wbt.go", "-c", "./src/config/example.toml"]
