package filter

import data "test/data"

// SearchBar city filter
func SearchBarCity(artistsObject data.Art, cityState string) data.Art {
	var temp2 data.Art
	for i := range artistsObject.Artists {
		for j := 0; j < len(artistsObject.Artists[i].Locations); j++ {
			if j+len(cityState) < len(artistsObject.Artists[i].Locations) && artistsObject.Artists[i].Locations[j:j+len(cityState)] == cityState {
				temp2.Artists = append(temp2.Artists, artistsObject.Artists[i])
				break
			}
		}
	}
	return temp2
}
