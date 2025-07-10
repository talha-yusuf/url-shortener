package utils

import (
	"fmt"
	"net/http"
)

// SendHTMLResponse sends an HTML response with the given template and arguments
func SendHTMLResponse(w http.ResponseWriter, template string, args ...interface{}) {
	w.Header().Set("Content-Type", "text/html")
	if len(args) > 0 {
		fmt.Fprintf(w, template, args...)
	} else {
		fmt.Fprint(w, template)
	}
}
