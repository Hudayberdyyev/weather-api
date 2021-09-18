package models

type Coordinates struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	FeelsLike float64 `json:"feels_like"`
	Pressure  float64 `json:"pressure"`
	SeaLevel  float64 `json:"sea_level,omitempty"`
	GrndLevel float64 `json:"grnd_level,omitempty"`
	Humidity  int     `json:"humidity"`
	TempKf    float64 `json:"temp_kf,omitempty"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg,omitempty"`
	Gust  float64 `json:"gust,omitempty"`
}

type Clouds struct {
	All int `json:"all"`
}

type Rain struct {
	OneH   float64 `json:"1h,omitempty"`
	ThreeH float64 `json:"3h,omitempty"`
}

type Snow struct {
	OneH   float64 `json:"1h,omitempty"`
	ThreeH float64 `json:"3h,omitempty"`
}

type Sys struct {
	Pod string `json:"pod"`
}

type City struct {
	Id         int         `json:"id"`
	Name       string      `json:"name"`
	Coord      Coordinates `json:"coord"`
	Country    string      `json:"country"`
	Population int         `json:"population,omitempty"`
	Timezone   int         `json:"timezone"`
	Sunrise    int         `json:"sunrise"`
	Sunset     int         `json:"sunset"`
}

type WeatherInfo struct {
	Dt          int       `json:"dt"`
	MainData    Main      `json:"main"`
	WeatherData []Weather `json:"weather"`
	CloudsData  Clouds    `json:"clouds"`
	WindData    Wind      `json:"wind"`
	Visibility  int       `json:"visibility,omitempty"`
	Pop         float64   `json:"pop"`
	RainData    Rain      `json:"rain,omitempty"`
	SnowData    Snow      `json:"snow,omitempty"`
	SysData     Sys       `json:"sys,omitempty"`
	DtTxt       string    `json:"dt_txt"`
}

type OwmResponse struct {
	Cod      string        `json:"cod"`
	Message  int           `json:"message"`
	Cnt      int           `json:"cnt"`
	ListData []WeatherInfo `json:"list"`
	CityData City          `json:"city"`
}

type Regions struct {
	RegionId int
	Name     string
}
