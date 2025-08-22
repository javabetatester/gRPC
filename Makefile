PROTO_DIR = proto
PROTO_FILES = $(wildcard $(PROTO_DIR)/*.proto)

.PHONY: deps proto build clean

deps:
	go mod tidy
	go mod download

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)

build: deps
	go build -o bin/server main.go

run: build
	./bin/server

clean:
	rm -rf bin/
	go clean

test:
	go test ./...

install-protoc:
	@echo "Instale o protoc manualmente:"
	@echo "Windows: baixe de https://github.com/protocolbuffers/protobuf/releases"
	@echo "Ubuntu: sudo apt install protobuf-compiler"
	@echo "macOS: brew install protobuf"
