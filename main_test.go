package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", setupHandlers())
	return mux
}

func TestAddMessageSuccess(t *testing.T) {
	body := `{"body":"hello"}`
	req, err := http.NewRequest("POST", "/messages", bytes.NewBuffer([]byte(body)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := setupTestHandler()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	if rr.Body.String() != body {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), body)
	}
}

func TestAddMessageFailed(t *testing.T) {
	notBody := `"not_body":"not body"`
	req, err := http.NewRequest("POST", "/messages", bytes.NewBuffer([]byte(notBody)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := setupTestHandler()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestGetMessages(t *testing.T) {
	req, err := http.NewRequest("GET", "/messages", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := setupTestHandler()
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
