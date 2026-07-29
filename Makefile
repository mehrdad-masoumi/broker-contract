.PHONY: proto clean deps tidy test vet fmt buf-generate

PROTOC ?= protoc
PROTOC_GEN_GO_VERSION ?= v1.36.6
PROTOC_GEN_GO_GRPC_VERSION ?= v1.5.1

# Optional local protoc (e.g. copied from another repo's tools/protoc)
ifneq ($(wildcard tools/protoc/bin/protoc$(EXE)),)
PROTOC := tools/protoc/bin/protoc$(EXE)
endif

PROTOC_INCLUDE := --proto_path=proto
ifneq ($(wildcard tools/protoc/include),)
PROTOC_INCLUDE += --proto_path=tools/protoc/include
endif

PROTO_FILES := proto/notification/v1/notification.proto

proto:
	$(PROTOC) --go_out=gen/go --go_opt=paths=source_relative \
		--go-grpc_out=gen/go --go-grpc_opt=paths=source_relative \
		$(PROTOC_INCLUDE) \
		$(PROTO_FILES)

buf-generate:
	buf generate

clean:
	rm -f gen/go/notification/v1/*.pb.go

deps:
	go mod download
	go install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)

tidy:
	go mod tidy

fmt:
	gofmt -w go gen/go examples

vet:
	go vet ./...

test:
	go test ./...
