package utils

import (
	"testing"
)

func TestMarshalForm(t *testing.T) {
	data := map[string]string{
		"foo": "bar",
	}
	res := MarshalForm(data)
	if res == "" {
		t.Error("marshal form failed")
	}

	rev := UnmarshalForm[map[string]string](res)
	if rev == nil || (*rev)["foo"] != "bar" {
		t.Error("form is invalid")
	}
}
