package errorhandling

import (
	"net/http"

	"groupie/vars"
)

// handles HTTP errors by writing the specified status code and error message
// to the response writer, and then rendering an error template with the provided data.
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
	errorsTmplErr := vars.Templates.ExecuteTemplate(w, "errors.html", &data)
	if errorsTmplErr != nil {
		http.Error(w, "Currently unable to display this page. Please try again later.", http.StatusInternalServerError)
	}
}
