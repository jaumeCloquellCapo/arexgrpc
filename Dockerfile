FROM golang:1.15.6

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 8888

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main