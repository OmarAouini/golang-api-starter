package store

import (
	"github.com/OmarAouini/golang-api-starter/entities"
)

type EmployeeStore interface {
	Get(id int) (*entities.Employee, error)
	GetByName(name string) (*entities.Employee, error)
}

type MySqlEmployeeStore struct{}
