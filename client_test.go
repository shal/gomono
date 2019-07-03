package mono

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_GetJSON(t *testing.T) {
	cli := New("fake_token")

	// Mock fake server.
	expected := "This is fake response from the fake API"
	server := httptest.NewServer(
		http.HandlerFunc(
			func(rw http.ResponseWriter, req *http.Request) {
				rw.Write([]byte(expected))
			},
		),
	)
	defer server.Close()

	// Get request to the fake server to validate, that method works properly.
	BaseURL = server.URL
	body, status, err := cli.GetJSON("/fake")
	if err != nil || status != 200 || string(body) != expected {
		t.Fail()
	}
	BaseURL = DefaultBaseURL
}
