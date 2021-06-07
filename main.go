package main

import (
	"log"
	"fmt"
	"os"
	"net/http"
	"net/url"
	"io/ioutil"
)

func getWeather(city, owApiHost, owApiKey string) string {
	req, err := http.NewRequest("GET", owApiHost, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("appid", owApiKey)
	q.Add("q", city)
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
		bodyString := string(bodyBytes)
		return bodyString
	}
	return ""

}


func main() {
	owApiHost := "http://api.openweathermap.org/data/2.5/weather"
	owApiKey := os.Getenv("OW_API_KEY")
	city := "Pittsburgh"


	result := getWeather(city, owApiHost, owApiKey)
	fmt.Println(result)

}

