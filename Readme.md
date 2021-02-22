### Golang - GRPC
#### Full list what has been used:
* [GRPC](https://grpc.io/) - gRPC
* [sqlx](https://github.com/jmoiron/sqlx) - Extensions to database/sql.
* [pgx](https://github.com/jackc/pgx) - PostgreSQL driver and toolkit for Go
* [viper](https://github.com/spf13/viper) - Go configuration with fangs
* [go-redis](https://github.com/go-redis/redis) - Redis client for Golang
* [zap](https://github.com/uber-go/zap) - Logger
* [validator](https://github.com/go-playground/validator) - Go Struct and Field validation
* [migrate](https://github.com/golang-migrate/migrate) - Database migrations. CLI and Golang library.
* [testify](https://github.com/stretchr/testify) - Testing toolkit
* [gomock](https://github.com/golang/mock) - Mocking framework
* [Docker](https://www.docker.com/) - Docker

#### Recommendation for local development most comfortable usage:
    make run // run all containers
    make client // run the example client grpc to cal to our container

#### Docker-compose files:
    docker-compose.yml - run postgresql, redis


### Docker development usage:
    make docker

### Local development usage:
    make migrate_dev // to run the migrations 


This project is a implementation from the project https://github.com/jaumeCloquellCapo/go-api-boirplate to GRPC