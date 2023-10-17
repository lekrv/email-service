package emails

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"email-service/internal/domain"
	"email-service/pkg/mtools/mjsonapi"
	"email-service/pkg/mtools/mlogger"
	"email-service/pkg/mtools/mtypes"

	"github.com/xeipuuv/gojsonschema"
)

//go:embed schemas/validator.json
var validationSchemaJSON string
var validationSchema = gojsonschema.NewStringLoader(validationSchemaJSON)

// EmailAction is an interface that defines the behavior of sending an email.
type EmailAction interface {
	SendEmail(email domain.Email) error
}

// emailRequest represents the request body for sending an email.
type emailRequest struct {
	Email domain.Email `json:"data"`
}

// Handler defines event controller.
type Handler struct {
	emailAction EmailAction
	errors      mjsonapi.ErrorsJSONAPIProvider
	response    mjsonapi.ResponseJSONAPIProvider
	logger      mlogger.Logger
}

// defines response structure
type response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Date    string `json:"date"`
}

// Handle handles the HTTP request for sending an email.
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	//init logger
	logger := h.logger.WithFields(
		"file", "cx_handler",
		"method", "handler",
		"headers", r.Header,
		"body", r.Body,
	)

	// set response headers with the allowed origin.
	//update content type
	w.Header().Set(mtypes.ContentType, mtypes.ApplicationJSONContentType)

	var request emailRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		logger.WithError(err).Errorf("Error decoding request body.")
		h.errors.Add(mjsonapi.ErrorJSONAPI{
			Status: strconv.Itoa(http.StatusBadRequest),
			Code:   "FORMAT_ERROR",
			Title:  "Error",
			Detail: "Error decoding request body.",
		})

		http.Error(w, "Error decoding request body.", http.StatusBadRequest)
		return
	}

	paramsValidate := gojsonschema.NewGoLoader(&request)
	result, err := gojsonschema.Validate(validationSchema, paramsValidate)
	if err != nil {
		fmt.Println(err)
		logger.WithError(err).Errorf("Error validate schema with the fields")
		h.errors.Add(mjsonapi.ErrorJSONAPI{
			Status: strconv.Itoa(http.StatusBadRequest),
			Code:   "FORMAT_ERROR",
			Title:  "Error",
			Detail: "Error validate schema with the fielValidateds",
		})

		http.Error(w, "Error validate schema with the fields", http.StatusBadRequest)

		return
	}

	if !result.Valid() {
		fmt.Println(result.Errors())
		logger.WithError(err).Errorf("Error the fields are not valid.")
		h.errors.Add(mjsonapi.ErrorJSONAPI{
			Status: strconv.Itoa(http.StatusBadRequest),
			Code:   "FORMAT_ERROR",
			Title:  "Error",
			Detail: "Error the fields are not valid.",
		})

		http.Error(w, "Error the fields are not valid.", http.StatusBadRequest)

		return
	}

	err = h.emailAction.SendEmail(request.Email)
	if err != nil {
		logger.WithError(err).Errorf("Error to send email")
		h.errors.Add(mjsonapi.ErrorJSONAPI{
			Status: strconv.Itoa(http.StatusBadRequest),
			Code:   "FORMAT_ERROR",
			Title:  "Error",
			Detail: "Error to send email",
		})

		http.Error(w, "Error to send email", http.StatusBadRequest)

		return
	}

	// response success
	responseBody := transformerResponse(http.StatusOK, "Successfully sent email.", time.UnixDate)
	response, _ := json.Marshal(responseBody)

	_, err = w.Write(response)
	if err != nil {
		return
	}
}

// transformerResponse function to return response structure
func transformerResponse(status int, message string, date string) *response {
	return &response{
		Status:  status,
		Message: message,
		Date:    date,
	}
}

// NewHandler create a new instances of Handler
func NewHandler(
	emailAction EmailAction,
	errors mjsonapi.ErrorsJSONAPIProvider,
	response mjsonapi.ResponseJSONAPIProvider,
	logger mlogger.Logger,
) *Handler {
	return &Handler{
		emailAction: emailAction,
		errors:      errors,
		response:    response,
		logger:      logger,
	}
}
