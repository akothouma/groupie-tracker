package errorhandling
import (
	"html/template"
	"net/http"
)

func RenderErr(w http.ResponseWriter, statusCode int, errMessage string) {
	if w.WriteHeader==nil {
		w.WriteHeader(statusCode)
        // Avoid writing headers again
    }

	// load the err template
	tmpl, err := template.ParseFiles("templates/errors.html")
	if err != nil {
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return
	}
	// data to pass to the error template

	data := struct {
		StatusCode int
		Message    string
	}{
		StatusCode: statusCode,
		Message:    errMessage,
	}

	// render the page
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}