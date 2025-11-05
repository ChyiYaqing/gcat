.PHONY: generate
generate:
	protoc \
		-I . \
		-I third_party/vtprotobuf \
		-I third_party/googleapis \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		--go-vtproto_out=. \
		--go-vtproto_opt=paths=source_relative \
		--go-vtproto_opt=features=marshal+unmarshal+size+pool \
		pkg/proto/demo/v1/myservice.proto \
		pkg/proto/users/v1/users.proto