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

	temp := w.CurrentConditions.Temperature
	if (w.CurrentConditions.Humidex != weather.Unit{}) {
		temp = w.CurrentConditions.Humidex
	}

	if (w.CurrentConditions.WindChill != weather.Unit{}) {
		temp = w.CurrentConditions.WindChill
	}

	return fmt.Sprintf("%s, %s",
		w.ForecastGroup.Forcast[0].AbbreviatedForecast.Summary,
		temp.String(),
	)
}
