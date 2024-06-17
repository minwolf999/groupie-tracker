package filter

import (
	"strconv"
	data "test/data"
)

// Sort by filters choose
func FilterChoose(artistsObject data.Art, doc, doc1, doc2, yearfirst, yearlast, fa, fa2, nm, nm2 string) []struct {
	ID            int      "json:\"id\""
	Image         string   "json:\"image\""
	Name          string   "json:\"name\""
	Members       []string "json:\"members\""
	CreationDate  int      "json:\"creationDate\""
	FirstAlbum    string   "json:\"firstAlbum\""
	Locations     string   "json:\"locations\""
	ConcertDates  string   "json:\"concertDates\""
	MembersString string
	TabLocations  []string
} {
	var temp data.Art
	for i := range artistsObject.Artists {
		if doc == "" && doc1 == "" && doc2 == "" {
			break
		}
		yearfirstint, _ := strconv.Atoi(yearfirst)
		yearlastint, _ := strconv.Atoi(yearlast)

		nmInt, _ := strconv.Atoi(nm)
		nm2Int, _ := strconv.Atoi(nm2)

		faint, _ := strconv.Atoi(fa)
		faint2, _ := strconv.Atoi(fa2)

		FirstAlbumDate, _ := strconv.Atoi(artistsObject.Artists[i].FirstAlbum[len(artistsObject.Artists[i].FirstAlbum)-4:])

		if doc == "0" && doc1 == "" && doc2 == "" {
			if artistsObject.Artists[i].CreationDate >= yearfirstint && artistsObject.Artists[i].CreationDate <= yearlastint {
				temp.Artists = append(temp.Artists, artistsObject.Artists[i])
			}
		}
		if doc == "0" && doc1 == "0" && doc2 == "" {
			if artistsObject.Artists[i].CreationDate >= yearfirstint && artistsObject.Artists[i].CreationDate <= yearlastint && FirstAlbumDate >= faint && FirstAlbumDate <= faint2 {
				temp.Artists = append(temp.Artists, artistsObject.Artists[i])
			}
		}
		if doc == "0" && doc1 == "0" && doc2 == "0" {
			if artistsObject.Artists[i].CreationDate >= yearfirstint && artistsObject.Artists[i].CreationDate <= yearlastint && FirstAlbumDate >= faint && FirstAlbumDate <= faint2 && len(artistsObject.Artists[i].Members) >= nmInt && len(artistsObject.Artists[i].Members) <= nm2Int {
				temp.Artists = append(temp.Artists, artistsObject.Artists[i])
			}
		}
		if doc == "" && doc1 == "0" && doc2 == "0" {
			if FirstAlbumDate >= faint && FirstAlbumDate <= faint2 && len(artistsObject.Artists[i].Members) >= nmInt && len(artistsObject.Artists[i].Members) <= nm2Int {
				temp.Artists = append(temp.Artists, artistsObject.Artists[i])
			}
		}
		if doc == "" && doc1 == "" && doc2 == "0" {
			if len(artistsObject.Artists[i].Members) >= nmInt && len(artistsObject.Artists[i].Members) <= nm2Int {
				temp.Artists = append(temp.Artists, artistsObject.Artists[i])
			}
		}
		if doc == "0" && doc1 == "" && doc2 == "0" {
			if artistsObject.Artists[i].CreationDate >= yearfirstint && artistsObject.Artists[i].CreationDate <= yearlastint && len(artistsObject.Artists[i].Members) >= nmInt && len(artistsObject.Artists[i].Members) <= nm2Int {
				temp.Artists = append(temp.Artists, artistsObject.Artists[i])
			}
		}
		if doc == "" && doc1 == "0" && doc2 == "" {
			if FirstAlbumDate >= faint && FirstAlbumDate <= faint2 {
				temp.Artists = append(temp.Artists, artistsObject.Artists[i])
			}
		}
	}
	if doc == "0" || doc1 == "0" || doc2 == "0" {
		return temp.Artists
	} else {
		return artistsObject.Artists
	}
}
