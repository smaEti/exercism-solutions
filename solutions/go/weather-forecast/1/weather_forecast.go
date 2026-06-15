// Package weather provides functionality for weather forecasting.
package weather

var (
	// CurrentCondition stores the current weather condition.
	CurrentCondition string

	// CurrentLocation stores the current location.
	CurrentLocation string
)

// Forecast returns a formatted string containing the location and weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}