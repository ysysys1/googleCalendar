package function

import (
	"net/http"

	"github.com/calendar-open/hello"
)

func Hello(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		hello.HelloWorld(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
