
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "tangram-api" -ldflags="-s -w" *.go