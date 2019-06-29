package vcard

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

const (
	dateFormat     = "20060102"
	dateTimeFormat = "20060102T150405Z0700"
)

// FieldFormatter is an interface containing the minimal behaviour of any VCard field formatter
// Receives the version as a string, returns a formatted string and an error
// Current versions to support are 2.1, 3.0 and 4,0
type FieldFormatter interface {
	Format(string) (string, error)
}

// BEGIN:VCARD
// VERSION:4.0
// N:Gump;Forrest;;Mr.;
// FN:Forrest Gump
// ORG:Bubba Gump Shrimp Co.
// TITLE:Shrimp Man
// PHOTO;MEDIATYPE=image/gif:http://www.example.com/dir_photos/my_photo.gif
// TEL;TYPE=work,voice;VALUE=uri:tel:+1-111-555-1212
// TEL;TYPE=home,voice;VALUE=uri:tel:+1-404-555-1212
// ADR;TYPE=WORK;PREF=1;LABEL="100 Waters Edge\nBaytown\, LA 30314\nUnited States of America":;;100 Waters Edge;Baytown;LA;30314;United States of America
// ADR;TYPE=HOME;LABEL="42 Plantation St.\nBaytown\, LA 30314\nUnited States of America":;;42 Plantation St.;Baytown;LA;30314;United States of America
// EMAIL:forrestgump@example.com
// REV:20080424T195243Z
// x-qq:21588891
// END:VCARD

// ErrVersion is used by FieldFormatters when a request is made for an unsupported vcard version
var ErrVersion = errors.New("unsupported verson")

// N type definition to specify the components of the name of the object the vCard represents
type N struct {
	FamilyName        string
	GivenName         string
	AdditionalNames   string
	HonorificPrefixes string
	HonorificSuffixes string
}

// Format implements the FieldFormatter interface
func (f N) Format(v string) (string, error) {
	switch v {
	case "2.1", "3.0", "4.0":
		return fmt.Sprintf("N:%s;%s;%s;%s;%s",
			f.FamilyName,
			f.GivenName,
			f.AdditionalNames,
			f.HonorificPrefixes,
			f.HonorificSuffixes,
		), nil
	}
	return "", ErrVersion
}

// FN type definition to specify the formatted text corresponding to the name of the object the vCard represents.
type FN struct {
	FormattedName string
}

// Format implements the FieldFormatter interface
func (f FN) Format(v string) (string, error) {
	switch v {
	case "2.1", "3.0", "4.0":
		return fmt.Sprintf("FN:%s", f.FormattedName), nil
	}
	return "", ErrVersion
}

// Org type definition to specify the organizational name and units associated with the vCard.
type Org struct {
	Name  string
	Units []string
}

// Format implements the FieldFormatter interface
func (f Org) Format(v string) (string, error) {
	switch v {
	case "2.1", "3.0", "4.0":
		l := fmt.Sprintf("ORG:%s", f.Name)
		if len(f.Units) > 0 {
			l = fmt.Sprintf("%s;%s", l, strings.Join(f.Units, ";"))
		}
		return l, nil
	}
	return "", ErrVersion
}

// Title type definition to specify the job title, functional position or function of the object the vCard represents.
type Title struct {
	Title string
}

// Format implements the FieldFormatter interface
func (f Title) Format(v string) (string, error) {
	switch v {
	case "2.1", "3.0", "4.0":
		return fmt.Sprintf("TITLE:%s", f.Title), nil
	}
	return "", ErrVersion
}

// Role type definition to specify information concerning the role, occupation,
// or business category of the object the vCard represents.
type Role struct {
	Role string
}

// Format implements the FieldFormatter interface
func (f Role) Format(v string) (string, error) {
	switch v {
	case "2.1", "3.0", "4.0":
		return fmt.Sprintf("ROLE:%s", f.Role), nil
	}
	return "", ErrVersion
}

func mediaString(v, field, tp, b64 string, uri *url.URL) (string, error) {
	var b bytes.Buffer
	fmt.Fprintf(&b, field)

	switch v {
	case "2.1":
		if tp != "" {
			fmt.Fprintf(&b, ";%s", tp)
		}
		if len(b64) != 0 {
			fmt.Fprintf(&b, ";ENCODING=BASE64:%s", b64)
			return b.String(), nil
		}

		fmt.Fprintf(&b, ":%s", uri)
		return b.String(), nil

	case "3.0":
		if tp != "" {
			fmt.Fprintf(&b, ";TYPE=%s", tp)
		}

		if len(b64) > 0 {
			fmt.Fprintf(&b, ";ENCODING=b:%s", b64)
			return b.String(), nil
		}

		fmt.Fprintf(&b, ";VALUE=uri:%s", uri)
		return b.String(), nil
	case "4.0":
		if len(b64) != 0 {
			fmt.Fprintf(&b, ":data:%s;base64,%s", tp, b64)
			return b.String(), nil
		}

		if tp != "" {
			fmt.Fprintf(&b, ";MEDIATYPE=%s", tp)
		}
		fmt.Fprintf(&b, ":%s", uri)
		return b.String(), nil
	}
	return "", ErrVersion
}

// Photo type definition to specify an image or photograph information that
// annotates some aspect of the object the vCard represents.
type Photo struct {
	Type       string
	URI        *url.URL
	Base64Data string
}

// Format implements the FieldFormatter interface
func (f Photo) Format(v string) (string, error) {
	return mediaString(v, "PHOTO", f.Type, f.Base64Data, f.URI)
}

const (
	// TelHome is a telephone number associated with a residence
	TelHome = "home"

	// TelMsg indicates that the telephone number has voice messaging support
	TelMsg = "msg"

	// TelWork indicates a telephone number associated with a place of work
	TelWork = "work"

	// TelPref indicates a preferred-use telephone number
	TelPref = "pref"

	// TelVoice indicates a voice telephone number
	TelVoice = "voice"

	// TelFax indicates a facsimile telephone number
	TelFax = "fax"

	// TelCell indicates a cellular phone number
	TelCell = "cell"

	// TelVideo indicates video calling support
	TelVideo = "video"

	// TelPager indicates a paging device telephone number
	TelPager = "pager"

	// TelBBS indicates a buletin board system telephone number
	TelBBS = "bbs"

	// TelModem indicates a moden telephone number
	TelModem = "modem"

	// TelCar indicates a car-phone telephone
	TelCar = "car"

	// TelISDN indicates an ISDN service telephone number
	TelISDN = "isdn"

	// TelPCS indicates a personal communication service telephone number
	TelPCS = "pcs"
)

// Tel type definition to the telephone number for telephony communication with the object the vCard represents.
type Tel struct {
	Types  []string
	Number string
}

// Format implements the FieldFormatter interface
func (f Tel) Format(v string) (string, error) {
	switch v {
	case "2.1", "3.0", "4.0":
		t := strings.Join(f.Types, ",")
		if t == "" {
			t = TelVoice
		}
		return fmt.Sprintf("TEL;TYPE=%s:%s", t, f.Number), nil
	}
	return "", ErrVersion
}

const (
	// AdrDom indicates a domestic address.
	AdrDom = "dom"

	// AdrIntl indicates an international address.
	AdrIntl = "intl"

	// AdrPostal indicates a postal address.
	AdrPostal = "postal"

	// AdrParcel indicates a parcel delivery address.
	AdrParcel = "parcel"

	// AdrHome indicates a home delivery address.
	AdrHome = "home"

	// AdrWork indicates a delivery address for a place of work.
	AdrWork = "work"

	// AdrPref indicates a preferred address for when more addresses are supplied.
	AdrPref = "pref"
)

// Adr type definition to specify the components of the delivery address for the vCard object.
type Adr struct {
	Types           []string
	PostOfficeBox   string
	ExtendedAddress string
	StreetAddress   string
	Locality        string
	Region          string
	PostalCode      string
	CountryName     string
}

// Format implements the FieldFormatter interface
func (f Adr) Format(v string) (string, error) {
	switch v {
	case "2.1", "3.0", "4.0":
		t := strings.Join(f.Types, ",")
		if t == "" {
			t = "intl,postal,parcel,work"
		}
		return fmt.Sprintf("ADR;TYPE=%s:%s;%s;%s;%s;%s;%s;%s",
			t,
			f.PostOfficeBox,
			f.ExtendedAddress,
			f.StreetAddress,
			f.Locality,
			f.Region,
			f.PostalCode,
			f.CountryName,
		), nil
	}
	return "", ErrVersion
}

const (
	// EmailInternet indicates an internet addressing type.
	EmailInternet = "internet"

	// EmailX400 indicates a X.400 addressing type.
	EmailX400 = "x400"

	// EmailPref indicates a preferred address for when more addresses are supplied
	EmailPref = "pref"
)

// Email type definition to specify the formatted text corresponding to the name of the object the vCard represents.
type Email struct {
	Types []string
	Email string
}

// Format implements the FieldFormatter interface
func (f Email) Format(v string) (string, error) {
	switch v {
	case "2.1", "3.0", "4.0":
		var b bytes.Buffer
		fmt.Fprint(&b, "EMAIL")
		if len(f.Types) > 0 {
			fmt.Fprintf(&b, ";TYPE=%s", strings.Join(f.Types, ","))
		}
		fmt.Fprintf(&b, ":%s", f.Email)
		return b.String(), nil
	}
	return "", ErrVersion
}

// Rev type definition to specify revision information about the current vCard.
type Rev struct {
	Timestamp  time.Time
	TimeFormat string
}

// Format implements the FieldFormatter interface
func (f Rev) Format(v string) (string, error) {
	switch v {
	case "2.1", "3.0", "4.0":
		format := dateTimeFormat
		if f.TimeFormat != "" {
			format = f.TimeFormat
		}
		return fmt.Sprintf("REV:%s", f.Timestamp.Format(format)), nil
	}
	return "", ErrVersion
}

// Agent type definition to specify information about another person who will
// act on behalf of the individual or resource associated with the
// vCard. Can contain a VCard of the agent or a string.
type Agent struct {
	VCard *VCard
	Text  string
}

// Format implements the FieldFormatter interface
func (f Agent) Format(v string) (string, error) {
	switch v {
	case "2.1", "3.0":
		if f.VCard != nil {
			vcard, err := f.VCard.Generate()
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("AGENT:%s", vcard), nil
		}
		return fmt.Sprintf("AGENT:%s", f.Text), nil
	}
	return "", ErrVersion
}

// Anniversary type definition to specify a person's anniversary/
type Anniversary struct {
	Date       time.Time
	TimeFormat string
}

// Format implements the FieldFormatter interface
func (f Anniversary) Format(v string) (string, error) {
	switch v {
	case "4.0":
		format := dateFormat
		if f.TimeFormat != "" {
			format = f.TimeFormat
		}
		return fmt.Sprintf("ANNIVERSARY:%s", f.Date.Format(format)), nil
	}
	return "", ErrVersion
}

// Bday type definition to specify the date of birth of the individual associated with the vCard.
type Bday struct {
	Timestamp  time.Time
	TimeFormat string
}

// Format implements the FieldFormatter interface
func (f Bday) Format(v string) (string, error) {
	switch v {
	case "2.1", "3.0", "4.0":
		format := dateFormat
		if f.TimeFormat != "" {
			format = f.TimeFormat
		}
		return fmt.Sprintf("BDAY:%s", f.Timestamp.Format(format)), nil
	}
	return "", ErrVersion
}

// TODO: CLIENTPIDMAP

// FbURL type definition to specify a URL that shows when the person is "free" or "busy" on their calendar.
type FbURL struct {
	*url.URL
}

// Format implements the FieldFormatter interface
func (f FbURL) Format(v string) (string, error) {
	switch v {
	case "4.0":
		return fmt.Sprintf("FBURL:%s", f.URL), nil
	}
	return "", ErrVersion
}

// Gender type definition to specify a person's gender.
type Gender struct {
	Val string
}

// Format implements the FieldFormatter interface
func (f Gender) Format(v string) (string, error) {
	switch v {
	case "4.0":
		return fmt.Sprintf("GENDER:%s", f.Val), nil
	}
	return "", ErrVersion
}

// Geo type definition to specify a latitude and longitude.
// For vcard version 4.0
type Geo struct {
	Lat  float64
	Long float64
}

// Format implements the FieldFormatter interface
func (f Geo) Format(v string) (string, error) {
	switch v {
	case "2.1", "3.0":
		return fmt.Sprintf("GEO:%f,%f", f.Lat, f.Long), nil
	case "4.0":
		return fmt.Sprintf("GEO:geo:%f,%f", f.Lat, f.Long), nil
	}
	return "", ErrVersion
}

// IMPP type definition to specify instant messenger handle.
type IMPP struct {
	Platform string
	Handle   string
}

// Format implements the FieldFormatter interface
func (f IMPP) Format(v string) (string, error) {
	switch v {
	case "3.0", "4.0":
		return fmt.Sprintf("IMPP:%s:%s", strings.ToLower(f.Platform), f.Handle), nil
	}
	return "", ErrVersion
}

// IMPP type definition to specify instant messenger handle.
type Key struct {
	Type   string
	URI    *url.URL
	Data   string
	Binary bool
}

// Format implements the FieldFormatter interface
func (f Key) Format(v string) (string, error) {
	var b bytes.Buffer
	fmt.Fprintf(&b, "KEY")

	switch v {
	case "2.1":
		if len(f.Data) > 0 {
			encoding := ""
			if f.Binary {
				encoding = ";ENCODING=BASE64"
			}
			fmt.Fprintf(&b, ";%s%s:%s", f.Type, encoding, f.Data)
			return b.String(), nil
		}

		fmt.Fprintf(&b, ";%s:%s", f.Type, f.URI)
		return b.String(), nil

	case "3.0":
		if len(f.Data) > 0 {
			encoding := ""
			if f.Binary {
				encoding = ";ENCODING=b"
			}
			fmt.Fprintf(&b, ";TYPE=%s%s:%s", f.Type, encoding, f.Data)
			return b.String(), nil
		}

		fmt.Fprintf(&b, ";TYPE=%s:%s", f.Type, f.URI)
		return b.String(), nil
	case "4.0":
		if len(f.Data) > 0 {
			encoding := ""
			if f.Binary {
				encoding = "base64,"
			}
			fmt.Fprintf(&b, ":data:%s;%s%s", f.Type, encoding, f.Data)
			return b.String(), nil
		}

		fmt.Fprintf(&b, ";MEDIATYPE=%s:%s", f.Type, f.URI)
		return b.String(), nil
	}
	return "", ErrVersion
}
