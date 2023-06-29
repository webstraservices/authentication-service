BINARY=webstraAuth

build:
	@go build -o bin/${BINARY}

run: build
	@bin/${BINARY}

