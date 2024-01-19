PROTO_DIR=./proto
OUTPUT_DIR=./proto


generate-proto:
	@echo "Generating code..."
	find $(PROTO_DIR) -name "*.proto" | xargs protoc -I $(PROTO_DIR) \
		--go_out $(OUTPUT_DIR) --go_opt paths=source_relative \
		--go-grpc_out $(OUTPUT_DIR) --go-grpc_opt paths=source_relative \
		--grpc-gateway_out $(OUTPUT_DIR) --grpc-gateway_opt paths=source_relative
	@echo "Code generation completed."

install-protoc:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

	export PATH="$PATH:$(go env GOPATH)/bin"