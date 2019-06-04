package vcard_test

import (
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

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
	url, _ := url.Parse("https://www.w3schools.com/w3css/img_avatar3.png")
	v, err := vcard.New("4.0",
		vcard.N{FamilyName: "Gump", GivenName: "Forrest", HonorificPrefixes: "Mr"},
		vcard.FN{"Forrest Gumo"},
		vcard.Org{Name: "Bubba Gump Shrimp Co."},
		vcard.Title{"Shrimp man"},
		vcard.Photo{URL: url},
		vcard.Tel{Number: "+1-111-555-1212", Types: []string{vcard.TelWork, vcard.TelVoice}},
		vcard.Email{Email: "forrest@example.com"},
		vcard.Rev{Timestamp: time.Now()},
	)
	if err != nil {
		t.Fatalf("expected to pass, but got error %v", err)
	}

	if err := v.QRPng(500, 500, filename); err != nil {
		t.Fatalf("couldn't generate png: %v", err)
	}
}
