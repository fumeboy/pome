#!/usr/bin/env bash
go build -o main
mv ./main ./build/app/main
docker build -t pome ./build/