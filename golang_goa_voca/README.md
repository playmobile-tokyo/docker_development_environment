go 環境

docker-compose build
  go get -u github.com/goadesign/goa/...
  go get -u github.com/golang/dep/cmd/dep

docker-compose up -d
  goagen bootstrap -d vocatuku/design
  go build -o vocatukuapi

curl http://localhost:5000/voices/1
