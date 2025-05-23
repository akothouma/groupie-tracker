package errorhandling

import (
	"net"
	"net/http"
	"os"
	"time"

	"groupie/vars"
)

// handles HTTP errors by writing the specified status code and error message
// to the response writer, and then rendering an error template with the provided data.
func RenderErr(w http.ResponseWriter, statusCode int, headerMessage string, errMessage string) {
	if _, err := os.Stat(vars.Template_dir + "errors.html"); os.IsNotExist(err) {
		http.Error(w, "Currently unable to display this page. Please try again later.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)

	// data to pass to the error template
	data := struct {
		StatusCode int
		HeaderMessage     string
		Message    string
	}{
		StatusCode: statusCode,
		HeaderMessage:     headerMessage,
		Message:    errMessage,
	}

	// render the page
	vars.Templates.ExecuteTemplate(w, "errors.html", &data)
}

func ClientConnected() bool {
	timeout := 7 * time.Second
	_, err := net.DialTimeout("tcp", "8.8.8.8:53", timeout)
	return err == nil
}
