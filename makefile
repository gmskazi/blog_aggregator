BINARY_NAME=out

build:
	go build -o ${BINARY_NAME} && ./out

test:
	go test ./ ...

test_coverage:
	go test ./ ... -coverprofile=coverage.out

lint:
	golangci-lint run --enable-all
