package services

import (
	"publisher/config"

	"gorm.io/gorm"
)

type Context struct {
	DB *gorm.DB
	Config *config.Config
}
