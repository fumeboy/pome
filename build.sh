CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY=https://gocenter.io,https://goproxy.io,direct go build
docker build -t pome:v2 .