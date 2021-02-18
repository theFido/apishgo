package apish

import (
	"encoding/json"
	"io/ioutil"
)

// NewFromFile Reads an API raw definition file
func NewAPIFromFile(filename string) (*API, error) {
	var api API
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, &api)
	return &api, err
}

// NewAPISpecFromFile Reads an API definition file
func NewAPISpecFromFile(filename string) (*APISpec, error) {
	var spec map[string]SpecEndpoint
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, &spec)
	return &APISpec{spec, nil}, err
}
