package pkg

import (
	"bytes"
	"log"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	key, err := GeneratePrivateKey()
	if err != nil {
		t.Error(err)
	}
	log.Println(key.String())
}

func TestParseKey(t *testing.T) {
	key, err := GeneratePrivateKey()
	if err != nil {
		t.Error(err)
	}
	log.Println(key.String())
	okey, err := ParseKey(key.String())
	if err != nil {
		t.Error(err)
	}
	log.Println(okey.String())
	if bytes.Compare(key[:], okey[:]) != 0 {
		t.Error("not equal")
	}
}
