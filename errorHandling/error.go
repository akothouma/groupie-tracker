package errorhandling

import (
	"net/http"

	"groupie/vars"
)

func RenderErr(w http.ResponseWriter, statusCode int, errMessage string) {
	w.WriteHeader(statusCode)

	// data to pass to the error template

	data := struct {
		StatusCode int
		Message    string
	}{
		StatusCode: statusCode,
		Message:    errMessage,
	}

	// render the page
	vars.Templates.ExecuteTemplate(w, "errors.html", &data)
}
