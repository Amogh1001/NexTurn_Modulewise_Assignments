package inventory

import (
	"inventory-management/inventory/models"
	"inventory-management/inventory/services"
)

func AddProduct(id int, name string, price float64, stock int) error {
	return services.AddProduct(id, name, price, stock)
}

func UpdateStock(id int, stock int) error {
	return services.UpdateStock(id, stock)
}

func SearchProduct(input string) (models.Product, error) {
	return services.SearchProduct(input)
}

func DisplayInventory() {
	services.DisplayInventory()
}

func SortInventoryBy(option string) error {
	return services.SortInventoryBy(option)
}
