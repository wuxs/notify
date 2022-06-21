GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/linux/notify .
GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o bin/arm64/notify .
GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "-s -w" -o bin/armv7/notify .
