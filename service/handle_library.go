package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"
	"github.com/mikenomitch/bindle-packager/utils"
)

type nomadPackage struct {
	Name string `hcl:"key,label"`
}

type nomadCatalog struct {
	Packages []nomadPackage `hcl:"package,block"`
}

const catalogURL = "https://raw.githubusercontent.com/mikenomitch/nomad-packages/main/catalog.hcl"

func HandleCatalog(w http.ResponseWriter, req *http.Request) {
	header := w.Header()
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Content-Type", "application/json")

	packages := getCatalog()

	json.NewEncoder(w).Encode(packages)
}

func getCatalog() []nomadPackage {
	bodyBuffer, _ := utils.BufferFromURL(catalogURL)

	parser := hclparse.NewParser()
	file, diags := parser.ParseHCL(bodyBuffer, "catalog")

	res := nomadCatalog{}
	if diags = gohcl.DecodeBody(file.Body, nil, &res); diags.HasErrors() {
		log.Printf(diags.Error())
	}

	return res.Packages
}
