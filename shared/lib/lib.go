package lib

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// AppContext return custom application context
type (
	AppContext struct {
		echo.Context
		MysqlSession *gorm.DB
	}
)
