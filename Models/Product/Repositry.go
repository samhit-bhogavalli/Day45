package Product

import "github.com/samhit-bhogavalli/Day45/Config"

func CreateProduct(product *Product) error {
	if err := Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func GetAllProducts(product *[]Product) error {
	if err := Config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductById(id string, product *Product) error {
	if err := Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProduct(id string, product *Product) error {
	if err := Config.DB.Save(product).Error; err != nil {
		return err
	}
	return nil
}
