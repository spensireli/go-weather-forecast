package main

import (
	"encoding/json"
	"fmt"
)

type Results struct {
	Coord Coord
	Weather []Weather
	Base string
	Main Main
	Visibility int
	Wind Wind
	Clouds Clouds
	Dt string
	Sys Sys
	Timezone int
	Id int
	Name string
	Cod int
}

type Coord struct {
	Lon float32
	Lat float32
}

type Weather struct {
	Id int
	Main string
	Description string
	Icon string
}

type Main struct {
	Temp float32
	FeelsLike float32 
	TempMin float32 
	TempMax float32 
	Pressure int
	Humidity int
}

type Wind struct {
	Speed float32 
	Deg int 
	Gust float32 
}

type Clouds struct {
	All int
}

type Sys struct {
	Type int 
	Id int
	Country string 
	Sunrise int
	Sunset int
}

func main() {
	myJsonString := `{"coord":{"lon":-79.9959,"lat":40.4406},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"base":"stations","main":{"temp":300.04,"feels_like":301.09,"temp_min":295.63,"temp_max":302.01,"pressure":1018,"humidity":60},"visibility":10000,"wind":{"speed":1.18,"deg":164,"gust":1.2},"clouds":{"all":1},"dt":1623026345,"sys":{"type":1,"id":3247,"country":"US","sunrise":1622973015,"sunset":1623026849},"timezone":-14400,"id":5206379,"name":"Pittsburgh","cod":200}`

	var results Results

	json.Unmarshal([]byte(myJsonString), &results)
	fmt.Println(results)



}

