package mjsonapi

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"bitbucket.com/merqueo/product-services/common/mtools/mtypes"
)

func TestResponseJSONAPI_ResponseErrorLambda(t *testing.T) {
	var headersCors = map[string]string{"Access-Control-Allow-Origin": "*"}
	type args struct {
		provider func() ErrorsJSONAPIProvider
		status   int
		header   map[string]string
	}
	tests := []struct {
		name string
		args args
		want func() mtypes.ResponseLambda
	}{
		{
			name: "success 1 case",
			args: args{
				provider: func() ErrorsJSONAPIProvider {
					e := NewErrorsJSONAPI()

					error1 := ErrorJSONAPI{
						Id:     "",
						Links:  nil,
						Status: "404",
						Code:   "OBJECT_NOT_FOUND",
						Title:  "Error",
						Detail: "Elemento no encontrado",
						Source: nil,
						Meta:   nil,
					}

					e.Add(error1)

					return e
				},
				status: 404,
				header: nil,
			},
			want: func() mtypes.ResponseLambda {

				errorsJson := `{"errors":[{"id":"","links":null,"status":"404","code":"OBJECT_NOT_FOUND","title":"Error","detail":"Elemento no encontrado","source":null,"meta":null}]}`

				var buf bytes.Buffer

				json.HTMLEscape(&buf, []byte(errorsJson))

				return mtypes.ResponseLambda{
					StatusCode: 404,
					Headers:    headersCors,
					Body:       buf.String(),
				}
			},
		},
		{
			name: "success 2 case",
			args: args{
				provider: func() ErrorsJSONAPIProvider {
					e := NewErrorsJSONAPI()

					error1 := ErrorJSONAPI{
						Id:     "",
						Links:  nil,
						Status: "404",
						Code:   "OBJECT_NOT_FOUND",
						Title:  "Error",
						Detail: "Elemento no encontrado",
						Source: nil,
						Meta:   nil,
					}

					e.Add(error1).Add(error1)

					return e
				},
				status: 404,
				header: nil,
			},
			want: func() mtypes.ResponseLambda {

				errorsJson := `{"errors":[{"id":"","links":null,"status":"404","code":"OBJECT_NOT_FOUND","title":"Error","detail":"Elemento no encontrado","source":null,"meta":null},{"id":"","links":null,"status":"404","code":"OBJECT_NOT_FOUND","title":"Error","detail":"Elemento no encontrado","source":null,"meta":null}]}`

				var buf bytes.Buffer

				json.HTMLEscape(&buf, []byte(errorsJson))

				return mtypes.ResponseLambda{
					StatusCode: 404,
					Headers:    headersCors,
					Body:       buf.String(),
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := &ResponseJSONAPI{}

			if got := response.ResponseErrorLambda(tt.args.provider(), tt.args.status, tt.args.header); !reflect.DeepEqual(got, tt.want()) {
				t.Errorf("ResponseErrorLambda() = %v, want %v", got, tt.want())
			}
		})
	}
}

func TestResponseJSONAPI_ResponseLambda(t *testing.T) {
	var headersCors = map[string]string{"Access-Control-Allow-Origin": "*"}
	type Response struct {
		ID      string `jsonapi:"primary,whats_app_data_send"`
		UuidPin string `jsonapi:"attr,uuid_pin"`
		To      string `jsonapi:"attr,to"`
	}

	type args struct {
		v      interface{}
		status int
		header map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    func() mtypes.ResponseLambda
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				v: &Response{
					ID:      "SEND_PIN",
					UuidPin: "d97919b8-db68-46d1-84f3-2e15b052f32e",
					To:      "573045652958",
				},
				status: 200,
				header: nil,
			},
			wantErr: false,
			want: func() mtypes.ResponseLambda {

				errorsJson := `{"data":{"type":"whats_app_data_send","id":"SEND_PIN","attributes":{"to":"573045652958","uuid_pin":"d97919b8-db68-46d1-84f3-2e15b052f32e"}}}`

				var buf bytes.Buffer

				json.HTMLEscape(&buf, []byte(errorsJson))

				return mtypes.ResponseLambda{
					StatusCode: 200,
					Headers:    headersCors,
					Body:       buf.String(),
				}

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := &ResponseJSONAPI{}
			got, err := response.ResponseLambda(tt.args.v, tt.args.status, tt.args.header)

			if (err != nil) != tt.wantErr {
				t.Error(err)
				t.Errorf("ResponseLambda() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want()) {
				t.Errorf("ResponseLambda() got = %v, want %v", got, tt.want())
			}
		})
	}
}
