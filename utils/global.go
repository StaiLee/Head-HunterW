package changeme

import "text/template"

//nos differentes templates html
var TemplateArtists = template.Must(template.ParseFiles("./frontend/src/templates/artists.html"))
var TemplateStart = template.Must(template.ParseFiles("./frontend/src/templates/start.html"))
