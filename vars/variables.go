package vars

import "html/template"

var (
	Templates     *template.Template
	Artists_url   = "https://groupietrackers.herokuapp.com/api/artists"
	Template_dir = "web/templates/"
)