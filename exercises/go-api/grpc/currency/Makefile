.PHONY: protos

protos:
	export GO111MODULE=on
	go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
	export PATH="$PATH:$(go env GOPATH)/bin"
	protoc -I protos/ --go-grpc_out=protos/currency --go_out=protos/currency protos/currency.proto