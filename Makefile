
build:
	@go build -o bin/main main.go

run:
	@go run -race main.go server --file=dev.env

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