package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/legiorex/go-weather/geo"
)

func GetWeather(geo geo.GeoData, format int) string {

	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	params := url.Values{}

	params.Add("format", fmt.Sprint(format))

	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	if resp.StatusCode != 200 {

		fmt.Println(errors.New("ERROR_STATUS"))
		return ""
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	defer resp.Body.Close()

	return string(body)

}
