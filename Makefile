EXECUTORS_PLUGIN_DIRS=$(wildcard ./plugins/executors/*)
SOURCES_PLUGIN_DIRS=$(wildcard ./plugins/sources/*)

all: install protoc build-plugins run

build-plugins: $(EXECUTORS_PLUGIN_DIRS) $(SOURCES_PLUGIN_DIRS)

run:
	go run main.go

install:
	go get -u google.golang.org/protobuf@v1.28.1
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	go get -u google.golang.org/grpc@v1.50.0

protoc:
	# Clean
	rm -rf api/source/proto/*.pb.go api/executor/proto/*pb.go

	# Compile
	cd api/source
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative api/source/**/*.proto

	cd api/executor
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative api/executor/**/*.proto

	## Refresh dependencies
	cd api/source/proto && go mod tidy
	cd api/executor/proto && go mod tidy

$(EXECUTORS_PLUGIN_DIRS):
	$(MAKE) -C $@

$(SOURCES_PLUGIN_DIRS):
	$(MAKE) -C $@

.PHONY: all $(EXECUTORS_PLUGIN_DIRS) $(SOURCES_PLUGIN_DIRS)