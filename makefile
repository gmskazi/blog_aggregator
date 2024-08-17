BINARY_NAME=out

build:
	@go build -o ${BINARY_NAME}

run: build
	@./${BINARY_NAME}
test:
	@go test ./ ...

test_coverage:
	@go test ./ ... -coverprofile=coverage.out
