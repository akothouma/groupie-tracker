package vars

import "html/template"

// global variables for templates and API endpoint
var (
	Templates    *template.Template
	Artists_url  = "https://groupietrackers.herokuapp.com/api/artists"
	Template_dir = "web/templates/"
)
