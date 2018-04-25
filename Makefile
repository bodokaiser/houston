GOARM=7
GOARCH=arm
GOOS=linux

build: build-const build-reset build-sweep build-display

build-const:
	GOARM=${GOARM} GOARCH=${GOARCH} GOOS=${GOOS} \
		go build -o bin/dds/const ./cmd/dds/const

build-reset:
	GOARM=${GOARM} GOARCH=${GOARCH} GOOS=${GOOS} \
		go build -o bin/dds/reset ./cmd/dds/reset

build-sweep: build-sweep-amplitude build-sweep-frequency build-sweep-phase

build-display:
	GOARM=${GOARM} GOARCH=${GOARCH} GOOS=${GOOS} \
		go build -o bin/dds/display ./cmd/dds/display

build-sweep-amplitude:
	GOARM=${GOARM} GOARCH=${GOARCH} GOOS=${GOOS} \
		go build -o bin/dds/sweep/amplitude ./cmd/dds/sweep/amplitude

build-sweep-frequency:
	GOARM=${GOARM} GOARCH=${GOARCH} GOOS=${GOOS} \
		go build -o bin/dds/sweep/frequency ./cmd/dds/sweep/frequency

build-sweep-phase:
	GOARM=${GOARM} GOARCH=${GOARCH} GOOS=${GOOS} \
		go build -o bin/dds/sweep/phase ./cmd/dds/sweep/phase