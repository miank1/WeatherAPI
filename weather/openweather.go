package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OpenWeatherClient struct {
	APIKey string
}

func NewOpenWeatherClient(apiKey string) *OpenWeatherClient {
	return &OpenWeatherClient{
		APIKey: apiKey,
	}
}

func (ow *OpenWeatherClient) GetCurrentWeather(city string) (WeatherData, error) {
	var data WeatherData

	fmt.Println("CITY -------", city)
	fmt.Println("ow.APIKey ---------------", ow.APIKey)
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&key=%s&contentType=json", city, ow.APIKey)

	fmt.Println("URL Is ------------>", url)
	resp, err := http.Get(url)
	if err != nil {
		return data, err
	}
	fmt.Println("Response is ----------->", resp.Body)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return data, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	weatherResp := WeatherResponse{}

	err = json.NewDecoder(resp.Body).Decode(&weatherResp)
	if err != nil {
		return data, err
	}

	data = WeatherData{
		City:        weatherResp.Address,
		Temperature: weatherResp.Days[0].Temp,
		Humidity:    int(weatherResp.Days[0].Humidity),
		Description: weatherResp.Days[0].Description,
		RainChances: weatherResp.CurrentConditions.Precipprob,
	}

	fmt.Println("Data is ---------->", data)

	return data, nil
}
