package main

import (
	"log"
	"fmt"
	"os"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"flag"
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

func getWeather(city, owApiHost, owApiKey string) string {

	bodyString := ""
	req, err := http.NewRequest("GET", owApiHost, nil)
	if err != nil {
		log.Print(err)
	}

	q := url.Values{}
	q.Add("appid", owApiKey)
	q.Add("q", city)
	q.Add("units", "imperial")
	req.URL.RawQuery = q.Encode()
	requestUri := req.URL.String()

	resp, err := http.Get(requestUri)
	if err != nil {
		log.Fatalln(err)
	  }
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString = string(bodyBytes)
		return bodyString
	}
	return bodyString

}


func main() {
	var city string
	owApiHost := "http://api.openweathermap.org/data/2.5/weather"
	owApiKey := os.Getenv("OW_API_KEY")
	wordPtr := flag.String("city", "foo", "a string")
	flag.Parse()
	city = *wordPtr
	var results Results
	result := getWeather(city, owApiHost, owApiKey)
	json.Unmarshal([]byte(result), &results)
	
	fmt.Printf("The city is: %s \n", results.Name)
	fmt.Printf("The temperature is: %g \n", results.Main.Temp)
	fmt.Printf("The humidity is: %d \n", results.Main.Humidity)
	fmt.Printf("The forecast is: %s", results.Weather[0].Description)
}

