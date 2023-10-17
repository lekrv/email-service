package emails

import (
	"email-service/pkg/mtools/mjsonapi"
	"email-service/pkg/mtools/mlogger"
	"net/http"
	"reflect"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	type fields struct {
		emailAction EmailAction
		errors      mjsonapi.ErrorsJSONAPIProvider
		response    mjsonapi.ResponseJSONAPIProvider
		logger      mlogger.Logger
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				emailAction: tt.fields.emailAction,
				errors:      tt.fields.errors,
				response:    tt.fields.response,
				logger:      tt.fields.logger,
			}
			h.Handle(tt.args.w, tt.args.r)
		})
	}
}

func TestNewHandler(t *testing.T) {
	type args struct {
		emailAction EmailAction
		errors      mjsonapi.ErrorsJSONAPIProvider
		response    mjsonapi.ResponseJSONAPIProvider
		logger      mlogger.Logger
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.emailAction, tt.args.errors, tt.args.response, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transformerResponse(t *testing.T) {
	type args struct {
		status  int
		message string
		date    string
	}
	tests := []struct {
		name string
		args args
		want *response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transformerResponse(tt.args.status, tt.args.message, tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transformerResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
