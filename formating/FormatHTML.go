package formating

import (
	data "test/data"
)

// Add formating for HTML
func FormatHTML(artistsObject data.Art) data.Art {
	for i := range artistsObject.Artists {
		artistsObject.Artists[i].MembersString = ""
		for y := range artistsObject.Artists[i].Members {
			if y != len(artistsObject.Artists[i].Members)-1 {
				artistsObject.Artists[i].MembersString += artistsObject.Artists[i].Members[y] + "\n"
			} else {
				artistsObject.Artists[i].MembersString += artistsObject.Artists[i].Members[y]
			}
		}
	}
	return artistsObject
}
