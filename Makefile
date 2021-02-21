
build:
	@go build -o bin/main main.go

run:
	@go run -race main.go server --file=dev.env

test:							## Run all tests
	@go test ./...

migrate_pro :
	@goose -dir ./migrations mysql "db:db@tcp(db)/db?parseTime=true" up

migrate_dev:
	@goose -dir ./migrations mysql "db:db@tcp(localhost)/db?parseTime=true" up

dev_up:
	@docker-compose -f docker-compose.dev.yml up

dev_down:
	@docker-compose -f docker-compose.dev.yml down

pro_up:
	@docker-compose -f docker-compose.yml up

pro_down:
	@docker-compose -f docker-compose.yml down

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