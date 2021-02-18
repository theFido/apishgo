package main

import "apishgo/pkg/apish"

// APISpecPlugin produces artifacts from the raw API definition
type APIPlugin interface {
	CreateFromAPI(api apish.API) error
}

// APISpecPlugin produces artifacts from the full spec file
type APISpecPlugin interface {
	CreateFromSpec(spec apish.APISpec) error
}

// HybridPlugin component capable of producing an output using both definitions
type HybridPlugin interface {
	APIPlugin
	APISpecPlugin
}

func main() {

}
