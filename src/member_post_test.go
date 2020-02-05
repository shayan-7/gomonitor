package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMemberRegister(t *testing.T) {

	req, _ := http.NewRequest(
		"POST",
		"/apiv1/members",
		strings.NewReader(`"{"username": "shahan", "password": "password"}"`),
	)
	//if err != nil {
	//	t.Fatalf("Couldn't create request")
	//}

	rec := httptest.NewRecorder()
	registerMember(rec, req)
	resp := rec.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected '200 OK', got: %+v", resp.StatusCode)
	}
}
