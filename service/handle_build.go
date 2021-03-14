package service

import (
	"io"
	"net/http"
)

func HandleBuild(w http.ResponseWriter, req *http.Request) {
	header := w.Header()
	header.Set("Content-Type", "application/json")

	io.WriteString(w, "Hello, world!\n")
}
