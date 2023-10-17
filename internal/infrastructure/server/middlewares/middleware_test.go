package middlewares

import (
	"net/http"
	"reflect"
	"testing"
)

func TestRecovery(t *testing.T) {
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Recovery(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Recovery() = %v, want %v", got, tt.want)
			}
		})
	}
}
