GOARM=7
GOARCH=arm
GOOS=linux

build: build-dds build-httpdev build-http

build-dds:
	GOARM=${GOARM} GOARCH=${GOARCH} GOOS=${GOOS} \
		go build -o bin/dds ./cmd/dds

build-http:
	GOARM=${GOARM} GOARCH=${GOARCH} GOOS=${GOOS} \
		go build -o bin/http ./cmd/http

build-httpdev:
	GOARM=${GOARM} GOARCH=${GOARCH} GOOS=${GOOS} \
		go build -o bin/httpdev ./cmd/httpdev
