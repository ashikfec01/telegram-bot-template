package repositories

import (
	"gorm.io/gorm"
	"shopinbiz.telegram.bot/database"
)

var DB *gorm.DB = database.Init()
