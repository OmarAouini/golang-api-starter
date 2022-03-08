package entities

import "time"

type Project struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Expenses  float64   `json:"expenses"`
	Incomes   float64   `json:"incomes"`
	StartAt   time.Time `json:"start_at"`
	UpdatedAt time.Time `json:"update_at"`
	EndAt     time.Time `json:"end_at"`
	CompanyID int       `json:"company_id"`
}

func (c *Project) TableName() string {
	return "projects"
}
