package container

import (
	"github.com/Axit88/UserService/src/domain/userService/application"
	core "github.com/Axit88/UserService/src/domain/userService/core"
	infrastructure "github.com/Axit88/UserService/src/domain/userService/infrastructure"
	"github.com/Axit88/UserService/src/utils/loggerUtil"
	"go.uber.org/dig"
)

func UserServiceContainer() (*dig.Container, error) {
	container := dig.New()
	err := container.Provide(loggerUtil.InitLogger)
	err = container.Provide(infrastructure.NewUserServiceClient)
	err = container.Provide(application.NewUserServiceApplication)
	err = container.Provide(core.NewUserServiceFacade)

	return container, err
}
