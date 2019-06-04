package vcard_test

import (
	"os"
	"strings"
	"testing"

	"github.com/arjanvaneersel/vcard"
)

func TestNew(t *testing.T) {
	v, err := vcard.New("4.0",
		vcard.N{FamilyName: "Person", GivenName: "Test"},
		vcard.FN{"Test Person"},
	)
	if err != nil {
		t.Fatalf("expected to pass, but got error %v", err)
	}

	if err := v.Validate(); err != nil {
		t.Fatalf("expected validation to pass, but got %v", err)
	}

	vcard := v.String()
	if len(vcard) == 0 {
		t.Fatalf("length should not be zero")
	}

	if !strings.Contains(vcard, "N:Person") {
		t.Fatalf("vcard should contain an N field for Test Persom")
	}
	t.Log(v)
}

func TestQRPng(t *testing.T) {
	filename := "qr.png"
	defer os.Remove(filename)
	v, err := vcard.New("4.0",
		vcard.N{FamilyName: "Person", GivenName: "Test"},
		vcard.FN{"Test Person"},
	)
	if err != nil {
		t.Fatalf("expected to pass, but got error %v", err)
	}

	if err := v.QRPng(500, 500, filename); err != nil {
		t.Fatalf("couldn't generate png: %v", err)
	}
}
