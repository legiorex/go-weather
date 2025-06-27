package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type CityResponse struct {
	Error bool `json:"error"`
}

type GeoData struct {
	City string `json:"cityName"`
}

func GetLocation(city string) (*GeoData, error) {

	if city != "" {

		isCity := CheckCity(city)

		if !isCity {
			panic("Такого города нет!")
		}

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

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("STATUS_ERROR")
	}

	var geo GeoData

	json.Unmarshal(body, &geo)

	return &geo, nil
}

func CheckCity(city string) bool {

	bodyData := map[string]string{
		"city": city,
	}

	body, _ := json.Marshal(bodyData)

	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(body))

	if err != nil {
		return true
	}

	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)

	var isError CityResponse

	json.Unmarshal(body, &isError)

	return !isError.Error

}
