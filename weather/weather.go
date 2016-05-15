package weather;

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
	"github.com/ngraves95/run-temp/tokens"
	"github.com/ngraves95/util"
)

const baseUrl string = "http://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s";

const indicator string = "Temperature";

// Checks if a description contains weather data or not, as added by this.
func ContainsWeatherData(desc string) bool {

	return strings.Contains(desc, indicator);
}

// Adds the weather data to the description and returns the new description.
func AddWeatherData(lat float64, long float64, desc string) string {
	humidity, temperature := GetCurrentHumidityAndTempFarenheit(lat, long);

	return fmt.Sprintf("%s: %.2f oF\n%s: %.0f%%\n%s",
		indicator, temperature, "Humidity", humidity, desc);
}

func formatWeatherUrl(lat float64, long float64, key string) string {
	return fmt.Sprintf(baseUrl, lat, long, key);
}

type weatherResponseRoot struct {
	main map[string]float64 `json:"main"`
}

// TODO: implement via weather api.
func GetCurrentHumidityAndTempFarenheit(lat float64, long float64) (float64, float64) {
	tokenManager := tokens.NewTokenManager("/home/ngraves3/gocode/src/github.com/ngraves95/run-temp/tokens/");
	var weatherToken string = tokenManager.GetToken("openweathermap");

	resp, err := http.Get(formatWeatherUrl(lat, long, weatherToken));
	util.Check(err);
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	util.Check(err);

	var f interface{};
	err = json.Unmarshal(contents, &f);
	util.Check(err);

	m := f.(map[string]interface{});
	relevantParts := m["main"].(map[string]interface{});
	temperature := relevantParts["temp"].(float64);
	humidity := relevantParts["humidity"].(float64);

	// Conver from Kelvin to Farenheit
	temperature = (temperature * (9.0/5.0)) - 459.67;

	return humidity, temperature;
}
