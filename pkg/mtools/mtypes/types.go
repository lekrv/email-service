// Package mtypes contain the pkg types
package mtypes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"

	"github.com/aws/aws-lambda-go/events"
)

// constants related to errors
const (
	// CodeHTTPBusinessError Code 280 for http business error
	CodeHTTPBusinessError int = 280

	// CodeGeneralError this error code is for unexpected errors responses
	CodeGeneralError string = "CODE_GENERAL_ERROR"

	// IDGeneralError this error id is used for unexpected errors s
	IDGeneralError string = "ID_GENERAL_ERROR"

	// GeneralErrorTitle General title when throws message Error.
	GeneralErrorTitle string = "Error"
)

// list of header keys.
const (
	// ContentType key for content type of the response.
	ContentType string = "Content-Type"
	// AccessControlAllowOrigin key for allowed origins.
	AccessControlAllowOrigin string = "Access-Control-Allow-Origin"
	// Origin key for origin request header.
	Origin string = "origin"
	// Accept key for accept content.
	Accept string = "Accept"
)

// response header values list.
const (
	// ApplicationJSONContentType application/json content type.
	ApplicationJSONContentType string = "application/json"
	// AccessControlAllowOriginDev allowed origin for localhost
	AccessControlAllowOriginDev string = "http://localhost:5173/"
)

// AllowedOrigins list of allowed origins for all requests of API Gateway.
var AllowedOrigins = []string{
	AccessControlAllowOriginDev,
}

// -----------------------------------------------------------------------------------------------------------

// ResponseJSONAPI basic structure to represent JSON API Spec response
type ResponseJSONAPI struct {
	Data Data `json:"data"`
}

// Data contains the basic information to respond in JSON API Spec
type Data struct {
	Type       string      `json:"type"`
	Attributes interface{} `json:"attributes"`
}

// -----------------------------------------------------------------------------------------------------------

// Response struct to service response
type Response events.APIGatewayProxyResponse

// RequestLambda struct to service request
type RequestLambda events.APIGatewayProxyRequest

// GetInsensitiveHeaders get headers allowed.
func (r *RequestLambda) GetInsensitiveHeaders() http.Header {
	headers := http.Header{}
	for header, value := range r.Headers {
		headers.Add(header, value)
	}
	return headers
}

// GetResponseHeaders CORS generate valid headers to response
func (r *RequestLambda) GetResponseHeaders() map[string]string {
	return map[string]string{
		ContentType:              ApplicationJSONContentType,
		AccessControlAllowOrigin: ValidateAllowedOrigins(r.GetInsensitiveHeaders().Get(Origin)),
	}
}

// ValidateAllowedOrigins check if the given origin exist in the list of allowed origins and return this
// or return a default CORS according ENV.
func ValidateAllowedOrigins(origin string) string {
	if SliceItemExists(AllowedOrigins, origin) {
		return origin
	}
	return os.Getenv("ACCESS_CONTROL_ALLOW_ORIGIN")
}

// SliceItemExists We check if exist an item inside a slice.
func SliceItemExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("SliceExists() given a non-slice type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

// GetDataFromGoldenFile This method reads the golden file located in the path given and return the content (string)
func GetDataFromGoldenFile(filePath string) string {
	goldenFile, _ := os.Open(filePath)
	defer func(goldenFile *os.File) {
		err := goldenFile.Close()
		if err != nil {
			panic(fmt.Errorf("golden file (%s) not found - %w", filePath, err))
		}
	}(goldenFile)

	fileBytes, _ := ioutil.ReadAll(goldenFile)
	data := &bytes.Buffer{}
	_ = json.Compact(data, fileBytes)

	return data.String()
}
