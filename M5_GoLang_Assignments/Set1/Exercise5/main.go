package main

import (
	"bufio"
	climate "climate-analysis/Climate"
	"climate-analysis/Climate/models"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Println("1. Add new climate data for the city")
		fmt.Println("2. View city with the highest temperature")
		fmt.Println("3. View city with the lowest temperature")
		fmt.Println("4. Calculate average rainfall across all cities")
		fmt.Println("5. Filter cities by rainfall threshold")
		fmt.Println("6. Search city by name")
		fmt.Println("7. Exit")
		fmt.Print("Input your choice: ")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			var id int
			var name string
			var avgtemp, rainfall float64
			fmt.Print("Input Climate ID: ")
			fmt.Scanln(&id)
			fmt.Print("Input City Name: ")
			reader := bufio.NewReader(os.Stdin)
			input, readErr := reader.ReadString('\n')
			if readErr != nil {
				fmt.Println("Error occured in reading name. Error message : ", readErr)
				name = "Default City"
			} else {
				name = strings.TrimSpace(input)
			}
			fmt.Print("Input Avg Temperature in celsius: ")
			fmt.Scanln(&avgtemp)
			fmt.Print("Input Rainfall in mm: ")
			fmt.Scanln(&rainfall)
			err := climate.AddInput(id, name, avgtemp, rainfall)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 2:
			err := climate.HighestTemperature()
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 3:
			err := climate.LowestTemperature()
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 4:
			err := climate.CalculateAverageRainfall()
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 5:
			var threshold float64
			var err error
			var filteredCities []models.Climate
			fmt.Print("Input the threshold in mm: ")
			fmt.Scanln(&threshold)
			filteredCities, err = climate.FilterCitiesByRainfall(threshold)
			if err != nil {
				fmt.Println("Error:", err)
			}
			for _, city := range filteredCities {
				fmt.Printf("City : %s\n", city.Name)
				fmt.Printf("Rainfall: %.2f mm\n", city.Rainfall)
			}
		case 6:
			var name string
			fmt.Print("Input City Name: ")
			reader := bufio.NewReader(os.Stdin)
			input, readErr := reader.ReadString('\n')
			if readErr != nil {
				fmt.Println("Error occured in reading name. Error message : ", readErr)
				name = "Default City"
			} else {
				name = strings.TrimSpace(input)
			}
			err := climate.SearchByCityName(name)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
