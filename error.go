package openweathermap

import (
	"fmt"
	"net/http"
)

type apiError struct {
	Code         string `json:"cod"`
	Message      string `json:"message"`
	HTTPResponse *http.Response
}

func (a *apiError) Error() string {
	return fmt.Sprintf("API error %s %s", a.Code, a.Message)
}
