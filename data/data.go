package date

//  All struct

type Status struct {
	Code int
	Msg  string
}

type Locations []struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Dates []struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Art struct {
	MegaTableauJVscript struct {
		Truc []string
	}
	Artists []struct {
		ID           int      `json:"id"`
		Image        string   `json:"image"`
		Name         string   `json:"name"`
		Members      []string `json:"members"`
		CreationDate int      `json:"creationDate"`
		FirstAlbum   string   `json:"firstAlbum"`
		Locations    string   `json:"locations"`
		ConcertDates string   `json:"concertDates"`

		MembersString string
		TabLocations  []string
	}
}

type Coo struct {
	Features []Features `json:"features"`
}

type Features struct {
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
}

type Properties struct {
	Name     string   `json:"name"`
	State    string   `json:"state"`
	City     string   `json:"city"`
	Village  string   `json:"village"`
	Suburb   string   `json:"suburb"`
	Timezone Timezone `json:"timezone"`

	Country    string `json:"country"`
	Dependency string `json:"dependency"`
}

type Timezone struct {
	Name string `json:"name"`
}

type Geometry struct {
	Coordinates []float64 `json:"coordinates"`
}
