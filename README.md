# go-openweathermap <img src="https://img.icons8.com/color/48/000000/golang.png" height="25" width="25"/> <img src="https://img.icons8.com/external-flatart-icons-outline-flatarticons/64/000000/external-map-map-pin-flatart-icons-outline-flatarticons.png" height="25" width="25"/>

Go SDK for interacting with https://openweathermap.org/api

## Installation

```sh
go get github.com/DillonStreator/go-openweathermap
```

## Usage

```go
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
	// Uses environment variable from key OPENWEATHERMAP_APPID
	client, err := openweathermap.NewHTTPClient(&http.Client{}, "")
	if err != nil {
		log.Fatal(err.Error())
	}

	city, err := client.GetByCityName(context.Background(), "chicago")
	if err != nil {
		log.Fatal(err.Error())
	}

	b, _ := json.Marshal(city)
	fmt.Println(string(b))
}
```

[Examples](./example/main.go)
