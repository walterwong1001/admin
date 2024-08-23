package crypto

import (
	"testing"
)

func TestEncode(t *testing.T) {
	encode, err := Encode("weitianonly")
	if err == nil {
		t.Log(encode)
		t.Log(len(encode))
	}
}
