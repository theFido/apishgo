package apish

// APISpec Expanded API definition
type APISpec struct {
	specItems  map[string]SpecEndpoint
	sortedAPIs []string
}

// SpecParam reusable parameter
type SpecParam struct {
	Name         string `json:"name"`
	DataType     string `json:"data_type"`
	Required     bool   `json:"required"`
	DefaultValue string `json:"default_value"`
	Description  string `json:"description"`
}

// SpecEndpoint full API specification
type SpecEndpoint struct {
	Get    SpecDefinition `json:"get"`
	Post   SpecDefinition `json:"post"`
	Put    SpecDefinition `json:"put"`
	Delete SpecDefinition `json:"delete"`
	Patch  SpecDefinition `json:"patch"`
}

// SpecDefinition full API definition
type SpecDefinition struct {
	Description  string                 `json:"description"`
	Operation    string                 `json:"operation"`
	UseCases     []string               `json:"use_cases"`
	Headers      []SpecParam            `json:"headers"`
	QueryStrings []SpecParam            `json:"query_strings"`
	PathParams   []SpecParam            `json:"params"`
	StatusCodes  []StatusCode           `json:"status_codes"`
	Produces     []string               `json:"produces"`
	Consumes     []string               `json:"consumes"`
	Example      map[string]interface{} `json:"example"`
}

// GetAPIs returns the sorted API endpoints as a list
func (s *APISpec) GetAPIs() []string {
	if s.sortedAPIs != nil {
		return s.sortedAPIs
	}
	keys := make([]string, len(s.specItems))

	ndx := 0
	for k, _ := range s.specItems {
		keys[ndx] = k
	}

	s.sortedAPIs = keys

	return keys
}
