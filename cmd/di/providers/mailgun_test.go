package providers

import (
	"reflect"
	"testing"
)

func TestNewMailgunProvider(t *testing.T) {
	tests := []struct {
		name string
		want *MailgunProvider
	}{
		{
			name: "mailgun provider ",
			want: NewMailgunProvider(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMailgunProvider(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMailgunProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}
