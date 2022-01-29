//go:build integration

package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestStoreSuccess(t *testing.T) {
	p, _ := rand.Prime(rand.Reader, 10)
	key := "key" + p.String()
	val := "val" + p.String()
	url := fmt.Sprintf("http://localhost:%s/db/%s", os.Getenv("PORT"), key)

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(val))
	req.Header.Add("Content-Type", "application/octet-stream")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Error("Expected status code to be 201, got", resp.StatusCode)
	}
}
