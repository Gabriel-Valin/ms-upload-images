package handlers

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "../public/index.html")
}
