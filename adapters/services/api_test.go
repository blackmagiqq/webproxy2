package services

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServiceGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET method, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer ts.Close()

	service := &APIService{}
	headers := map[string]string{"Authorization": "Bearer token"}

	resp, err := service.Handle(ts.URL, "GET", headers, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	expectedBody := `{"message": "success"}`
	if string(resp.Body) != expectedBody {
		t.Errorf("expected body %s, got %s", expectedBody, string(resp.Body))
	}
}

func TestAPIServicePost(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST method, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer ts.Close()

	service := &APIService{}
	headers := map[string]string{"Authorization": "Bearer token"}

	resp, err := service.Handle(ts.URL, "POST", headers, struct{ Test string }{Test: "test"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	expectedBody := `{"message": "success"}`
	if string(resp.Body) != expectedBody {
		t.Errorf("expected body %s, got %s", expectedBody, string(resp.Body))
	}
}
