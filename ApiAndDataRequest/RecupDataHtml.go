package ApiAndDataRequest

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	errors "test/errors"
)

// Recup the data from HTML page
func RecupDataHtml(writer http.ResponseWriter, r *http.Request) (int, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string) {
	na := r.FormValue("North-America")
	if na != "" && na != "North-America" {
		errors.Error(writer, 500)
		return 1, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
	}

	sa := r.FormValue("South-America")
	if sa != "" && sa != "South-America" {
		errors.Error(writer, 500)
		return 1, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
	}

	eu := r.FormValue("Europe")
	if eu != "" && eu != "Europe" {
		errors.Error(writer, 500)
		return 1, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
	}

	me := r.FormValue("Middle-East")
	if me != "" && me != "Middle-East" {
		errors.Error(writer, 500)
		return 1, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
	}

	as := r.FormValue("Asia")
	if as != "" && as != "Asia" {
		errors.Error(writer, 500)
		return 1, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
	}

	oc := r.FormValue("Oceania")
	if oc != "" && oc != "Oceania" {
		errors.Error(writer, 500)
		return 1, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
	}

	doc := r.FormValue("Zero1")
	if doc != "" && doc != "0" {
		errors.Error(writer, 500)
		return 1, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
	}

	doc1 := r.FormValue("Zero2")
	if doc1 != "" && doc1 != "0" {
		errors.Error(writer, 500)
		return 1, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
	}

	doc2 := r.FormValue("Zero3")
	if doc2 != "" && doc2 != "0" {
		errors.Error(writer, 500)
		return 1, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
	}

	yearfirst := r.FormValue("Year")
	if yearfirst != "" {
		if yearint, err := strconv.Atoi(yearfirst); err != nil || (yearint < 1950 || yearint > 2020) {
			fmt.Println(err)
			errors.Error(writer, 500)
			return 1, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
		}
	}

	yearlast := r.FormValue("Yearlast")
	if yearlast != "" {
		if yearint, err := strconv.Atoi(yearlast); err != nil || (yearint < 1955 || yearint > 2023) {
			fmt.Println(err)
			errors.Error(writer, 500)
			return 1, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
		} else {
			yearint, _ := strconv.Atoi(yearfirst)
			yearint1, _ := strconv.Atoi(yearlast)
			if yearint1 < yearint {
				yearlast = yearfirst
			}
		}
	}

	fa := r.FormValue("FirstAlbum")
	if fa != "" {
		if faint, err := strconv.Atoi(fa); err != nil || (faint < 1950 || faint > 2010) {
			fmt.Println(err)
			errors.Error(writer, 500)
			return 1, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
		}
	}

	fa2 := r.FormValue("FirstAlbum2")
	if fa2 != "" {
		if faint, err := strconv.Atoi(fa2); err != nil || (faint < 1955 || faint > 2023) {
			fmt.Println(err)
			errors.Error(writer, 500)
			return 1, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
		} else {
			faint, _ := strconv.Atoi(fa)
			faint1, _ := strconv.Atoi(fa2)
			if faint1 < faint {
				fa2 = fa
			}
		}
	}

	nm := r.FormValue("numberMumbers")
	if nm != "" {
		if nmint, err := strconv.Atoi(nm); err != nil || (nmint < 1 || nmint > 8) {
			fmt.Println(err)
			errors.Error(writer, 500)
			return 1, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
		}
	}

	nm2 := r.FormValue("numberMumbers2")
	if nm2 != "" {
		if nmint, err := strconv.Atoi(nm2); err != nil || (nmint < 1 || nmint > 8) {
			fmt.Println(err)
			errors.Error(writer, 500)
			return 1, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
		} else {
			nmint, _ := strconv.Atoi(nm)
			nm2int, _ := strconv.Atoi(nm2)
			if nm2int < nmint {
				nm2 = nm
			}
		}
	}

	cityState := r.FormValue("ville")
	if cityState != "" {
		cityState = strings.ToLower(cityState)
		cityState = strings.ReplaceAll(cityState, ", ", "-")
		cityState = strings.ReplaceAll(cityState, " ", "_")
	}

	searchBar := r.FormValue("searchbar")

	return 0, na, sa, eu, me, as, oc, doc, doc1, doc2, yearfirst, yearlast, fa, fa2, nm, nm2, cityState, searchBar
}
