package changeme

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// notre fonction qui nous permet de recuperer les donnees de l'api, cette derniere contient egalement notre systeme de barre de recherche pour pourvoir afficher les donnees
//que l'on veut a l'interieur de notre artist.html
func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	var apiResponse APIResponse
	GetFromApi("https://groupietrackers.herokuapp.com/api/artists", &apiResponse.Artists)
	name := r.URL.Query().Get("name")
	if name == "" || name == " " {
		apiResponse.GetParam = false
	} else {
		apiResponse.GetParam = true
		searchTab := strings.Split(name, "-")
		if len(searchTab) > 1 {
			switch strings.TrimSpace(searchTab[1]) {
			case "Creation Date":
				val, err := strconv.Atoi(strings.TrimSpace(searchTab[0]))
				if err != nil {
					apiResponse.GetParam = false
				}
				apiResponse.SearchType = strings.TrimSpace(searchTab[1])
				for range apiResponse.Artists {
					for i, v := range apiResponse.Artists {
						if v.CreationDate != val {
							if i < len(apiResponse.Artists) {
								apiResponse.Artists = RemoveArtist(apiResponse.Artists, i)
							}
						}
					}
				}
				break
			case "Group":
				for range apiResponse.Artists {
					for i, v := range apiResponse.Artists {
						if v.Name != strings.TrimSpace(searchTab[0]) {
							if i < len(apiResponse.Artists) {
								apiResponse.Artists = RemoveArtist(apiResponse.Artists, i)
							}
						}
					}
				}
				break
			case "Member":
				for range apiResponse.Artists {
					for i, group := range apiResponse.Artists {
						match := false
						for _, v := range group.Members {
							if v == strings.TrimSpace(searchTab[0]) {
								match = true
							}
						}
						if !match {
							if i < len(apiResponse.Artists) {
								apiResponse.Artists = RemoveArtist(apiResponse.Artists, i)
							}
						}
						match = false
					}
				}
				break
			default:
				fmt.Print("Non reconnu : ", searchTab[1])
				break
			}
		}
	}
	TemplateArtists.ExecuteTemplate(w, "artists.html", apiResponse)
}

//notre template qui contient la barre, meme principe que pour la template precedente
func StartHandler(w http.ResponseWriter, r *http.Request) {
	var apiResponse APIResponse
	GetFromApi("https://groupietrackers.herokuapp.com/api/artists", &apiResponse.Artists)
	TemplateStart.ExecuteTemplate(w, "start.html", apiResponse)
}
