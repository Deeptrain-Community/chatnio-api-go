package utils

import "testing"

func TestGet(t *testing.T) {
	_, err := Get("https://example.com", Headers{})
	if err != nil {
		t.Error(err)
	}
}

func TestGetForm(t *testing.T) {
	TestGet(t)
}

func TestPost(t *testing.T) {
	_, err := Post("https://example.com", Headers{}, "")
	if err != nil {
		t.Error(err)
	}
}

func TestPostForm(t *testing.T) {
	TestPost(t)
}
