package entities

type Company struct {
	ID          int         `json:"id" gorm:"primaryKey"`
	Name        string      `json:"name"`
	Email       string      `json:"email"`
	Address     string      `json:"address"`
	PhoneNumber string      `json:"phone_number"`
	VatCode     string      `json:"vat_code"`
	Employee    *[]Employee `json:"employees" gorm:"foreignKey:company_id"`
	Project     *[]Project  `json:"projects" gorm:"foreignKey:company_id"`
}

type Tabler interface {
	TableName() string
}

func (c *Company) TableName() string {
	return "companies"
}
