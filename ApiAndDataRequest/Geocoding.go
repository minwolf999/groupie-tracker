package ApiAndDataRequest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	data "test/data"
)

// convert address in coordinates gps
func GeoCode(tab string, chanSender chan string) {
	if tab == "playa_del_carmen-mexico" {
		tab = "playa_del_carmen"
	}

	url := "https://api.geoapify.com/v1/geocode/search?text=" + strings.ReplaceAll(strings.ReplaceAll(tab, "-", " "), "_", " ") + "&apiKey=bef0bdf320fb423f9745fd06cbd824cd"

	if tab == "playa_del_carmen" {
		tab = "playa_del_carmen-mexico"
	}

	req, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var temp5 data.Coo
	err = json.Unmarshal(body, &temp5)
	if err != nil {
		panic(err)
	}

	city_state := strings.Split(strings.ReplaceAll(tab, "_", " "), "-")
	city_state[0] = RemoveAccent(city_state[0])
	city_state[1] = RemoveAccent(city_state[1])

	if city_state[1] == "usa" {
		city_state[1] = "unitedstates"
	} else if city_state[1] == "uk" {
		city_state[1] = "unitedkingdom"
	}

	for z := range temp5.Features {
		properties := temp5.Features[z].Properties

		if properties.Name == "" && properties.Timezone.Name != "" {
			properties.Name = properties.Timezone.Name
		}

		if (RemoveAccent(properties.Name) == city_state[0] || RemoveAccent(properties.State) == city_state[0] || RemoveAccent(properties.City) == city_state[0] || RemoveAccent(properties.Village) == city_state[0] || RemoveAccent(properties.Suburb) == city_state[0]) && (RemoveAccent(properties.Country) == city_state[1] || RemoveAccent(properties.Dependency) == city_state[1]) {
			tab = fmt.Sprintf("%v %v", temp5.Features[z].Geometry.Coordinates[1], temp5.Features[z].Geometry.Coordinates[0])
			break
		}
	}

	chanSender <- tab
}
