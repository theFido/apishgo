package apish

// API Main API container
type API struct {
	Headers     []Param             `json:"headers"`
	QueryParams []Param             `json:"query"`
	PathParams  []Param             `json:"params"`
	StatusCodes []StatusCode        `json:"status_codes"`
	Endpoints   map[string]Endpoint `json:"endpoints"`
}

// Param reusable parameter
type Param struct {
	Name         string `json:"name"`
	DataType     string `json:"data_type"`
	Alias        string `json:"alias"`
	Required     bool   `json:"required"`
	DefaultValue string `json:"default_value"`
	Description  string `json:"description"`
}

// StatusCode custom definition
type StatusCode struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	IsRetryable bool   `json:"is_retryable"`
}

// Endpoint container
type Endpoint struct {
	Get    EndpointDefinition `json:"get"`
	Post   EndpointDefinition `json:"post"`
	Put    EndpointDefinition `json:"put"`
	Delete EndpointDefinition `json:"delete"`
	Patch  EndpointDefinition `json:"patch"`
}

// EndpointDefinition operation definition
type EndpointDefinition struct {
	Description string   `json:"description"`
	Operation   string   `json:"operation"`
	Headers     []string `json:"headers"`
	QueryParams []string `json:"query_string"`
	PathParams  []string `json:"path_params"`
	StatusCodes []string `json:"status_codes"`
	Produces    []string `json:"produces"`
	Consumes    []string `json:"consumes"`
	Example     string   `json:"example"`
}
