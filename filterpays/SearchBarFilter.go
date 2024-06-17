package filter

import (
	"strconv"
	"strings"
	data "test/data"
)

// Global searchBar filter
func SearchBarFilter(artistsObject, temp3 data.Art, searchBar string) data.Art {
	for i := range artistsObject.Artists {
		if len(searchBar) >= 16 && searchBar[len(searchBar)-16:] == "\u00a0-\u00a0Group\u00a0Name" && strings.Contains(artistsObject.Artists[i].Name, strings.ReplaceAll(searchBar[:len(searchBar)-16], "\u00a0", " ")) {
			temp3.Artists = append(temp3.Artists, artistsObject.Artists[i])
		} else if len(searchBar) >= 24 && searchBar[len(searchBar)-24:] == "\u00a0-\u00a0Group\u00a0First\u00a0Album" && strings.Contains(strings.ReplaceAll(artistsObject.Artists[i].FirstAlbum, "-", "/"), strings.ReplaceAll(searchBar[:len(searchBar)-24], "\u00a0", " ")) {
			temp3.Artists = append(temp3.Artists, artistsObject.Artists[i])
		} else if len(searchBar) >= 26 && searchBar[len(searchBar)-26:] == "\u00a0-\u00a0Group\u00a0Creation\u00a0Date" && strings.Contains(strconv.Itoa(artistsObject.Artists[i].CreationDate), strings.ReplaceAll(searchBar[:len(searchBar)-26], "\u00a0", " ")) {
			temp3.Artists = append(temp3.Artists, artistsObject.Artists[i])
		} else if len(searchBar) >= 20 && searchBar[len(searchBar)-20:] == "\u00a0-\u00a0Group\u00a0Location" && strings.Contains(artistsObject.Artists[i].Locations, strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(searchBar[:len(searchBar)-20]), "\u00a0", " "), " in the ", "-"), " ", "_")) {
			temp3.Artists = append(temp3.Artists, artistsObject.Artists[i])
		} else {
			for y := range artistsObject.Artists[i].Members {
				if len(searchBar) >= 18 && searchBar[len(searchBar)-18:] == "\u00a0-\u00a0Group\u00a0Member" && strings.Contains(artistsObject.Artists[i].Members[y], strings.ReplaceAll(searchBar[:len(searchBar)-18], "\u00a0", " ")) {
					temp3.Artists = append(temp3.Artists, artistsObject.Artists[i])
					break
				}
			}
		}
	}
	return temp3
}
