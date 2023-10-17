package mjsonapi

// ErrorJSONAPI struct base from error response
type ErrorJSONAPI struct {
	Id     string                 `json:"id"`
	Links  map[string]interface{} `json:"links"`
	Status string                 `json:"status"`
	Code   string                 `json:"code"`
	Title  string                 `json:"title"`
	Detail string                 `json:"detail"`
	Source map[string]interface{} `json:"source"`
	Meta   map[string]interface{} `json:"meta"`
}

// ErrorsJSONAPIProvider interface to add or get errors
type ErrorsJSONAPIProvider interface {
	Add(error ErrorJSONAPI) *ErrorsJSONAPI
	Get() *ErrorsJSONAPI
}

// ErrorsJSONAPI list of errors
type ErrorsJSONAPI struct {
	Errors []ErrorJSONAPI `json:"errors"`
}

// Add an error
func (e *ErrorsJSONAPI) Add(error ErrorJSONAPI) *ErrorsJSONAPI {
	e.Errors = append(e.Errors, error)
	return e
}

// Get an error
func (e *ErrorsJSONAPI) Get() *ErrorsJSONAPI {
	return e
}

// NewErrorsJSONAPI init error response
func NewErrorsJSONAPI() *ErrorsJSONAPI {
	return &ErrorsJSONAPI{
		Errors: []ErrorJSONAPI{},
	}
}
