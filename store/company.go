package store

import (
	"github.com/OmarAouini/golang-api-starter/entities"
	"gorm.io/gorm"
)

type CompanyStore interface {
	Get(id int) (*entities.Company, error)
	GetByName(name string) (*entities.Company, error)
}

type MySqlCompanyStore struct {
	DB *gorm.DB
}

func (s *MySqlCompanyStore) Get(id int) (*entities.Company, error) {
	var comp entities.Company
	err := s.DB.Preload("Employee").Preload("Project").Where("id = ?", id).First(&comp).Error
	if err != nil {
		return nil, err
	}
	return &comp, nil
}
