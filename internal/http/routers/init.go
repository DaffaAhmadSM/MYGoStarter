package routers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouterInit(db *gorm.DB) *echo.Echo {
	mux := echo.New()

	web(mux, db)
	api(mux, db)

	return mux
}
