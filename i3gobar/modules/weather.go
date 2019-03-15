package modules

import (
	"fmt"

	"github.com/zwzn/weather"
)

func Weather() string {
	w, err := weather.Load()

	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s, %v",
		w.ForecastGroup.Forcast[0].AbbreviatedForecast.Summary,
		w.CurrentConditions.Temperature.String())
}
