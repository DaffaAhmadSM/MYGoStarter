package controllers

import "gorm.io/gorm"

type MainController struct {
	Database *gorm.DB
}
