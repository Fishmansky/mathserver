generate_proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false proto/request.proto

helper:
	@echo "Available options:"
	@echo "  make generate_proto - generate proto files"

.DEFAULT_GOAL := helper
