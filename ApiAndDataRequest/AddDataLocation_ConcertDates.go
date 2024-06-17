package ApiAndDataRequest

import (
	"strings"
	data "test/data"
)

// Add data from Locations, ConcertDates
func AddDataLocation_ConcertDates(artistsObject data.Art, locationsObject data.Locations, datesObject data.Dates) data.Art {
	for i := 0; i < len(artistsObject.Artists); i++ {
		artistsObject.Artists[i].Locations = strings.Join(locationsObject[i].Locations, " :*") + " :*"
		for k := 0; k < len(datesObject[i].Dates); k++ {
			for _, c := range datesObject[i].Dates[k] {
				if c == '*' {
					datesObject[i].Dates[k] = "" + datesObject[i].Dates[k]
				}
			}
		}
		artistsObject.Artists[i].ConcertDates = strings.Join(datesObject[i].Dates, ", ")
	}

	return artistsObject
}
