proto:
	protoc internal/auth/pb/auth.proto --go_out=plugins=grpc:.

build:
	docker build -t kit .
