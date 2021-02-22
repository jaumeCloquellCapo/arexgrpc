
run:
	@go run cmd/server/main.go

client:
	@go run cmd/client/main.go

test:							## Run all tests
	@go test ./...

migrate_dev:
	@goose -dir ./migrations mysql "db:db@tcp(localhost)/db?parseTime=true" up

proto:
	protoc -I grpc -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:grpc \
		grpc/*.proto
	protoc -I grpc -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:grpc \
		grpc/*.proto