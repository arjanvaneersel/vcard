package vcard_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/arjanvaneersel/vcard"
)

func TestN(t *testing.T) {
	f := vcard.N{
		FamilyName:        "Person",
		GivenName:         "Test",
		HonorificSuffixes: "PhD",
	}

	if expected, got := "N:Person;Test;;;PhD", f.String(); expected != got {
		t.Fatalf("expected %q, but got %q", expected, got)
	}
}

func TestFN(t *testing.T) {
	f := vcard.FN{"Test Person"}

	if expected, got := "FN:Test Person", f.String(); expected != got {
		t.Fatalf("expected %q, but got %q", expected, got)
	}
}

func TestORG(t *testing.T) {
	tt := []struct {
		field    vcard.Org
		expected string
	}{
		{vcard.Org{Name: "UAB Test Company"}, "ORG:UAB Test Company"},
		{vcard.Org{Name: "UAB Test Company", Units: []string{"Software development", "Programming"}}, "ORG:UAB Test Company;Software development;Programming"},
	}

	for _, tc := range tt {
		if got := tc.field.String(); got != tc.expected {
			t.Fatalf("expected %q, but got %q", tc.expected, got)
		}
	}
}

func TestTitle(t *testing.T) {
	f := vcard.Title{"V.P. Research and Development"}

	if expected, got := "TITLE:V.P. Research and Development", f.String(); expected != got {
		t.Fatalf("expected %q, but got %q", expected, got)
	}
}

func TestRole(t *testing.T) {
	f := vcard.Role{"Programmer"}

	if expected, got := "ROLE:Programmer", f.String(); expected != got {
		t.Fatalf("expected %q, but got %q", expected, got)
	}
}

func TestPhoto(t *testing.T) {
	url, _ := url.Parse("http://www.abc.com/pub/photos/jqpublic.gif")
	tt := []struct {
		field    vcard.Photo
		expected string
	}{
		{
			vcard.Photo{
				Binary: true,
				Type:   "JPEG",
				Data:   []byte("MIICajCCAdOgAwIBAgICBEUwDQYJKoZIhvcNAQEEBQAwdzELMAkGA1UEBhMCVVMxLDAqBgNVBAoTI05ldHNjYXBlIENvbW11bmljYXRpb25zIENvcnBvcmF0aW9uMRwwGgYDVQQLExNJbmZvcm1hdGlvbiBTeXN0"),
			},
			"PHOTO;ENCODING=b;TYPE=JPEG:MIICajCCAdOgAwIBAgICBEUwDQYJKoZIhvcNAQEEBQAwdzELMAkGA1UEBhMCVVMxLDAqBgNVBAoTI05ldHNjYXBlIENvbW11bmljYXRpb25zIENvcnBvcmF0aW9uMRwwGgYDVQQLExNJbmZvcm1hdGlvbiBTeXN0",
		},
		{
			vcard.Photo{
				URL: url,
			},
			"PHOTO;VALUE=uri:http://www.abc.com/pub/photos/jqpublic.gif",
		},
		{
			vcard.Photo{
				Type: "JPEG",
				URL:  url,
			},
			"PHOTO;TYPE=JPEG;VALUE=uri:http://www.abc.com/pub/photos/jqpublic.gif",
		},
	}

	for _, tc := range tt {
		if got := tc.field.String(); got != tc.expected {
			t.Fatalf("expected %q, but got %q", tc.expected, got)
		}
	}
}

func TestTel(t *testing.T) {
	tt := []struct {
		field    vcard.Tel
		expected string
	}{
		{vcard.Tel{Number: "+3701234567890"}, "TEL;TYPE=voice:+3701234567890"},
		{vcard.Tel{Number: "+3701234567890", Types: []string{vcard.TelVoice, vcard.TelFax}}, "TEL;TYPE=voice,fax:+3701234567890"},
	}

	for _, tc := range tt {
		if got := tc.field.String(); got != tc.expected {
			t.Fatalf("expected %q, but got %q", tc.expected, got)
		}
	}
}

func TestAdr(t *testing.T) {
	tt := []struct {
		field    vcard.Adr
		expected string
	}{
		{
			vcard.Adr{
				StreetAddress: "Bandymus gatve 1-1",
				PostalCode:    "10000",
				Locality:      "Testius",
				CountryName:   "Testland",
			},
			"ADR;TYPE=intl,postal,parcel,work:;;Bandymus gatve 1-1;Testius;;10000;Testland",
		},
		{
			vcard.Adr{
				Types:         []string{vcard.AdrWork},
				StreetAddress: "Bandymus gatve 1-1",
				PostalCode:    "10000",
				Locality:      "Testius",
				CountryName:   "Testland",
			},
			"ADR;TYPE=work:;;Bandymus gatve 1-1;Testius;;10000;Testland",
		},
	}

	for _, tc := range tt {
		if got := tc.field.String(); got != tc.expected {
			t.Fatalf("expected %q, but got %q", tc.expected, got)
		}
	}
}

func TestEmail(t *testing.T) {
	tt := []struct {
		field    vcard.Email
		expected string
	}{
		{vcard.Email{Email: "test@example.com"}, "EMAIL:test@example.com"},
		{vcard.Email{Types: []string{vcard.EmailInternet}, Email: "test@example.com"}, "EMAIL;TYPE=internet:test@example.com"},
	}

	for _, tc := range tt {
		if got := tc.field.String(); got != tc.expected {
			t.Fatalf("expected %q, but got %q", tc.expected, got)
		}
	}
}

func TestRev(t *testing.T) {
	now := time.Now()
	tt := []struct {
		Timestamp time.Time
		Format    string
		Expected  string
	}{
		{
			Timestamp: now,
			Expected:  fmt.Sprintf("REV:%s", now.Format("20060102T150405Z0700")),
		},
		{
			Timestamp: now,
			Format:    time.RFC3339,
			Expected:  fmt.Sprintf("REV:%s", now.Format(time.RFC3339)),
		},
	}

	for _, tc := range tt {
		f := vcard.Rev{Timestamp: tc.Timestamp, Format: tc.Format}
		if got := f.String(); got != tc.Expected {
			t.Fatalf("expected %q, but got %q", tc.Expected, got)
		}
	}
}
