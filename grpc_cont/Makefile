generate_grpc_code:
	protoc \
	--go_out=GRPC \
	--go_opt=paths=source_relative \
	--go-grpc_out=GRPC \
	--go-grpc_opt=paths=source_relative \
	--plugin=/home/nemael/go/bin/protoc-gen-go \
	--plugin=/home/nemael/go/bin/protoc-gen-go-grpc \
	GRPC.proto