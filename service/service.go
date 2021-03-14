package service

import (
	"log"
	"net/http"
)

func Run() {

	http.HandleFunc("/", HandleHealth)
	http.HandleFunc("/health", HandleHealth)

	http.HandleFunc("/package", HandlePackage)
	http.HandleFunc("/library", HandleLibrary)

	// TODO: MAKE THIS A POST
	http.HandleFunc("/build", HandleBuild)

	log.Println("Listing for requests at http://localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
