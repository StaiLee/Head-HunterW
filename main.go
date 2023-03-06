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

const port = 1706

var Templatesartists = template.Must(template.ParseFiles("./frontend/src/templates/artists.html"))

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	homeHandler()
}

func homeHandler() {
	go http.ListenAndServe(":1706", nil)
	http.Handle("/", http.FileServer(http.Dir("./frontend/src")))
	http.HandleFunc("/artists", utils.ArtistsHandler)
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
