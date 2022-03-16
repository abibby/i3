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

	current := w.CurrentConditions
	forcast := w.ForecastGroup.Forcast[0]

	temp := current.Temperature
	if (current.Humidex != weather.Unit{}) {
		temp = current.Humidex
	}

	if (temp == weather.Unit{}) {
		temp = forcast.Temperature
		if (forcast.Humidex != weather.Unit{}) {
			temp = forcast.Humidex
		}
	}

	condition := current.Condition
	if condition == "" {
		condition = w.ForecastGroup.Forcast[0].AbbreviatedForecast.Summary
	}

	// if (forcast.WindChill != weather.Unit{}) {
	// 	temp = forcast.WindChill
	// }

	return fmt.Sprintf("%s, %s",
		condition,
		temp.String(),
	)
}
