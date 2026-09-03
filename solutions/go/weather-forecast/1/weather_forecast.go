//Package weather is intended to help tell the weather information for the country of Goblinocus. 
package weather


var (
    //CurrentCondition describes the current weather conditions.
	CurrentCondition string 
    //CurrentLocation describes the current location or city the weather program is intended to information for.
	CurrentLocation  string 
)

//Forecast function describes the current weather condition basedon the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
