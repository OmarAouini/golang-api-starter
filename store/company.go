package store

import (
	"github.com/OmarAouini/golang-api-starter/database"
	"github.com/OmarAouini/golang-api-starter/entities"
)

type CompanyStore interface {
	All() (*[]entities.Company, error)
	Get(id int) (*entities.Company, error)
	GetByName(name string) (*entities.Company, error)
	Create(new *entities.Company) error
}

type MySqlCompanyStore struct{}

func (s *MySqlCompanyStore) All() (*[]entities.Company, error) {
	var comp *[]entities.Company
	err := database.DB.Preload("Employee").Preload("Project").Find(&comp).Error
	if err != nil {
		return nil, err
	}
	return comp, nil
}

func (s *MySqlCompanyStore) Get(id int) (*entities.Company, error) {
	var comp entities.Company
	err := database.DB.Preload("Employee").Preload("Project").Where("id = ?", id).First(&comp).Error
	if err != nil {
		return nil, err
	}
	return &comp, nil
}

func (s *MySqlCompanyStore) GetByName(name string) (*entities.Company, error) {
	var comp entities.Company
	err := database.DB.Preload("Employee").Preload("Project").Where("name = ?", name).First(&comp).Error
	if err != nil {
		return nil, err
	}
	return &comp, nil
}

func (s *MySqlCompanyStore) Create(new *entities.Company) error {
	err := database.DB.Save(&new).Error
	if err != nil {
		return err
	}
	return nil
}
