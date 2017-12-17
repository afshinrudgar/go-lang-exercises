package space

type Planet string

// - Earth: orbital period 365.25 Earth days, or 31557600 seconds
const EarthYearInSeconds float64 = 31557600

var planets map[Planet]float64 = map[Planet]float64{
	"Mercury": 0.2408467 * EarthYearInSeconds,
	"Earth":   1.0 * EarthYearInSeconds,
	"Venus":   0.61519726 * EarthYearInSeconds,
	"Mars":    1.8808158 * EarthYearInSeconds,
	"Jupiter": 11.862615 * EarthYearInSeconds,
	"Saturn":  29.447498 * EarthYearInSeconds,
	"Uranus":  84.016846 * EarthYearInSeconds,
	"Neptune": 164.79132 * EarthYearInSeconds,
}

func Round(f float64) float64 {
	return float64(int64(f*1e2+.5)) / 1e2
}

func Age(seconds float64, planetName Planet) float64 {
	yearInSeconds, _ := planets[planetName]
	return Round(seconds / yearInSeconds)
}
