package entities

// Product struct (Model)
type Product struct {
	Id     int `gorm:"primary_key, AUTO_INCREMENT"`
	Name   string 
	Price  float64
	Quantity int
	Status bool
}

func (product *Product) TableName() string {
	return "product"
}
