package space

// Planet is the location where age will be calculated
type Planet string

// Define number of seconds of 1 Earth year
const earthYearInSeconds = 31557600

// Map planet orbital ratio as compared to Earth
var orbitalRatio = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age calculates number of years in the given planet
func Age(seconds float64, planet Planet) float64 {
	return seconds / (orbitalRatio[planet] * earthYearInSeconds)
}
