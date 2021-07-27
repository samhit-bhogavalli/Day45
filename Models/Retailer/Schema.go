package Retailer

type Retailer struct {
	Name string `json:"name"`
}

func (r *Retailer) TableName() string {
	return "retailer"
}
