package meta

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"

)

func TestGetAvailabilityZone(t *testing.T) {

	method := "GET"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			t.Errorf("r.Method = '%s', want '%s'", r.Method, method)
		}

		correctPath := "/latest/meta-data/placement/availability-zone"
		if r.URL.Path != correctPath {
			t.Errorf("r.URL.Path ='%v', want '%v'", r.URL.Path, correctPath)
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ap-northeast-1a"))
	}))

	defer ts.Close()

	client := ClientImpl{
		httpCli: &http.Client{},
		urlBase: ts.URL,
	}

	result, err := client.GetAvailabilityZone()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", result)

	assert.Equal(t, "ap-northeast-1a", result.Name)
	assert.Equal(t, "ap-northeast-1", result.GetRegion())
}

func TestGetInstanceID(t *testing.T) {

	method := "GET"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			t.Errorf("r.Method = '%s', want '%s'", r.Method, method)
		}

		correctPath := "/latest/meta-data/instance-id"
		if r.URL.Path != correctPath {
			t.Errorf("r.URL.Path ='%v', want '%v'", r.URL.Path, correctPath)
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("i-fffffff"))
	}))

	defer ts.Close()

	client := ClientImpl{
		httpCli: &http.Client{},
		urlBase: ts.URL,
	}

	result, err := client.GetInstanceID()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", result)

	assert.Equal(t, "i-fffffff", result)
}