echo "building..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "./build/bin/api" -ldflags="-s -w" *.go