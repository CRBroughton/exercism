// Package weather - Forecasts the current weather conditions for a given city.
// Provide the Forecast function with a city name and condition
// to recieve the forcast.
package weather

// CurrentCondition - Sets the current weather condition.
var CurrentCondition string

// CurrentLocation - Set the current weather location.
var CurrentLocation string

// Forecast - Returns the Forecast.
// Requires a city and condition arguments.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
