package container

import (
	"fmt"

	"github.com/fgrosse/goldi"

	"github.com/hrz8/rahman-tennis/shared/config"
	"github.com/hrz8/rahman-tennis/shared/database"
)

// DefaultContainer returns default given depedency injections
func DefaultContainer() *goldi.Container {
	goldiRegistry := goldi.NewTypeRegistry()
	goldiConfig := make(map[string]interface{})
	container := goldi.NewContainer(goldiRegistry, goldiConfig)

	appConfigInterface, err := config.NewConfig()
	if err != nil {
		panic(fmt.Sprintf("[ERROR] no config file: %s", err.Error()))
	}

	container.InjectInstance("shared.config", appConfigInterface)
	container.RegisterType("shared.mysql", database.NewMysql, "@shared.config")

	return container
}
