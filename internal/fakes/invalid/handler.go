package invalid

import "net/http"

type Handler struct {
}

func (Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	code := http.StatusInternalServerError

	switch req.Method {
	case "POST":
		// Return the correct response code, with an invalid JSON response
		code = http.StatusCreated
	case "GET":
		// Return the correct response code, with an invalid JSON response
		code = http.StatusOK
	case "DELETE":
		// Return an invalid response code
		code = http.StatusTeapot
	}

	w.WriteHeader(code)
	w.Write([]byte("%%%"))
}
