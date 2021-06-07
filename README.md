# Go Weather Forecast

Go Weather Forecast was inspired from [Python Weather Forecast](https://github.com/gonzaloplaza/python-weather-forecast). However this application does not leverage the same API's or have the exact capability. 

The Go Weather Forecast leverages the [OpenWeather API](https://openweathermap.org/current) using any `City` as a parameter. 

## Requirements

- `go1.14.15`
- [OpenWeather API Key](https://openweathermap.org/)

## Installation (Linux) and Usage
Declare the environment variable for your API key to Open Weather as so.

```bash
$ export OW_API_KEY=<your_api_key_from_openweathermap>
```
Clone this project:
```bash
$ git clone https://github.com/spensireli/go-weather-forecast
```

Run `main.go` passing a city.

```bash
$ go run main.go -city 
$ go run main.go -city=Pittsburgh
The city is: Pittsburgh 
The temperature is: 72.59 
The humidity is: 76 
```

## Licensing

The code in this project is licensed under [MIT LICENSE](https://github.com/spensireli/go-weather-forecast/blob/master/LICENSE). Read file for more information.
