package usecase

import (
	"email-service/internal/domain"
	"reflect"
	"testing"
)

func TestEmailUseCase_SendEmail(t *testing.T) {
	type fields struct {
		emailRepository domain.EmailRepository
	}
	type args struct {
		email domain.Email
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &EmailUseCase{
				emailRepository: tt.fields.emailRepository,
			}
			if err := uc.SendEmail(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("SendEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewEmailUseCase(t *testing.T) {
	type args struct {
		emailRepository domain.EmailRepository
	}
	tests := []struct {
		name string
		args args
		want *EmailUseCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmailUseCase(tt.args.emailRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmailUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
