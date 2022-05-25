proto:
	protoc pkg/pb/*.proto --go_out=plugins=grpc:.

order-server:
	go run cmd/main.go