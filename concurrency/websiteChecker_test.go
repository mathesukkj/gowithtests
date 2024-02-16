package main

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "wtf://notworkingbro" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"wtf://notworkingbro",
	}

	want := map[string]bool{
		"http://google.com":   true,
		"wtf://notworkingbro": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted %v got %v", want, got)
	}
}
