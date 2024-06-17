package main

import (
	"fmt"
	"net/http"

	request "test/ApiAndDataRequest"
	"test/add"
	"test/server"
)

const PORT = ":8080"

func init() {
	artistsObject, locationsObject, datesObject := request.ApiRequest()
	server.ArtistsObject = artistsObject

	var tabLocations []string
	for i := range locationsObject {
		tabLocations = append(tabLocations, locationsObject[i].Locations...)
	}
	tabLocations = add.Tri(tabLocations)

	server.MapLocation = make(map[string]string)
	for i := range tabLocations {
		if i%4 == 0 {
			fmt.Printf("Loading ◜ %.2f%% \n", float64(i+1)*100.0/float64(len(tabLocations)))
		} else if i%4 == 1 {
			fmt.Printf("Loading ◝ %.2f%% \n", float64(i+1)*100.0/float64(len(tabLocations)))
		} else if i%4 == 2 {
			fmt.Printf("Loading ◞ %.2f%% \n", float64(i+1)*100.0/float64(len(tabLocations)))
		} else if i%4 == 3 {
			fmt.Printf("Loading ◟ %.2f%% \n", float64(i+1)*100.0/float64(len(tabLocations)))
		}

		chanRecever := make(chan string)
		go request.GeoCode(tabLocations[i], chanRecever)
		server.MapLocation[tabLocations[i]] = <-chanRecever

		fmt.Print("\033[1A\033[K")
	}

	server.ArtistsObject = request.AddDataLocation_ConcertDates(server.ArtistsObject, locationsObject, datesObject)
}

func main() {
	http.HandleFunc("/", server.Root)

	fmt.Printf("Starting server at port %v\nServing on http://localhost%v\n", PORT[1:], PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		panic(err)
	}
}
