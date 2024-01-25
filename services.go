package main

import (
	sharedApplication "api.kyoris.com/core/_shared/application"
	sharedDomain "api.kyoris.com/core/_shared/domain"
	userApplication "api.kyoris.com/core/user/application"
	userInfrastructure "api.kyoris.com/core/user/infrastructure"
	"gorm.io/gorm"
)

// Inicializamos los repsitorios para retornar un objeto de tipo Repositories
func initRepositories(dbConn *gorm.DB) *sharedDomain.Repositories {
	userRepository := userInfrastructure.NewMySQLRepository(dbConn)

	// Retornamos los repositorios inicializados para usarlos en los servicios
	return &sharedDomain.Repositories{
		UserRepository: userRepository,
	}
}

// Inicializamos los servicios para retornar un objeto de tipo Services
func initServices(dbConn *gorm.DB) *sharedApplication.Services {
	// Obtenemos todos los Repositorios inicializados
	repositories := initRepositories(dbConn)

	userService := userApplication.NewService(repositories.UserRepository)

	// Retornamos los servicios inicializados para usarlos en los Handlers
	return &sharedApplication.Services{
		UserService: userService,
	}
}
