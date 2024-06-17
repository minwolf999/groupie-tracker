package filter

import (
	"io/ioutil"
	"strings"

	data "test/data"
)

// Filter to find the continent of a country from a ListePays database
func FilterPays(fileName string, artistsObject data.Art) data.Art {
	var temp3 data.Art
	skip := false
	temp1, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	NA := strings.Split(string(temp1), " ")
	for i := 0; i < len(artistsObject.Artists); i++ {
		temp2 := strings.Split(artistsObject.Artists[i].Locations, "-")
		for a := 1; a < len(temp2); a++ {
			b := strings.IndexRune(temp2[a], ' ')
			for _, c := range NA {
				if temp2[a][0:b] == c {
					temp3.Artists = append(temp3.Artists, artistsObject.Artists[i])
					skip = true
					break
				}
			}
			if skip {
				break
			}
		}
	}
	return temp3
}
