package hsp

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"testing"
)

func TestNewClientBase64EncodesCredentials(t *testing.T) {
	email := "test@test.com"
	password := "password"

	c := NewClient(email, password)

	creds := fmt.Sprintf("%s:%s", email, password)
	sEncCreds := base64.StdEncoding.EncodeToString([]byte(creds))
	expected := fmt.Sprintf(sEncCreds)

	if expected != c.auth {
		t.Errorf("Expected auth to be %s but instead got \"%s\"!", expected, c.auth)
	}
}

func TestNewRequestSetsAcceptHeader(t *testing.T) {
	c := NewClient("test@test.com", "password")

	req, _ := c.newRequest("/path", struct{}{})

	accept := req.Header.Get("Accept")
	if accept != "application/json" {
		t.Errorf("Expected Accept header to be \"application/json\" but instead got \"%s\"!", accept)
	}
}

func TestNewRequestSetsContentTypeHeader(t *testing.T) {
	c := NewClient("test@test.com", "password")

	req, _ := c.newRequest("/path", struct{}{})

	contentType := req.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type header to be \"application/json\" but instead got \"%s\"!", contentType)
	}
}

func TestNewRequestSetsAuthorizationHeader(t *testing.T) {
	c := NewClient("test@test.com", "password")
	c.auth = "TestAuth"

	req, _ := c.newRequest("/path", struct{}{})

	expected := "Basic " + c.auth
	authorization := req.Header.Get("Authorization")
	if authorization != expected {
		t.Errorf("Expected Authorization header to be \"%s\" but instead got \"%s\"!", expected, authorization)
	}
}

func TestNewRequestSetsPOSTMethod(t *testing.T) {
	c := NewClient("test@test.com", "password")

	req, _ := c.newRequest("/path", struct{}{})

	method := req.Method
	if method != http.MethodPost {
		t.Errorf("Expected method to be \"%s\" but instead got \"%s\"!", http.MethodPost, method)
	}
}

func TestNewRequestSetsUrl(t *testing.T) {
	basePath := "https://test.co.uk"
	path := "/path"

	c := NewClient("test@test.com", "password", BaseURL(basePath))

	req, _ := c.newRequest(path, struct{}{})

	url := req.URL
	expected, _ := url.Parse(basePath + path)
	if *expected != *url {
		t.Errorf("Expected url to be \"%s\" but instead got \"%s\"!", expected, url)
	}
}
