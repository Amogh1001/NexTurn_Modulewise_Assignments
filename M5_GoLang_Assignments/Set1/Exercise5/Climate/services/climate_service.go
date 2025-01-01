package services

import (
	"climate-analysis/Climate/models"
	"errors"
	"fmt"
	"strings"
)

var Climates []models.Climate

func AddInput(id int, name string, avgtemp float64, rainfall float64) error {
	for _, climate := range Climates {
		if climate.ID == id {
			return errors.New("climate ID must be unique")
		}
	}

	climate := models.Climate{ID: id, Name: name, AvgTemperature: avgtemp, Rainfall: rainfall}

	Climates = append(Climates, climate)

	fmt.Printf("Climate for city added: %v\n", climate)

	return nil
}

func HighestTemperature() error {
	if len(Climates) == 0 {
		return errors.New("currently no data available")
	}
	maxAvgTemp := Climates[0].AvgTemperature
	city := Climates[0].Name

	for _, climate := range Climates {
		if climate.AvgTemperature > maxAvgTemp {
			maxAvgTemp = climate.AvgTemperature
			city = climate.Name
		}
	}

	fmt.Printf("City with the highest temperature (%.2f °C) is : %s", maxAvgTemp, city)

	return nil
}

func LowestTemperature() error {
	if len(Climates) == 0 {
		return errors.New("currently no data available")
	}
	minAvgTemp := Climates[0].AvgTemperature
	city := Climates[0].Name

	for _, climate := range Climates {
		if climate.AvgTemperature < minAvgTemp {
			minAvgTemp = climate.AvgTemperature
			city = climate.Name
		}
	}

	fmt.Printf("City with the lowest temperature (%.2f °C) is : %s\n", minAvgTemp, city)

	return nil
}

func CalculateAverageRainfall() error {
	if len(Climates) == 0 {
		return errors.New("currently no data available")
	}
	var totalRainfall float64
	for _, climate := range Climates {
		totalRainfall += climate.Rainfall
	}

	totalAvgRainfall := totalRainfall / float64(len(Climates))

	fmt.Printf("Average Rainfall: %.2f mm\n", totalAvgRainfall)

	return nil
}

func FilterCitiesByRainfall(threshold float64) ([]models.Climate, error) {
	if threshold < 0 {
		return nil, errors.New("invalid threshold value. Threshold cannot be negative")
	}

	var filteredCities []models.Climate
	for _, climate := range Climates {
		if climate.Rainfall >= threshold {
			filteredCities = append(filteredCities, climate)
		}
	}

	if len(filteredCities) == 0 {
		return nil, errors.New("no cities with rainfall above the threshold")
	}

	return filteredCities, nil
}

func SearchByCityName(name string) error {
	for _, climate := range Climates {
		if strings.Contains(climate.Name, name) {
			fmt.Printf("City found: %s\n", climate.Name)
			fmt.Printf("Average Temperature: %.2f°C\n", climate.AvgTemperature)
			fmt.Printf("Rainfall: %.2f mm\n", climate.Rainfall)
			return nil
		}
	}
	return errors.New("city not found")
}
