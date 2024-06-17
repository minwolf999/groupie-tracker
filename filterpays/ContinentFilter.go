package filter

import data "test/data"

func ContinentFilter(artistsObject data.Art, na, sa, eu, me, as, oc string) data.Art {
	// Continents filters
	if na != "" {
		artistsObject = FilterPays("ListePays/North-America.txt", artistsObject)
	}
	if sa != "" {
		artistsObject = FilterPays("ListePays/South-America.txt", artistsObject)
	}
	if eu != "" {
		artistsObject = FilterPays("ListePays/Europe.txt", artistsObject)
	}
	if me != "" {
		artistsObject = FilterPays("ListePays/Moyen-Orient.txt", artistsObject)
	}
	if as != "" {
		artistsObject = FilterPays("ListePays/Asia.txt", artistsObject)
	}
	if oc != "" {
		artistsObject = FilterPays("ListePays/Oceania.txt", artistsObject)
	}

	return artistsObject
}
