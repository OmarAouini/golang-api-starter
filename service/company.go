package service

import (
	"github.com/OmarAouini/golang-api-starter/entities"
	"github.com/OmarAouini/golang-api-starter/store"
)

type CompanyService struct {
	Store store.CompanyStore
}

func (s *CompanyService) All() (*[]entities.Company, error) {
	comps, err := s.Store.All()
	if err != nil {
		return nil, err
	}
	return comps, nil
}

func (s *CompanyService) Get(id int) (*entities.Company, error) {
	comps, err := s.Store.Get(id)
	if err != nil {
		return nil, err
	}
	return comps, nil
}

func (s *CompanyService) GetByName(name string) (*entities.Company, error) {
	comps, err := s.Store.GetByName(name)
	if err != nil {
		return nil, err
	}
	return comps, nil
}

func (s *CompanyService) Create(comp *entities.Company) error {
	err := s.Store.Create(comp)
	if err != nil {
		return err
	}
	return nil
}

func (s *CompanyService) Delete(id int) error {
	err := s.Store.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
