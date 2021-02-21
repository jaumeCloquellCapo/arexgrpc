package dic

import (
	"github.com/jaumeCloquellCapo/authGrpc/app/repository"
	"github.com/jaumeCloquellCapo/authGrpc/app/service"
	"github.com/jaumeCloquellCapo/authGrpc/internal/middleware"
	"github.com/jaumeCloquellCapo/authGrpc/internal/storage"
	"github.com/sarulabs/dingo/generation/di"
	"log"
)

//var Builder *di.Builder
//var Container di.Container

const DbService = "db"
const CacheService = "cache"

const AuthMiddleware = "middleware.auth"
const CorsMiddleware = "middleware.cors"

const UserRepository = "repository.user"
const UserService = "service.user"
const UserController = "controller.user"

const AuthRepository = "repository.auth"
const AuthService = "service.auth"
const AuthController = "controller.auth"

// dependency injection container
func InitContainer() di.Container {
	builder, err := di.NewBuilder()
	if err != nil {
		log.Fatal(err.Error())
	}
	RegisterServices(builder)
	return builder.Build()
}

func RegisterServices(builder *di.Builder) {
	builder.Add(di.Def{
		Name: DbService,
		Build: func(ctn di.Container) (interface{}, error) {
			return storage.InitializeDB(), nil
		},
		Close: func(obj interface{}) error {
			obj.(*storage.DbStore).Close()
			return nil
		},
	})
	builder.Add(di.Def{
		Name: CacheService,
		Build: func(ctn di.Container) (interface{}, error) {
			return storage.InitializeCache(), nil
		},
		Close: func(obj interface{}) error {
			obj.(*storage.DbCache).Close()
			return nil
		},
	})

	builder.Add(di.Def{
		Name: AuthMiddleware,
		Build: func(ctn di.Container) (interface{}, error) {
			return middleware.NewAuthMiddleware(ctn.Get(AuthService).(service.AuthServiceInterface)), nil
		},
	})

	builder.Add(di.Def{
		Name: CorsMiddleware,
		Build: func(ctn di.Container) (interface{}, error) {
			return middleware.NewCorsMiddleware(), nil
		},
	})

	builder.Add(di.Def{
		Name: UserRepository,
		Build: func(ctn di.Container) (interface{}, error) {
			return repository.NewUserRepository(ctn.Get(DbService).(*storage.DbStore)), nil
		},
	})
	builder.Add(di.Def{
		Name: AuthRepository,
		Build: func(ctn di.Container) (interface{}, error) {
			return repository.NewAuthRepository(ctn.Get(CacheService).(*storage.DbCache)), nil
		},
	})

	builder.Add(di.Def{
		Name: UserService,
		Build: func(ctn di.Container) (interface{}, error) {
			return service.NewUserService(ctn.Get(UserRepository).(repository.UserRepositoryInterface)), nil
		},
	})

	builder.Add(di.Def{
		Name: AuthService,
		Build: func(ctn di.Container) (interface{}, error) {
			return service.NewAuthService(ctn.Get(AuthRepository).(repository.AuthRepositoryInterface), ctn.Get(UserRepository).(repository.UserRepositoryInterface)), nil
		},
	})
	/*
		builder.Add(di.Def{
			Name: UserController,
			Build: func(ctn di.Container) (interface{}, error) {
				return controller.NewUserController(ctn.Get(UserService).(service.UserServiceInterface)), nil
			},
		})

		builder.Add(di.Def{
			Name: AuthController,
			Build: func(ctn di.Container) (interface{}, error) {
				return controller.NewAuthController(ctn.Get(AuthService).(service.AuthServiceInterface), ctn.Get(UserService).(service.UserServiceInterface)), nil
			},
		})*/
}
