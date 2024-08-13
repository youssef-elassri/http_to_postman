package parser

import (
	"http-to-postman/pkg/models"
	"reflect"
	"testing"
)

func TestParseHttpContent_GetRequest(t *testing.T) {
	content := "# Get User Info\nGET https://localhost:8080/api/v1/users/12345\nAuthorization: Bearer YOUR_ACCESS_TOKEN\n"
	want := []models.Request{
		{
			Description: "# Get User Info",
			Url:         "https://localhost:8080/api/v1/users/12345",
			Method:      "GET",
			Headers: []models.Header{{
				Key:   "Authorization",
				Value: "Bearer YOUR_ACCESS_TOKEN",
			}},
			Body: "",
		},
	}
	got, err := ParseHttpContent(content)
	if !reflect.DeepEqual(want, got) || err != nil {
		t.Fatalf(`ParseHttpContent(content) = %q, %v, want match for %#q, nil`, got, err, want)
	}
}

func TestParseHttpContent_PostRequest(t *testing.T) {
	content := "# Create User\nPOST https://localhost:8080/api/v1/users\nContent-Type: application/json\nAuthorization: Bearer YOUR_ACCESS_TOKEN\n\n{\"name\": \"John Doe\"}"
	want := []models.Request{
		{
			Description: "# Create User",
			Url:         "https://localhost:8080/api/v1/users",
			Method:      "POST",
			Headers: []models.Header{
				{Key: "Content-Type", Value: "application/json"},
				{Key: "Authorization", Value: "Bearer YOUR_ACCESS_TOKEN"},
			},
			Body: `{"name": "John Doe"}`,
		},
	}
	got, err := ParseHttpContent(content)

	if !reflect.DeepEqual(want, got) || err != nil {
		t.Fatalf(`ParseHttpContent(content) = %q, %v, want match for %#q, nil`, got, err, want)
	}
}

func TestParseHttpContent_EmptyContent(t *testing.T) {
	content := ""
	var want []models.Request
	got, err := ParseHttpContent(content)

	if !reflect.DeepEqual(want, got) || err != nil {
		t.Fatalf(`ParseHttpContent(content) = %q, %v, want match for %#q`, got, err, want)
	}
}
