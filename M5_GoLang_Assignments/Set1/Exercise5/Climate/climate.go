package climate

import (
	"climate-analysis/Climate/models"
	"climate-analysis/Climate/services"
)

func AddInput(id int, name string, avgtemp float64, rainfall float64) error {
	return services.AddInput(id, name, avgtemp, rainfall)
}

func HighestTemperature() error {
	return services.HighestTemperature()
}

func LowestTemperature() error {
	return services.LowestTemperature()
}

func CalculateAverageRainfall() error {
	return services.CalculateAverageRainfall()
}

func FilterCitiesByRainfall(threshold float64) ([]models.Climate, error) {
	return services.FilterCitiesByRainfall(threshold)
}

func SearchByCityName(name string) error {
	return services.SearchByCityName(name)
}
