package changeme

//nos differentes struct que nous allons manipuler pour obtenir le json dans notre page html

type Data struct {
	ArtistsURL   string `json:"artists"`
	LocationsURL string `json:"locations"`
	DatesURL     string `json:"dates"`
	RelationURL  string `json:"relation"`
}

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	LocationsURL string   `json:"locations"`
	DatesURL     string   `json:"concertDates"`
	RelationsURL string   `json:"relations"`
	Locations    []string
	Dates        []string
	Relations    map[string][]string
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	DatesURL  string   `json:"dates"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID        int                 `json:"id"`
	Relations map[string][]string `json:"datesLocations"`
}

type APIResponse struct {
	Artists    []Artist
	SearchType string
	Item       string
	GetParam   bool
}
