package vcard

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// FieldFormatter is an interface containing the minimal behaviour of any VCard field formatter
type FieldFormatter interface {
	// Every FieldFormatter needs to implement the Stringer interface
	fmt.Stringer

	// Versions returns a slice containing which VCard versions this field supports
	Versions() []string
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

// N type definition to specify the components of the name of the object the vCard represents
type N struct {
	FamilyName        string
	GivenName         string
	AdditionalNames   string
	HonorificPrefixes string
	HonorificSuffixes string
}

// String implements the Stringer interface
func (f N) String() string {
	return fmt.Sprintf("N:%s;%s;%s;%s;%s",
		f.FamilyName,
		f.GivenName,
		f.AdditionalNames,
		f.HonorificPrefixes,
		f.HonorificSuffixes,
	)
}

// Versions implements the FieldFormatter interface.
func (f N) Versions() []string {
	return []string{"2.1", "3.0", "4.0"}
}

// FN type definition to specify the formatted text corresponding to the name of the object the vCard represents.
type FN struct {
	FormattedName string
}

// String implements the Stringer interface
func (f FN) String() string {
	return fmt.Sprintf("FN:%s", f.FormattedName)
}

// Versions implements the FieldFormatter interface
func (f FN) Versions() []string {
	return []string{"2.1", "3.0", "4.0"}
}

// Org type definition to specify the organizational name and units associated with the vCard.
type Org struct {
	Name  string
	Units []string
}

// String implements the Stringer interface
func (f Org) String() string {
	l := fmt.Sprintf("ORG:%s", f.Name)
	if len(f.Units) > 0 {
		l = fmt.Sprintf("%s;%s", l, strings.Join(f.Units, ";"))
	}
	return l
}

// Versions implements the FieldFormatter interface
func (f Org) Versions() []string {
	return []string{"2.1", "3.0", "4.0"}
}

// Title type definition to specify the job title, functional position or function of the object the vCard represents.
type Title struct {
	Title string
}

// String implements the Stringer interface
func (f Title) String() string {
	return fmt.Sprintf("TITLE:%s", f.Title)
}

// Versions implements the FieldFormatter interface
func (f Title) Versions() []string {
	return []string{"2.1", "3.0", "4.0"}
}

// Role type definition to specify information concerning the role, occupation,
// or business category of the object the vCard represents.
type Role struct {
	Role string
}

// String implements the Stringer interface
func (f Role) String() string {
	return fmt.Sprintf("ROLE:%s", f.Role)
}

// Versions implements the FieldFormatter interface
func (f Role) Versions() []string {
	return []string{"2.1", "3.0", "4.0"}
}

// Photo type definition to specify an image or photograph information that
// annotates some aspect of the object the vCard represents.
type Photo struct {
	Type   string
	URL    *url.URL
	Binary bool
	Data   []byte
}

// String implements the Stringer interface
func (f Photo) String() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "PHOTO")
	if f.Binary {
		fmt.Fprintf(&b, ";ENCODING=b;TYPE=%s:%s", f.Type, f.Data)
		return b.String()
	}

	if f.Type != "" {
		fmt.Fprintf(&b, ";TYPE=%s", f.Type)
	}
	fmt.Fprintf(&b, ";VALUE=uri:%s", f.URL.String())
	return b.String()
}

// Versions implements the FieldFormatter interface
func (f Photo) Versions() []string {
	return []string{"4.0"}
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

// String implements the Stringer interface
func (f Tel) String() string {
	t := strings.Join(f.Types, ",")
	if t == "" {
		t = TelVoice
	}
	return fmt.Sprintf("TEL;TYPE=%s:%s", t, f.Number)
}

// Versions implements the FieldFormatter interface
func (f Tel) Versions() []string {
	return []string{"2.1", "3.0", "4.0"}
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

// String implements the Stringer interface
func (f Adr) String() string {
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
	)
}

// Versions implements the FieldFormatter interface
func (f Adr) Versions() []string {
	return []string{"2.1", "3.0", "4.0"}
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

// String implements the Stringer interface
func (f Email) String() string {
	var b bytes.Buffer
	fmt.Fprint(&b, "EMAIL")
	if len(f.Types) > 0 {
		fmt.Fprintf(&b, ";TYPE=%s", strings.Join(f.Types, ","))
	}
	fmt.Fprintf(&b, ":%s", f.Email)
	return b.String()
}

// Versions implements the FieldFormatter interface
func (f Email) Versions() []string {
	return []string{"2.1", "3.0", "4.0"}
}

// Rev type definition to specify revision information about the current vCard.
type Rev struct {
	Timestamp time.Time
	Format    string
}

// String implements the Stringer interface
func (f Rev) String() string {
	format := "20060102T150405Z0700"
	if f.Format != "" {
		format = f.Format
	}
	return fmt.Sprintf("REV:%s", f.Timestamp.Format(format))
}

// Versions implements the FieldFormatter interface
func (f Rev) Versions() []string {
	return []string{"2.1", "3.0", "4.0"}
}
