package geo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"cityName"`
}

func GetLocation(city string) (*GeoData, error) {

	if city != "" {
		return &GeoData{
			City: city,
		}, nil
	}

	// resp, err := http.Get("https://ipapi.co/json/")
	resp, err := http.Get("https://free.freeipapi.com/api/json")

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("STATUS_ERROR")
	}

	var geo GeoData

	json.Unmarshal(body, &geo)

	return &geo, nil
}
