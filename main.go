package main

import (
	"fmt"
	"io"
	"net/http"
)

type CityWeather struct {
	QueryCost       int     `json:"queryCost"`
	Latitude        string  `json:"latitude"`
	Longitude       string  `json:"longitude"`
	ResolvedAddress string  `json:"resolvedAddress"`
	Address         string  `json:"address"`
	Timezone        string  `json:"timezone"`
	Tzoffset        float64 `json:"tzoffset"`
	Description     string  `json:"description"`
	Days            []struct {
		Datetime       string   `json:"datetime"`
		DatetimeEpoch  int      `json:"datetimeEpoch"`
		Tempmax        float64  `json:"tempmax"`
		Tempmin        float64  `json:"tempmin"`
		Temp           float64  `json:"temp"`
		Feelslikemax   float64  `json:"feelslikemax"`
		Feelslikemin   float64  `json:"feelslikemin"`
		Feelslike      float64  `json:"feelslike"`
		Dew            int      `json:"dew"`
		Humidity       float64  `json:"humidity"`
		Precip         int      `json:"precip"`
		Precipprob     int      `json:"precipprob"`
		Precipcover    int      `json:"precipcover"`
		Preciptype     any      `json:"preciptype"`
		Snow           int      `json:"snow"`
		Snowdepth      int      `json:"snowdepth"`
		Windgust       float64  `json:"windgust"`
		Windspeed      float64  `json:"windspeed"`
		Winddir        float64  `json:"winddir"`
		Pressure       float64  `json:"pressure"`
		Cloudcover     float64  `json:"cloudcover"`
		Visibility     float64  `json:"visibility"`
		Solarradiation int      `json:"solarradiation"`
		Solarenergy    int      `json:"solarenergy"`
		Uvindex        int      `json:"uvindex"`
		Severerisk     int      `json:"severerisk"`
		Sunrise        string   `json:"sunrise"`
		SunriseEpoch   int      `json:"sunriseEpoch"`
		Sunset         string   `json:"sunset"`
		SunsetEpoch    int      `json:"sunsetEpoch"`
		Moonphase      float64  `json:"moonphase"`
		Conditions     string   `json:"conditions"`
		Description    string   `json:"description"`
		Icon           string   `json:"icon"`
		Stations       []string `json:"stations"`
		Source         string   `json:"source"`
		Hours          []struct {
			Datetime       string   `json:"datetime"`
			DatetimeEpoch  int      `json:"datetimeEpoch"`
			Temp           float64  `json:"temp"`
			Feelslike      float64  `json:"feelslike"`
			Humidity       float64  `json:"humidity"`
			Dew            float64  `json:"dew"`
			Precip         int      `json:"precip"`
			Precipprob     int      `json:"precipprob"`
			Snow           int      `json:"snow"`
			Snowdepth      int      `json:"snowdepth"`
			Preciptype     any      `json:"preciptype"`
			Windgust       float64  `json:"windgust"`
			Windspeed      float64  `json:"windspeed"`
			Winddir        int      `json:"winddir"`
			Pressure       int      `json:"pressure"`
			Visibility     float64  `json:"visibility"`
			Cloudcover     int      `json:"cloudcover"`
			Solarradiation int      `json:"solarradiation"`
			Solarenergy    int      `json:"solarenergy"`
			Uvindex        int      `json:"uvindex"`
			Severerisk     int      `json:"severerisk"`
			Conditions     string   `json:"conditions"`
			Icon           string   `json:"icon"`
			Stations       []string `json:"stations"`
			Source         string   `json:"source"`
		} `json:"hours"`
	} `json:"days"`
	Alerts   []any `json:"alerts"`
	Stations struct {
		Vidp struct {
			Distance     int     `json:"distance"`
			Latitude     float64 `json:"latitude"`
			Longitude    float64 `json:"longitude"`
			UseCount     int     `json:"useCount"`
			ID           string  `json:"id"`
			Name         string  `json:"name"`
			Quality      int     `json:"quality"`
			Contribution int     `json:"contribution"`
		} `json:"VIDP"`
		Av559 struct {
			Distance     int     `json:"distance"`
			Latitude     float64 `json:"latitude"`
			Longitude    float64 `json:"longitude"`
			UseCount     int     `json:"useCount"`
			ID           string  `json:"id"`
			Name         string  `json:"name"`
			Quality      int     `json:"quality"`
			Contribution int     `json:"contribution"`
		} `json:"AV559"`
	} `json:"stations"`
	CurrentConditions struct {
		Datetime       string   `json:"datetime"`
		DatetimeEpoch  int      `json:"datetimeEpoch"`
		Temp           float64  `json:"temp"`
		Feelslike      float64  `json:"feelslike"`
		Humidity       float64  `json:"humidity"`
		Dew            float64  `json:"dew"`
		Precip         any      `json:"precip"`
		Precipprob     int      `json:"precipprob"`
		Snow           int      `json:"snow"`
		Snowdepth      int      `json:"snowdepth"`
		Preciptype     any      `json:"preciptype"`
		Windgust       any      `json:"windgust"`
		Windspeed      float64  `json:"windspeed"`
		Winddir        int      `json:"winddir"`
		Pressure       int      `json:"pressure"`
		Visibility     float64  `json:"visibility"`
		Cloudcover     int      `json:"cloudcover"`
		Solarradiation int      `json:"solarradiation"`
		Solarenergy    int      `json:"solarenergy"`
		Uvindex        int      `json:"uvindex"`
		Conditions     string   `json:"conditions"`
		Icon           string   `json:"icon"`
		Stations       []string `json:"stations"`
		Source         string   `json:"source"`
		Sunrise        string   `json:"sunrise"`
		SunriseEpoch   int      `json:"sunriseEpoch"`
		Sunset         string   `json:"sunset"`
		SunsetEpoch    int      `json:"sunsetEpoch"`
		Moonphase      float64  `json:"moonphase"`
	} `json:"currentConditions"`
}

func main() {
	fmt.Println("Weather API")

	// Note: You'll need to replace YOUR_API_KEY with your actual Visual Crossing API key
	url := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/delhi?unitGroup=us&key=X2AM8B666SQSMPZNKMD7QBEAE&contentType=json"

	// redis connection string

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making HTTP request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	// Print the response
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))

	// Start HTTP server
	http.HandleFunc("/weather", weatherHandler)
	fmt.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Weather API Server is running!")
}
