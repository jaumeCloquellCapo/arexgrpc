FROM golang:1.15.6

WORKDIR /app

COPY ./ /app

RUN go mod download
RUN go get -d -v ./...
RUN go install -v ./...
RUN GOOS=linux CGO_ENABLED=0 go build cmd/server/main.go

EXPOSE 8888

CMD ["/src/main"]