CGO_ENABLED=0 GOPROXY=https://gocenter.io,https://goproxy.io,direct go build
docker build -t pome:v2 .