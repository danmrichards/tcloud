GOARCH=amd64

build: linux darwin windows

linux:
	GO111MODULE=on CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=linux go build -mod=vendor -o ./bin/tcloud-linux-${GOARCH}

darwin:
	GO111MODULE=on CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=darwin go build -mod=vendor -o ./bin/tcloud-darwin-${GOARCH}

windows:
	GO111MODULE=on CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=windows go build -mod=vendor -o ./bin/tcloud-windows-${GOARCH}.exe

.PHONY: build
