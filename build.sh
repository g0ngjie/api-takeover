#/bin/bash

# 版本号
VERSION=0.0.1

# windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./output/takeover-amd64-${VERSION}.exe main.go
