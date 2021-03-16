package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"
	"github.com/mikenomitch/bindle-packager/utils"
)

const baseURL = "https://raw.githubusercontent.com/mikenomitch/nomad-packages/main/"

type nomadResources struct {
	Resources []nomadResource `hcl:"nomad_resource,block"`
}

type nomadResource struct {
	Name         string `hcl:"key,label"`
	TemplateFile string `hcl:"template_file"`
	VariableFile string `hcl:"variable_file"`
	Description  string `hcl:"description"`
	Type         string `hcl:"type"`
}

type nomadVariables struct {
	Variables    []nomadVariable    `hcl:"variable,block"`
	VariableSets []nomadVariableSet `hcl:"variable_set,block"`
}

type nomadVariableMeta struct{}

type nomadVariable struct {
	Key         string            `hcl:"key,label"`
	Default     string            `hcl:"default"`
	Description string            `hcl:"description"`
	Type        string            `hcl:"type"`
	Meta        nomadVariableMeta `hcl:"meta,block"`
}

type nomadVariableSet struct {
	Key      string   `hcl:"key,label"`
	Contents []string `hcl:"contents"`
}

type packageResponse struct {
	Variables    []nomadVariable
	VariableSets []nomadVariableSet
	PackageInfo  []nomadResource
}

func HandlePackage(w http.ResponseWriter, req *http.Request) {
	header := w.Header()
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Content-Type", "application/json")

	query := req.URL.Query()
	name, present := query["name"]
	if !present || len(name) == 0 {
		fmt.Println("packageName not present")
	}
	packageName := name[0]

	resources := getResources(packageName)
	variables := getVariables(packageName)

	res := packageResponse{
		PackageInfo:  resources.Resources,
		Variables:    variables.Variables,
		VariableSets: variables.VariableSets,
	}

	// manifestBody
	json.NewEncoder(w).Encode(res)
}

func getResources(packageName string) nomadResources {
	manifestURL := baseURL + packageName + "/manifest.hcl"
	bodyBuffer, _ := utils.BufferFromURL(manifestURL)

	parser := hclparse.NewParser()
	file, diags := parser.ParseHCL(bodyBuffer, "ignoreme")

	res := nomadResources{}
	if diags = gohcl.DecodeBody(file.Body, nil, &res); diags.HasErrors() {
		log.Printf(diags.Error())
	}

	return res
}

func getVariables(packageName string) nomadVariables {
	varsURL := baseURL + packageName + "/variables.tf"
	bodyBuffer, _ := utils.BufferFromURL(varsURL)

	parser := hclparse.NewParser()
	file, diags := parser.ParseHCL(bodyBuffer, "ignoreme")

	vars := nomadVariables{}
	if diags = gohcl.DecodeBody(file.Body, nil, &vars); diags.HasErrors() {
		log.Printf(diags.Error())
	}

	return vars
}
