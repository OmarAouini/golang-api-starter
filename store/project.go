package store

import (
	"github.com/OmarAouini/golang-api-starter/entities"
	"gorm.io/gorm"
)

type ProjectStore interface {
	Get(id int) (*entities.Project, error)
	GetByName(name string) (*entities.Project, error)
}

type MySqlProjectStore struct {
	DB *gorm.DB
}
