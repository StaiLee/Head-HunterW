package changeme

import "text/template"

var Templatesartists = template.Must(template.ParseFiles("./frontend/src/templates/artists.html"))
