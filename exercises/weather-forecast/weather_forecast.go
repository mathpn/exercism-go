// Package weather provides tools to forecast weather conditions.
package weather

// CurrentCondition represents the current weather condition for forecasting.
var CurrentCondition string

// CurrentLocation represents the current weather location for forecasting.
var CurrentLocation string

// Forecast returns a string with the current location and current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
