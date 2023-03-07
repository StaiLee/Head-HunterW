package main

import (
	utils "changeme/utils"
	"embed"
	"fmt"
	"net/http"
	"text/template"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//port pour notre serveur
const port = 1706

//notre premiere template artistes dans la quelle tout sera afficher
var Templatesartists = template.Must(template.ParseFiles("./frontend/src/templates/artists.html"))

//go:embed all:frontend/dist
var assets embed.FS

//notre main qui execute la fonction qui fait tout fonctionner
func main() {
	homeHandler()
}

func homeHandler() {
	//execution de notre serveur, une goroutine a ete utiliser car nous avons besoin d'un programme asynchrone pour executer l'app wails ainsi que le serveur web
	go http.ListenAndServe(":1706", nil)
	// on handle notre dossier qui contient notre js ainsi que notre css template ect
	http.Handle("/", http.FileServer(http.Dir("./frontend/src")))
	//nos deux templates artists qui contient les donnees de l'api ainsi que start ou il ya la barre de recherche
	http.HandleFunc("/artists", utils.ArtistsHandler)
	http.HandleFunc("/start", utils.StartHandler)
	//print pour la console pour que se soit plus accessible
	fmt.Println("http://localhost:1706/artists - server started on port", port)
	// Create an instance of the app structure
	app := NewApp()
	// Create application with options
	err := wails.Run(&options.App{
		Title:      "Head-HunterW",
		Width:      1600,
		Height:     900,
		Fullscreen: false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
