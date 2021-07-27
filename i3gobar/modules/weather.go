package modules

import (
	"fmt"

	"github.com/abibby/weather"
)

func Weather() string {
	w, err := weather.Load()

	if err != nil {
		return ""
	}

	forcast := w.ForecastGroup.Forcast[0]
	temp := forcast.Temperature
	if (forcast.Humidex != weather.Unit{}) {
		temp = forcast.Humidex
	}

	// if (forcast.WindChill != weather.Unit{}) {
	// 	temp = forcast.WindChill
	// }

	return fmt.Sprintf("%s, %s",
		w.ForecastGroup.Forcast[0].AbbreviatedForecast.Summary,
		temp.String(),
	)
}
