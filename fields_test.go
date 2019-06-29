package vcard_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/arjanvaneersel/vcard"
)

func TestN(t *testing.T) {
	got, err := vcard.N{
		FamilyName:        "Person",
		GivenName:         "Test",
		HonorificSuffixes: "PhD",
	}.Format("4.0")
	if err != nil {
		t.Fatalf("expected to pass, but got: %v", err)
	}

	if expected := "N:Person;Test;;;PhD"; expected != got {
		t.Fatalf("expected %q, but got %q", expected, got)
	}
}

func TestFN(t *testing.T) {
	got, err := vcard.FN{"Test Person"}.Format("4.0")
	if err != nil {
		t.Fatalf("expected to pass, but got: %v", err)
	}

	if expected := "FN:Test Person"; expected != got {
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
		got, err := tc.field.Format("4.0")
		if err != nil {
			t.Fatalf("expected to pass, but got: %v", err)
		}
		if got != tc.expected {
			t.Fatalf("expected %q, but got %q", tc.expected, got)
		}
	}
}

func TestTitle(t *testing.T) {
	got, err := vcard.Title{"V.P. Research and Development"}.Format("4.0")
	if err != nil {
		t.Fatalf("expected to pass, but got: %v", err)
	}

	if expected := "TITLE:V.P. Research and Development"; expected != got {
		t.Fatalf("expected %q, but got %q", expected, got)
	}
}

func TestRole(t *testing.T) {
	got, err := vcard.Role{"Programmer"}.Format("4.0")
	if err != nil {
		t.Fatalf("expected to pass, but got: %v", err)
	}

	if expected := "ROLE:Programmer"; expected != got {
		t.Fatalf("expected %q, but got %q", expected, got)
	}
}

func TestPhoto(t *testing.T) {
	uri, _ := url.Parse("http://example.com/photo.jpg")
	tt := []struct {
		version  string
		field    vcard.Photo
		expected string
	}{
		{
			"2.1",
			vcard.Photo{
				Type:       "JPEG",
				Base64Data: "MIICajCCAdOgAwIBAgICBEUwDQYJKoZIhvcNAQEEBQAwdzELMAkGA1UEBhMCVVMxLDAqBgNVBAoTI05ldHNjYXBlIENvbW11bmljYXRpb25zIENvcnBvcmF0aW9uMRwwGgYDVQQLExNJbmZvcm1hdGlvbiBTeXN0",
			},
			"PHOTO;JPEG;ENCODING=BASE64:MIICajCCAdOgAwIBAgICBEUwDQYJKoZIhvcNAQEEBQAwdzELMAkGA1UEBhMCVVMxLDAqBgNVBAoTI05ldHNjYXBlIENvbW11bmljYXRpb25zIENvcnBvcmF0aW9uMRwwGgYDVQQLExNJbmZvcm1hdGlvbiBTeXN0",
		},
		{
			"2.1",
			vcard.Photo{
				URI: uri,
			},
			"PHOTO:http://example.com/photo.jpg",
		},
		{
			"2.1",
			vcard.Photo{
				Type: "JPEG",
				URI:  uri,
			},
			"PHOTO;JPEG:http://example.com/photo.jpg",
		},

		{
			"3.0",
			vcard.Photo{
				Type:       "JPEG",
				Base64Data: "MIICajCCAdOgAwIBAgICBEUwDQYJKoZIhvcNAQEEBQAwdzELMAkGA1UEBhMCVVMxLDAqBgNVBAoTI05ldHNjYXBlIENvbW11bmljYXRpb25zIENvcnBvcmF0aW9uMRwwGgYDVQQLExNJbmZvcm1hdGlvbiBTeXN0",
			},
			"PHOTO;TYPE=JPEG;ENCODING=b:MIICajCCAdOgAwIBAgICBEUwDQYJKoZIhvcNAQEEBQAwdzELMAkGA1UEBhMCVVMxLDAqBgNVBAoTI05ldHNjYXBlIENvbW11bmljYXRpb25zIENvcnBvcmF0aW9uMRwwGgYDVQQLExNJbmZvcm1hdGlvbiBTeXN0",
		},
		{
			"3.0",
			vcard.Photo{
				URI: uri,
			},
			"PHOTO;VALUE=uri:http://example.com/photo.jpg",
		},
		{
			"3.0",
			vcard.Photo{
				Type: "JPEG",
				URI:  uri,
			},
			"PHOTO;TYPE=JPEG;VALUE=uri:http://example.com/photo.jpg",
		},

		{
			"4.0",
			vcard.Photo{
				Type:       "image/jpeg",
				Base64Data: "MIICajCCAdOgAwIBAgICBEUwDQYJKoZIhvcNAQEEBQAwdzELMAkGA1UEBhMCVVMxLDAqBgNVBAoTI05ldHNjYXBlIENvbW11bmljYXRpb25zIENvcnBvcmF0aW9uMRwwGgYDVQQLExNJbmZvcm1hdGlvbiBTeXN0",
			},
			"PHOTO:data:image/jpeg;base64,MIICajCCAdOgAwIBAgICBEUwDQYJKoZIhvcNAQEEBQAwdzELMAkGA1UEBhMCVVMxLDAqBgNVBAoTI05ldHNjYXBlIENvbW11bmljYXRpb25zIENvcnBvcmF0aW9uMRwwGgYDVQQLExNJbmZvcm1hdGlvbiBTeXN0",
		},
		{
			"4.0",
			vcard.Photo{
				URI: uri,
			},
			"PHOTO:http://example.com/photo.jpg",
		},
		{
			"4.0",
			vcard.Photo{
				Type: "image/jpeg",
				URI:  uri,
			},
			"PHOTO;MEDIATYPE=image/jpeg:http://example.com/photo.jpg",
		},
	}

	for _, tc := range tt {
		got, err := tc.field.Format(tc.version)
		if err != nil {
			t.Fatalf("expected to pass, but got: %v", err)
		}

		if got != tc.expected {
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
		got, err := tc.field.Format("4.0")
		if err != nil {
			t.Fatalf("expected to pass, but got: %v", err)
		}

		if got != tc.expected {
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
		got, err := tc.field.Format("4.0")
		if err != nil {
			t.Fatalf("expected to pass, but got: %v", err)
		}

		if got != tc.expected {
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
		got, err := tc.field.Format("4.0")
		if err != nil {
			t.Fatalf("expected to pass, but got: %v", err)
		}

		if got != tc.expected {
			t.Fatalf("expected %q, but got %q", tc.expected, got)
		}
	}
}

func TestRev(t *testing.T) {
	now := time.Now()
	tt := []struct {
		timestamp time.Time
		format    string
		expected  string
	}{
		{
			timestamp: now,
			expected:  fmt.Sprintf("REV:%s", now.Format("20060102T150405Z0700")),
		},
		{
			timestamp: now,
			format:    time.RFC3339,
			expected:  fmt.Sprintf("REV:%s", now.Format(time.RFC3339)),
		},
	}

	for _, tc := range tt {
		got, err := vcard.Rev{Timestamp: tc.timestamp, TimeFormat: tc.format}.Format("4.0")
		if err != nil {
			t.Fatalf("expected to pass, but got: %v", err)
		}

		if got != tc.expected {
			t.Fatalf("expected %q, but got %q", tc.expected, got)
		}
	}
}

func TestAgent(t *testing.T) {
	card, err := vcard.New("3.0", vcard.N{FamilyName: "Person", GivenName: "Test"}, vcard.FN{"Test Person"})
	if err != nil {
		t.Fatalf("expected to pass, but got: %v", err)
	}

	cardTxt, err := card.Generate()
	if err != nil {
		t.Fatalf("expected to pass, but got: %v", err)
	}

	tt := []struct {
		version     string
		vcard       *vcard.VCard
		text        string
		expectedErr error
		expected    string
	}{
		{
			version:  "3.0",
			text:     "Test Person",
			expected: "AGENT:Test Person",
		},
		{
			version:  "3.0",
			vcard:    card,
			expected: fmt.Sprintf("AGENT:%s", cardTxt),
		},
		{
			version:     "4.0",
			vcard:       card,
			expectedErr: vcard.ErrVersion,
		},
	}

	for _, tc := range tt {
		got, err := vcard.Agent{VCard: tc.vcard, Text: tc.text}.Format(tc.version)
		if err != tc.expectedErr {
			t.Fatalf("expected err %v, but got: %v", tc.expectedErr, err)
		}

		if got != tc.expected {
			t.Fatalf("expected %q, but got %q", tc.expected, got)
		}
	}
}

func TestAniversary(t *testing.T) {
	now := time.Now()
	tt := []struct {
		version     string
		timestamp   time.Time
		format      string
		expected    string
		expectedErr error
	}{
		{
			version:   "4.0",
			timestamp: now,
			expected:  fmt.Sprintf("ANNIVERSARY:%s", now.Format("20060102")),
		},
		{
			version:   "4.0",
			timestamp: now,
			format:    time.RFC3339,
			expected:  fmt.Sprintf("ANNIVERSARY:%s", now.Format(time.RFC3339)),
		},
		{
			version:     "3.0",
			expectedErr: vcard.ErrVersion,
		},
		{
			version:     "2.1",
			expectedErr: vcard.ErrVersion,
		},
	}

	for _, tc := range tt {
		got, err := vcard.Anniversary{Date: tc.timestamp, TimeFormat: tc.format}.Format(tc.version)
		if err != tc.expectedErr {
			t.Fatalf("expected err %v, but got: %v", tc.expectedErr, err)
		}

		if got != tc.expected {
			t.Fatalf("expected %q, but got %q", tc.expected, got)
		}
	}
}

func TestBday(t *testing.T) {
	now := time.Now()
	tt := []struct {
		timestamp time.Time
		format    string
		expected  string
	}{
		{
			timestamp: now,
			expected:  fmt.Sprintf("BDAY:%s", now.Format("20060102")),
		},
		{
			timestamp: now,
			format:    time.RFC3339,
			expected:  fmt.Sprintf("BDAY:%s", now.Format(time.RFC3339)),
		},
	}

	for _, tc := range tt {
		got, err := vcard.Bday{Timestamp: tc.timestamp, TimeFormat: tc.format}.Format("4.0")
		if err != nil {
			t.Fatalf("expected to pass, but got: %v", err)
		}

		if got != tc.expected {
			t.Fatalf("expected %q, but got %q", tc.expected, got)
		}
	}
}

func TestFbURL(t *testing.T) {
	u, _ := url.Parse("http://example.com/fb/jdoe")
	tt := []struct {
		version     string
		uri         *url.URL
		expected    string
		expectedErr error
	}{
		{
			version:  "4.0",
			uri:      u,
			expected: "FBURL:http://example.com/fb/jdoe",
		},
		{
			version:     "3.0",
			uri:         u,
			expectedErr: vcard.ErrVersion,
		},
		{
			version:     "2.1",
			uri:         u,
			expectedErr: vcard.ErrVersion,
		},
	}

	for _, tc := range tt {
		got, err := vcard.FbURL{tc.uri}.Format(tc.version)
		if err != tc.expectedErr {
			t.Fatalf("expected err %v, but got: %v", tc.expectedErr, err)
		}

		if got != tc.expected {
			t.Fatalf("expected %q, but got %q", tc.expected, got)
		}

	}
}

func TestGender(t *testing.T) {
	tt := []struct {
		version     string
		val         string
		expected    string
		expectedErr error
	}{
		{
			version:  "4.0",
			val:      "F",
			expected: "GENDER:F",
		},
		{
			version:     "3.0",
			val:         "F",
			expectedErr: vcard.ErrVersion,
		},
		{
			version:     "2.0",
			val:         "F",
			expectedErr: vcard.ErrVersion,
		},
	}

	for _, tc := range tt {
		got, err := vcard.Gender{tc.val}.Format(tc.version)
		if err != tc.expectedErr {
			t.Fatalf("expected err %v, but got: %v", tc.expectedErr, err)
		}

		if got != tc.expected {
			t.Fatalf("expected %q, but got %q", tc.expected, got)
		}
	}
}

func TestGeo(t *testing.T) {
	tt := []struct {
		version  string
		lat      float64
		long     float64
		expected string
	}{
		{
			version:  "4.0",
			lat:      39.95,
			long:     -75.1667,
			expected: "GEO:geo:39.950000,-75.166700",
		},
		{
			version:  "3.0",
			lat:      39.95,
			long:     -75.1667,
			expected: "GEO:39.950000,-75.166700",
		},
		{
			version:  "2.1",
			lat:      39.95,
			long:     -75.1667,
			expected: "GEO:39.950000,-75.166700",
		},
	}

	for _, tc := range tt {
		got, err := vcard.Geo{tc.lat, tc.long}.Format(tc.version)
		if err != nil {
			t.Fatalf("expected to pass, but got %v", err)
		}

		if got != tc.expected {
			t.Fatalf("expected %q, but got %q", tc.expected, got)
		}
	}
}

func TestIMPP(t *testing.T) {
	tt := []struct {
		version     string
		platform    string
		handle      string
		expected    string
		expectedErr error
	}{
		{
			version:  "4.0",
			platform: "AIM",
			handle:   "test@example.com",
			expected: "IMPP:aim:test@example.com",
		},
		{
			version:     "2.1",
			platform:    "AIM",
			handle:      "test@example.com",
			expectedErr: vcard.ErrVersion,
		},
	}

	for _, tc := range tt {
		got, err := vcard.IMPP{tc.platform, tc.handle}.Format(tc.version)
		if err != tc.expectedErr {
			t.Fatalf("expected err %v, but got: %v", tc.expectedErr, err)
		}

		if got != tc.expected {
			t.Fatalf("expected %q, but got %q", tc.expected, got)
		}
	}
}

func TestKey(t *testing.T) {
	uri, _ := url.Parse("http://example.com/key.pgp")
	tt := []struct {
		version  string
		field    vcard.Key
		expected string
	}{
		{
			"2.1",
			vcard.Key{
				Type:   "PGP",
				Binary: true,
				Data:   "MIICajCCAdOgAwIBAgICBEUwDQYJKoZIhvcNAQEEBQAwdzELMAkGA1UEBhMCVVMxLDAqBgNVBAoTI05ldHNjYXBlIENvbW11bmljYXRpb25zIENvcnBvcmF0aW9uMRwwGgYDVQQLExNJbmZvcm1hdGlvbiBTeXN0",
			},
			"KEY;PGP;ENCODING=BASE64:MIICajCCAdOgAwIBAgICBEUwDQYJKoZIhvcNAQEEBQAwdzELMAkGA1UEBhMCVVMxLDAqBgNVBAoTI05ldHNjYXBlIENvbW11bmljYXRpb25zIENvcnBvcmF0aW9uMRwwGgYDVQQLExNJbmZvcm1hdGlvbiBTeXN0",
		},
		{
			"2.1",
			vcard.Key{
				Type: "PGP",
				URI:  uri,
			},
			"KEY;PGP:http://example.com/key.pgp",
		},

		{
			"3.0",
			vcard.Key{
				Type:   "PGP",
				Binary: true,
				Data:   "MIICajCCAdOgAwIBAgICBEUwDQYJKoZIhvcNAQEEBQAwdzELMAkGA1UEBhMCVVMxLDAqBgNVBAoTI05ldHNjYXBlIENvbW11bmljYXRpb25zIENvcnBvcmF0aW9uMRwwGgYDVQQLExNJbmZvcm1hdGlvbiBTeXN0",
			},
			"KEY;TYPE=PGP;ENCODING=b:MIICajCCAdOgAwIBAgICBEUwDQYJKoZIhvcNAQEEBQAwdzELMAkGA1UEBhMCVVMxLDAqBgNVBAoTI05ldHNjYXBlIENvbW11bmljYXRpb25zIENvcnBvcmF0aW9uMRwwGgYDVQQLExNJbmZvcm1hdGlvbiBTeXN0",
		},
		{
			"3.0",
			vcard.Key{
				Type: "PGP",
				URI:  uri,
			},
			"KEY;TYPE=PGP:http://example.com/key.pgp",
		},

		{
			"4.0",
			vcard.Key{
				Type:   "application/pgp-keys",
				Binary: true,
				Data:   "MIICajCCAdOgAwIBAgICBEUwDQYJKoZIhvcNAQEEBQAwdzELMAkGA1UEBhMCVVMxLDAqBgNVBAoTI05ldHNjYXBlIENvbW11bmljYXRpb25zIENvcnBvcmF0aW9uMRwwGgYDVQQLExNJbmZvcm1hdGlvbiBTeXN0",
			},
			"KEY:data:application/pgp-keys;base64,MIICajCCAdOgAwIBAgICBEUwDQYJKoZIhvcNAQEEBQAwdzELMAkGA1UEBhMCVVMxLDAqBgNVBAoTI05ldHNjYXBlIENvbW11bmljYXRpb25zIENvcnBvcmF0aW9uMRwwGgYDVQQLExNJbmZvcm1hdGlvbiBTeXN0",
		},
		{
			"4.0",
			vcard.Key{
				Type: "application/pgp-keys",
				URI:  uri,
			},
			"KEY;MEDIATYPE=application/pgp-keys:http://example.com/key.pgp",
		},
	}

	for _, tc := range tt {
		got, err := tc.field.Format(tc.version)
		if err != nil {
			t.Fatalf("expected to pass, but got: %v", err)
		}

		if got != tc.expected {
			t.Fatalf("expected %q, but got %q", tc.expected, got)
		}
	}
}

func TestKind(t *testing.T) {
	tt := []struct {
		version     string
		text        string
		expected    string
		expectedErr error
	}{
		{
			version:  "4.0",
			text:     "INDIVIDUAL",
			expected: "KIND:individual",
		},
		{
			version:     "2.1",
			text:        "INDIVIDUAL",
			expectedErr: vcard.ErrVersion,
		},
	}

	for _, tc := range tt {
		got, err := vcard.Kind{tc.text}.Format(tc.version)
		if err != tc.expectedErr {
			t.Fatalf("expected err %v, but got: %v", tc.expectedErr, err)
		}

		if got != tc.expected {
			t.Fatalf("expected %q, but got %q", tc.expected, got)
		}
	}
}
