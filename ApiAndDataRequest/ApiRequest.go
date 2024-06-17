package ApiAndDataRequest

import (
	"encoding/json"
	"io"
	"net/http"

	data "test/data"
)

// Make the API requests for artists, locations and dates
func ApiRequest() (data.Art, data.Locations, data.Dates) {
	artists, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	defer artists.Body.Close()

	locations, _ := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	defer locations.Body.Close()

	dates, _ := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	defer dates.Body.Close()

	// Read the response bodies
	artistsByte, _ := io.ReadAll(artists.Body)

	var artistsObject data.Art
	json.Unmarshal(artistsByte, &artistsObject.Artists)

	locationsByte, _ := io.ReadAll(locations.Body)

	var locationsObject data.Locations
	json.Unmarshal(locationsByte[9:len(locationsByte)-2], &locationsObject)

	datesByte, _ := io.ReadAll(dates.Body)

	var datesObject data.Dates
	json.Unmarshal(datesByte[9:len(datesByte)-2], &datesObject)

	return artistsObject, locationsObject, datesObject
}
