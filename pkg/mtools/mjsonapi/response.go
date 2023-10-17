package mjsonapi

import (
	"bytes"
	"email-service/pkg/mtools/mtypes"
	"encoding/json"

	"github.com/google/jsonapi"
)

// ResponseJSONAPIProvider provider to error response and JSON API response
type ResponseJSONAPIProvider interface {
	ResponseError(provider ErrorsJSONAPIProvider, status int, header map[string]string) mtypes.Response
	ResponseJsonApi(v interface{}, status int, header map[string]string) (mtypes.Response, error)
}

// ResponseJSONAPI struct to response lambda
type ResponseJSONAPI struct{}

// ResponseError function to response error lambda
func (response *ResponseJSONAPI) ResponseError(provider ErrorsJSONAPIProvider, status int, header map[string]string) mtypes.Response {
	var buf bytes.Buffer

	headers := map[string]string{"Access-Control-Allow-Origin": "*"}
	for k, v := range header {
		headers[k] = v
	}

	r, _ := json.Marshal(provider.Get())

	json.HTMLEscape(&buf, r)

	return mtypes.Response{
		StatusCode: status,
		Headers:    headers,
		Body:       buf.String(),
	}
}

// ResponseJsonApi function to response JSON API
func (response *ResponseJSONAPI) ResponseJsonApi(v interface{}, status int, header map[string]string) (mtypes.Response, error) {
	p, err := jsonapi.Marshal(v)

	if err != nil {
		return mtypes.Response{}, err
	}

	var buf bytes.Buffer

	r, _ := json.Marshal(p)

	json.HTMLEscape(&buf, r)

	headers := map[string]string{"Access-Control-Allow-Origin": "*"}
	for k, v := range header {
		headers[k] = v
	}

	return mtypes.Response{
		StatusCode: status,
		Headers:    headers,
		Body:       buf.String(),
	}, nil
}

// NewResponseJSONAPI init JSON API response
func NewResponseJSONAPI() ResponseJSONAPIProvider {
	return &ResponseJSONAPI{}
}
