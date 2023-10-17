package mtypes

import (
	"testing"
)

// TestSliceItemExists Contain everything type validation.
func TestSliceItemExists(t *testing.T) {
	type args struct {
		slice interface{}
		item  interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Exist number in slice",
			args: args{
				slice: []int{1, 2, 3},
				item:  1,
			},
			want: true,
		},
		{
			name: "Not exist number in slice",
			args: args{
				slice: []int{1, 2, 3},
				item:  4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceItemExists(tt.args.slice, tt.args.item); got != tt.want {
				t.Errorf("SliceItemExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFromEmail(t *testing.T) {
	type args struct {
		countryCode string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test success colombia",
			args: args{
				countryCode: "co",
			},
			want: "soportecolombia@merqueo.com",
		},
		{
			name: "test success brasil",
			args: args{
				countryCode: "br",
			},
			want: "soportebrasil@merqueo.com",
		},
		{
			name: "test success mexico",
			args: args{
				countryCode: "mx",
			},
			want: "soportemx@merqueo.com",
		},
		{
			name: "test success other country",
			args: args{
				countryCode: "ch",
			},
			want: "servicioalcliente@merqueo.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFromEmail(tt.args.countryCode); got != tt.want {
				t.Errorf("GetFromEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
