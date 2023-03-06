package changeme

import (
	"net/http"
)

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	var apiResponse APIResponse
	GetFromApi("https://groupietrackers.herokuapp.com/api/artists", &apiResponse.Artists)
	Templatesartists.ExecuteTemplate(w, "artists.html", apiResponse)
}
