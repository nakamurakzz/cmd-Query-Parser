package main

import (
	"testing"
)

func Test_getQueryParams(t *testing.T) {
	want := map[string]string{
		"foo":  "bar",
		"hoge": "fuga",
	}
	rawURL := "https://example.com?foo=bar&hoge=fuga"

	queries := getQueryParams(rawURL)
	for key, value := range queries {
		if value != want[key] {
			t.Errorf("getQueryParams() = %v, want %v", value, want[key])
		}
	}
}
