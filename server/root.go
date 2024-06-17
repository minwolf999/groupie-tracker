package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	request "test/ApiAndDataRequest"
	"test/add"
	data "test/data"
	"test/errors"
	filter "test/filterpays"
	"test/formating"
)

var Init = 0
var ArtistsObject data.Art
var MapLocation map[string]string

// Take the html's datas, dispatch them inside a []structure and create the server
func Root(writer http.ResponseWriter, r *http.Request) {
	start := time.Now()

	if r.URL.Path != "/" && r.URL.Path != "/web" {
		errors.Error(writer, http.StatusNotFound)
		return
	}

	// Request all HTML datas
	var na, sa, eu, me, as, oc, doc, doc1, doc2, yearfirst, yearlast, fa, fa2, nm, nm2, cityState, searchBar string
	if Init != 0 {
		_, na, sa, eu, me, as, oc, doc, doc1, doc2, yearfirst, yearlast, fa, fa2, nm, nm2, cityState, searchBar = request.RecupDataHtml(writer, r)
	}
	Init++

	// ArtistsObject, _, _ := request.ApiRequest()

	ArtistsObject.Artists = filter.FilterChoose(ArtistsObject, doc, doc1, doc2, yearfirst, yearlast, fa, fa2, nm, nm2)

	ArtistsObject = filter.ContinentFilter(ArtistsObject, na, sa, eu, me, as, oc)

	if cityState != "" {
		ArtistsObject = filter.SearchBarCity(ArtistsObject, cityState)
	}

	ArtistsObject = formating.FormatHTML(ArtistsObject)

	// Global searchBar filter
	var temp3 data.Art
	if searchBar != "" {
		temp3 = filter.SearchBarFilter(ArtistsObject, temp3, searchBar)
	}

	mt := add.Add(ArtistsObject)

	// Take HTML datas to HTML pages for searchBar suggestion
	var te data.Art
	te.MegaTableauJVscript.Truc = append(te.MegaTableauJVscript.Truc, mt...)

	if searchBar == "" {
		te.Artists = ArtistsObject.Artists
	} else {
		te.Artists = temp3.Artists
	}

	// Take locations names for geolocalize
	for i := range ArtistsObject.Artists {
		temp4 := strings.ReplaceAll(strings.ReplaceAll(ArtistsObject.Artists[i].Locations, ":*", ""), "<br>", "")
		ArtistsObject.Artists[i].TabLocations = strings.Split(temp4[:len(temp4)-1], " ")

		for y := range ArtistsObject.Artists[i].TabLocations {
			ArtistsObject.Artists[i].TabLocations[y] = MapLocation[ArtistsObject.Artists[i].TabLocations[y]]
		}
	}

	end := time.Now()
	fmt.Println(end.Sub(start))

	// Execute the site
	t, err := template.ParseFiles("./templates/home.html")
	errors.Check(writer, err)

	err = t.Execute(writer, te)
	errors.Check(writer, err)
}
