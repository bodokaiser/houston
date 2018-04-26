GOARM=7
GOARCH=arm
GOOS=linux

build: build-dds

build-dds:
	GOARM=${GOARM} GOARCH=${GOARCH} GOOS=${GOOS} \
		go build -o bin/dds ./cmd/dds
