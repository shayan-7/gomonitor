package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMemberRegister(t *testing.T) {
	r := strings.NewReader("Hello, Reader!")

	req, err := http.NewRequest("GET", "localhost:8080/apiv1/members", r)
	if err != nil {
		t.Fatalf("Couldn't create request")
	}

	rec := httptest.NewRecorder()
	registerMember(rec, req)
	resp := rec.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected '200 OK', got: %+v", resp.StatusCode)
	}
}
