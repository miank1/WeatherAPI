package weather

type WeatherService interface {
	GetCurrentWeather(city string) (WeatherData, error)
}
