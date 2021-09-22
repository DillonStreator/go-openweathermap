package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/DillonStreator/go-openweathermap"
)

func main() {
	client, err := openweathermap.NewHTTPClient(&http.Client{}, "") // Uses environment variable from key OPENWEATHERMAP_APPID
	if err != nil {
		log.Fatal(err.Error())
	}

	cityResponse, err := client.GetByCityName(context.Background(), "chicago")
	if err != nil {
		log.Fatal(err.Error())
	}

	cBytes, _ := json.Marshal(cityResponse)
	fmt.Println(string(cBytes))

	citiesRectangle, err := client.GetCitiesWithinARectangleZone(context.Background(), openweathermap.BoundingBox{
		LonLeft:   "12",
		LatBottom: "32",
		LonRight:  "15",
		LatTop:    "37",
		Zoom:      10,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	crBytes, _ := json.Marshal(citiesRectangle)
	fmt.Println(string(crBytes))

	citiesCircle, err := client.GetCitiesInCircle(context.Background(), openweathermap.BoundingPoint{
		Count: 10,
		Point: openweathermap.Point{
			Lat: "55.5",
			Lon: "37.5",
		},
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	ccBytes, _ := json.Marshal(citiesCircle)
	fmt.Println(string(ccBytes))
}
