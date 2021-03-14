package service

import (
	"io"
	"net/http"
)

func HandleLibrary(w http.ResponseWriter, req *http.Request) {
	header := w.Header()
	header.Set("Content-Type", "application/json")

	io.WriteString(w, "library\n")
}
