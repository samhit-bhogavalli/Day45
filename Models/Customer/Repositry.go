package Customer

import "github.com/samhit-bhogavalli/Day45/Config"

func CreateCustomer(customer *Customer) error {
	if err := Config.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}
