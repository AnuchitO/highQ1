package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestBadVar(t *testing.T) {
	old := get
	defer func(o func(url string) (*http.Response, error)) {
		get = o
	}(old)

	get = func(url string) (*http.Response, error) {
		return &http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(`{"id": 1, "name": "AnuchitO", "description": "gopher"}`)),
		}, nil
	}

	resp, err := MakeHTTPCall("http://localhost:8080")
	if err != nil {
		t.Fatal(err)
	}

	want := &Response{
		ID:          1,
		Name:        "AnuchitO",
		Description: "gopher",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("expected (%v), got (%v)", want, resp)
	}
}
