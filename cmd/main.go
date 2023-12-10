package main

import (
	"net/http"

	"github.com/gabriel-valin/ms_upload_files/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload-files", handlers.UploadHandler)

	http.ListenAndServe(":1337", mux)
}
