// Package weather introduces handy vars and functions for weather forcasting.
package weather

var (
	// CurrentCondition variable conains text for current weather conditions.
	CurrentCondition string
	// CurrentLocation shows a geographical point where weather is presented from.
	CurrentLocation string
)

// Forecast function accepts location and current conditions so that it can use it to do a forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
