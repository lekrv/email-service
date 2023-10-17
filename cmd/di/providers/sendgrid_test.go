package providers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"email-service/pkg/mtools/mconfig"
)

func TestSendGridProvider_SendEmail(t *testing.T) {
	// Mock SendGrid API key and environment variables
	env := mconfig.Env{
		ApiKeySendgrid:    "your-sendgrid-api-key",
		UserNameSendgrid:  "your-username",
		FromEmailSendgrid: "your-email@example.com",
	}

	mconfig.GlobalEnv = &env

	// Create a test HTTP server to mock SendGrid API responses
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && r.URL.Path == "/v3/mail/send" {
			// Simulate a successful SendGrid API response
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "Email sent successfully"}`))
		} else {
			// Simulate a bad request
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Bad request"}`))
		}
	}))
	defer server.Close()

	// Set the SendGrid API base URL to the test server's URL
	//sendgrid.BaseURL = server.URL

	// Create a SendGridProvider instance
	provider := NewSendGridProvider()

	// Test successful email sending
	toEmail := "recipient@example.com"
	emailSubject := "Test Subject"
	emailBody := "Test Body"

	err := provider.SendEmail(toEmail, emailSubject, emailBody)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Test email sending with a bad request
	env.ApiKeySendgrid = "invalid-api-key" // Change API key to trigger a bad request
	err = provider.SendEmail(toEmail, emailSubject, emailBody)
	if err != nil {
		t.Errorf("Expected an error message containing 'Bad request', but got: %v", err)
	}
}
