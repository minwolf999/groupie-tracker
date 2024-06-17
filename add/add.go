package add

import (
	"slices"
	"strconv"
	"strings"

	data "test/data"
)

// Formating for searchBar
func Add(temp data.Art) []string {
	var megaTableauJVscript []string
	for i := range temp.Artists {
		megaTableauJVscript = append(megaTableauJVscript, temp.Artists[i].Name+"_|_Group_Name")
		for y := range temp.Artists[i].Members {
			megaTableauJVscript = append(megaTableauJVscript, strings.ReplaceAll(temp.Artists[i].Members[y], " ", "_")+"_|_Group_Member")
		}
		strTemp := temp.Artists[i].Locations
		strTemp = strings.ReplaceAll(strTemp, "_", "   ")
		strTemp = strings.Title(strTemp)
		strTemp = strings.ReplaceAll(strTemp, "   ", "_")
		strTemp = strings.ReplaceAll(strTemp, "<Br>", "")
		strTemp = strings.ReplaceAll(strTemp, " ", "")
		strTemp = strings.ReplaceAll(strTemp, "-", "_in_the_")
		tabtemp := strings.Split(strTemp[:len(strTemp)-2], ":*")

		for y := range tabtemp {
			megaTableauJVscript = append(megaTableauJVscript, tabtemp[y]+"_|_Group_Location")
		}
		megaTableauJVscript = append(megaTableauJVscript, strings.ReplaceAll(temp.Artists[i].FirstAlbum, "-", "/")+"_|_Group_First_Album")
		megaTableauJVscript = append(megaTableauJVscript, strconv.Itoa(temp.Artists[i].CreationDate)+"_|_Group_Creation_Date")

		for i := range megaTableauJVscript {
			megaTableauJVscript[i] = strings.ReplaceAll(megaTableauJVscript[i], "_", " ")
			megaTableauJVscript[i] = strings.ReplaceAll(megaTableauJVscript[i], "-", " ")
			megaTableauJVscript[i] = strings.ReplaceAll(megaTableauJVscript[i], "|", "-")
			megaTableauJVscript[i] = strings.ReplaceAll(megaTableauJVscript[i], " ", "!")
		}
	}
	megaTableauJVscript = Tri(megaTableauJVscript)

	megaTableauJVscript2 := strings.Title(strings.Join(megaTableauJVscript, " "))
	megaTableauJVscript2 = strings.ReplaceAll(megaTableauJVscript2, "!!!", "&nbsp;-&nbsp;")
	megaTableauJVscript2 = strings.ReplaceAll(megaTableauJVscript2, "!", "&nbsp;")
	megaTableauJVscript = strings.Split(megaTableauJVscript2, " ")
	return megaTableauJVscript
}

// Sort for searchBar
func Tri(tab []string) []string {
	// for i := 0; i < len(tab); i++ {
	// 	for y := len(tab) - 1; y > i; y-- {
	// 		if tab[i] == tab[y] {
	// 			tab = append(tab[:i], tab[i+1:]...)
	// 		}
	// 	}
	// }

	slices.Sort(tab)
	tab = slices.Compact(tab)
	return tab
}
