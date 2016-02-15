package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var city string

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	flag.StringVar(&city, "city", "Portland", "the city to get the weather for")
	flag.Parse()

	data, err := query(city)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current weather in %v is: ", data.Name)
	fmt.Printf("%d°F\n", int64(tempInF(data.Main.Kelvin)))
}

// Queries the OpenWeatherMap API for the passed in city. Returns
// weatherData info and an error.
func query(city string) (weatherData, error) {
	owmAPIKey := os.Getenv("WTHR_OWM_KEY")
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + owmAPIKey + "&q=" + city)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	return d, nil
}

// Converts the passed in Kelvin temperature to Fahrenheit.
func tempInF(tempInK float64) float64 {
	return (tempInK-273.15)*1.8 + 32.0
}
