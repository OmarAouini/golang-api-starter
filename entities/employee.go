package entities

type Employee struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	Name          string `json:"name"`
	Surname       string `json:"surname" `
	Age           int    `json:"age" `
	Email         string `json:"email" `
	Address       string `json:"address" `
	PhoneNumber   string `json:"phone_number" `
	VatCode       string `json:"vat_code" `
	Qualification string `json:"qualification"`
	CompanyID     int    `json:"company_id"`
}

func (c *Employee) TableName() string {
	return "employees"
}
