proto_server:
	del /Q server\server_grpc\*.pb.go
	C:/protobuff/bin/protoc.exe --proto_path=server --go_out=server/server_grpc --go_opt=paths=source_relative \
	--go-grpc_out=server/server_grpc --go-grpc_opt=paths=source_relative \
	server/*.proto

proto_client:
	del /Q client\client_grpc\*.pb.go
	C:/protobuff/bin/protoc.exe --proto_path=client --go_out=client/client_grpc --go_opt=paths=source_relative \
	--go-grpc_out=client/client_grpc --go-grpc_opt=paths=source_relative \
	client/*.proto