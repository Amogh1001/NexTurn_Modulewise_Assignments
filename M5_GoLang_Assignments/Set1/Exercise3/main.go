package main

import (
	"bufio"
	"fmt"
	"inventory-management/inventory"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Println("1. Add Product")
		fmt.Println("2. Update Stock")
		fmt.Println("3. Search Product")
		fmt.Println("4. Display Inventory")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			var id, stock int
			var name string
			var price float64
			fmt.Print("Enter Product ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Product Name: ")
			reader := bufio.NewReader(os.Stdin)
			input, readErr := reader.ReadString('\n')
			if readErr != nil {
				fmt.Println("Error:", readErr)
				name = "Default Product"
			} else {
				name = strings.TrimSpace(input)
			}
			fmt.Print("Enter Product Price: ")
			fmt.Scanln(&price)
			fmt.Print("Enter Product Stock: ")
			fmt.Scanln(&stock)
			err := inventory.AddProduct(id, name, price, stock)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case 2:
			var id, newStock int
			fmt.Print("Enter Product ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter New Stock: ")
			fmt.Scanln(&newStock)
			err := inventory.UpdateStock(id, newStock)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case 3:
			var query string
			fmt.Print("Enter Product ID or Name: ")
			reader := bufio.NewReader(os.Stdin)
			input, readErr := reader.ReadString('\n')
			if readErr != nil {
				fmt.Println("Error:", readErr)
				query = "Default Product"
			} else {
				query = strings.TrimSpace(input)
			}
			product, err := inventory.SearchProduct(query)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Product Found: %v\n", product)
			}

		case 4:
			inventory.DisplayInventory()

		case 5:
			var option string
			fmt.Print("Sort by 'price' or 'stock': ")
			fmt.Scanln(&option)
			err := inventory.SortInventoryBy(option)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				inventory.DisplayInventory()
			}

		case 6:
			fmt.Println("Exiting the program. Goodbye!")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
