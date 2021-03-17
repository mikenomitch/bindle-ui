package service

import (
	"io"
	"net/http"
)

func HandleDeploy(w http.ResponseWriter, req *http.Request) {
	header := w.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	header.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	io.WriteString(w, "Hello, world!\n")
}
