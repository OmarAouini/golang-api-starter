package database

import "github.com/OmarAouini/golang-api-starter/entities"

func Migrate() {
	DB.AutoMigrate(
		&entities.Company{},
		&entities.Employee{},
		&entities.Project{},
	)
}
