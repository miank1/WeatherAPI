package weather

type WeatherData struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Humidity    int     `json:"humidity"`
	Description string  `json:"description"`
	RainChances float64 `json:"rain_chances"`
}

type WeatherResponse struct {
	QueryCost       int     `json:"queryCost"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
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
		Dew            float64  `json:"dew"`
		Humidity       float64  `json:"humidity"`
		Precip         float64  `json:"precip"`
		Precipprob     float64  `json:"precipprob"`
		Precipcover    float64  `json:"precipcover"`
		Preciptype     any      `json:"preciptype"`
		Snow           float64  `json:"snow"`
		Snowdepth      float64  `json:"snowdepth"`
		Windgust       float64  `json:"windgust"`
		Windspeed      float64  `json:"windspeed"`
		Winddir        float64  `json:"winddir"`
		Pressure       float64  `json:"pressure"`
		Cloudcover     float64  `json:"cloudcover"`
		Visibility     float64  `json:"visibility"`
		Solarradiation float64  `json:"solarradiation"`
		Solarenergy    float64  `json:"solarenergy"`
		Uvindex        float64  `json:"uvindex"`
		Severerisk     float64  `json:"severerisk"`
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
			Precip         float64  `json:"precip"`
			Precipprob     float64  `json:"precipprob"`
			Snow           float64  `json:"snow"`
			Snowdepth      float64  `json:"snowdepth"`
			Preciptype     any      `json:"preciptype"`
			Windgust       float64  `json:"windgust"`
			Windspeed      float64  `json:"windspeed"`
			Winddir        float64  `json:"winddir"`
			Pressure       float64  `json:"pressure"`
			Visibility     float64  `json:"visibility"`
			Cloudcover     float64  `json:"cloudcover"`
			Solarradiation float64  `json:"solarradiation"`
			Solarenergy    float64  `json:"solarenergy"`
			Uvindex        float64  `json:"uvindex"`
			Severerisk     float64  `json:"severerisk"`
			Conditions     string   `json:"conditions"`
			Icon           string   `json:"icon"`
			Stations       []string `json:"stations"`
			Source         string   `json:"source"`
		} `json:"hours"`
	} `json:"days"`
	Alerts   []any `json:"alerts"`
	Stations struct {
		Vidp struct {
			Distance     float64 `json:"distance"`
			Latitude     float64 `json:"latitude"`
			Longitude    float64 `json:"longitude"`
			UseCount     int     `json:"useCount"`
			ID           string  `json:"id"`
			Name         string  `json:"name"`
			Quality      int     `json:"quality"`
			Contribution float64 `json:"contribution"`
		} `json:"VIDP"`
		Av559 struct {
			Distance     float64 `json:"distance"`
			Latitude     float64 `json:"latitude"`
			Longitude    float64 `json:"longitude"`
			UseCount     int     `json:"useCount"`
			ID           string  `json:"id"`
			Name         string  `json:"name"`
			Quality      int     `json:"quality"`
			Contribution float64 `json:"contribution"`
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
		Precipprob     float64  `json:"precipprob"`
		Snow           float64  `json:"snow"`
		Snowdepth      float64  `json:"snowdepth"`
		Preciptype     any      `json:"preciptype"`
		Windgust       any      `json:"windgust"`
		Windspeed      float64  `json:"windspeed"`
		Winddir        float64  `json:"winddir"`
		Pressure       float64  `json:"pressure"`
		Visibility     float64  `json:"visibility"`
		Cloudcover     float64  `json:"cloudcover"`
		Solarradiation float64  `json:"solarradiation"`
		Solarenergy    float64  `json:"solarenergy"`
		Uvindex        float64  `json:"uvindex"`
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
