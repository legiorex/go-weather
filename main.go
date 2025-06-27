package main

import (
	"flag"
	"fmt"

	"github.com/legiorex/go-weather/geo"
)

func main() {

	city := flag.String("city", "", "Город пользователя")

	// format := flag.Int("format", 1, "Формат вывода")

	flag.Parse()

	fmt.Println(*city)
	// fmt.Println(*format)

	geoData, err := geo.GetLocation(*city)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(geoData)

}
