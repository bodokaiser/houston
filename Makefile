GOARM=7
GOARCH=arm
GOOS=linux

build:
	GOARM=${GOARM} GOARCH=${GOARCH} GOOS=${GOOS} \
		go build -o bin/dds/const cmd/dds/const/main.go
