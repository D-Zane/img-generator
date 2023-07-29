.PHONY: build
build:
	rm -rf build && mkdir build && go build -o build/img_generator -v ./cmd/img-generator

run:
	go run cmd/img-generator/main.go