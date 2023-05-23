client-codegen:
	protoc --go_out=../client/rpc --go_opt=paths=source_relative \
	--go-grpc_out=../client/rpc  --go-grpc_opt=paths=source_relative \
	person.proto 
server-codegen:
	protoc --go_out=../server/rpc --go_opt=paths=source_relative \
	--go-grpc_out=../server/rpc  --go-grpc_opt=paths=source_relative \
	person.proto 
# client-codegen server-codegen
code-gen: server-codegen
	@echo "生成sever client代码成功" 