package openweathermap

import "context"

// Client describes a client for the openweathermap API
type Client interface {
	GetByCityName(ctx context.Context, cityName string) (*CityResponse, error)
	GetByCityID(ctx context.Context, cityID int) (*CityResponse, error)
	GetByGeographicCoordinate(ctx context.Context, point Point) (*CityResponse, error)
	GetByZIPCode(ctx context.Context, zipCode, countryCode string) (*CityResponse, error)
	GetCitiesWithinARectangleZone(ctx context.Context, bbox BoundingBox) (*CitiesRectangleResponse, error)
	GetCitiesInCircle(ctx context.Context, bpoint BoundingPoint) (*CitiesCircleResponse, error)
}
