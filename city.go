package openweathermap

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

type City struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CityResponse struct {
	City
	Cod int `json:"cod"`
}

type CitiesRectangleResponse struct {
	Cod      int     `json:"cod"`
	Calctime float32 `json:"calctime"`
	Cnt      int     `json:"cnt"`
	List     []City  `json:"list"`
}

type CitiesCircleResponse struct {
	Cod      string  `json:"cod"`
	Calctime float32 `json:"calctime"`
	Cnt      int     `json:"count"`
	List     []City  `json:"list"`
}

func (h *HTTPClient) GetByCityName(ctx context.Context, cityName string) (*CityResponse, error) {
	cityResponse := &CityResponse{}

	values := url.Values{
		"q": []string{cityName},
	}
	_, err := h.get(ctx, "/weather", values, cityResponse)

	return cityResponse, err
}

func (h *HTTPClient) GetByCityID(ctx context.Context, cityID int) (*CityResponse, error) {
	cityResponse := &CityResponse{}

	values := url.Values{
		"id": []string{strconv.Itoa(cityID)},
	}
	_, err := h.get(ctx, "/weather", values, cityResponse)

	return cityResponse, err
}

type Point struct {
	Lat string
	Lon string
}

func (h *HTTPClient) GetByGeographicCoordinate(ctx context.Context, point Point) (*CityResponse, error) {
	cityResponse := &CityResponse{}

	values := url.Values{
		"lat": []string{point.Lat},
		"lon": []string{point.Lon},
	}
	_, err := h.get(ctx, "/weather", values, cityResponse)

	return cityResponse, err
}

func (h *HTTPClient) GetByZIPCode(ctx context.Context, zipCode, countryCode string) (*CityResponse, error) {
	cityResponse := &CityResponse{}

	values := url.Values{
		"zip": []string{strings.Join([]string{zipCode, countryCode}, ",")},
	}
	_, err := h.get(ctx, "/weather", values, cityResponse)

	return cityResponse, err
}

type BoundingBox struct {
	LonLeft   string
	LonRight  string
	LatTop    string
	LatBottom string
	Zoom      int
}

func (h *HTTPClient) GetCitiesWithinARectangleZone(ctx context.Context, bbox BoundingBox) (*CitiesRectangleResponse, error) {
	citiesResponse := &CitiesRectangleResponse{}

	values := url.Values{
		"bbox": []string{strings.Join([]string{bbox.LonLeft, bbox.LatBottom, bbox.LonRight, bbox.LatTop, strconv.Itoa(bbox.Zoom)}, ",")},
	}
	_, err := h.get(ctx, "/box/city", values, citiesResponse)

	return citiesResponse, err

}

type BoundingPoint struct {
	Point
	Count int
}

func (h *HTTPClient) GetCitiesInCircle(ctx context.Context, bpoint BoundingPoint) (*CitiesCircleResponse, error) {
	citiesResponse := &CitiesCircleResponse{}

	values := url.Values{
		"lat": []string{bpoint.Lat},
		"lon": []string{bpoint.Lon},
		"cnt": []string{strconv.Itoa(bpoint.Count)},
	}
	_, err := h.get(ctx, "/find", values, citiesResponse)

	return citiesResponse, err

}
