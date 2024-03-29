package vcard

import (
	"bytes"
	"fmt"
	"image/png"
	"os"
	"reflect"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

/*
	Implement other versions than 4.0 for PHOTO field
*/

type version struct {
	// required contains all the required fields for a version
	required []reflect.Type
}

// versions contains all supported versions and their version specific settings
var versions = map[string]version{
	"2.1": version{
		required: []reflect.Type{
			reflect.TypeOf(N{}),
		},
	},
	"3.0": version{
		required: []reflect.Type{
			reflect.TypeOf(N{}),
			reflect.TypeOf(FN{}),
		},
	},
	"4.0": version{
		required: []reflect.Type{
			reflect.TypeOf(FN{}),
		},
	},
}

// VCard represents a formattable vcard
type VCard struct {
	Version string
	Fields  []FieldFormatter
}

func (v *VCard) fieldMap() map[reflect.Type]int {
	fmap := make(map[reflect.Type]int, len(v.Fields)-1)
	for i := range v.Fields {
		t := reflect.TypeOf(v.Fields[i])
		if _, ok := fmap[t]; !ok {
			fmap[t] = 1
		} else {
			fmap[t]++
		}
	}

	return fmap
}

// func (v *VCard) isSupportedField(i int) bool {
// 	for _, ver := range v.Fields[i].Versions() {
// 		if ver == v.Version {
// 			return true
// 		}
// 	}

// 	return false
// }

// Validate checks whether a VCard is valid by checking if all required fields are set for the version
// and if provided fields are supported by the required version
func (v *VCard) Validate() error {
	for i := range v.Fields {
		if _, err := v.Fields[i].Format(v.Version); err == ErrVersion {
			return fmt.Errorf("%T is an unsupported field for vCard version %s", v.Fields[i], v.Version)
		}
	}

	fmap := v.fieldMap()
	for _, req := range versions[v.Version].required {
		if _, ok := fmap[req]; !ok {
			return fmt.Errorf("%s is a required field for vCard version %s", req, v.Version)
		}
	}
	return nil
}

// Generate will generate the vcard string
func (v *VCard) Generate() (string, error) {
	var b bytes.Buffer
	fmt.Fprintf(&b, "BEGIN:VCARD\nVERSION:%s", v.Version)
	for i := range v.Fields {
		o, err := v.Fields[i].Format(v.Version)
		if err != nil {
			return "", err
		}
		fmt.Fprintf(&b, "\n%s", o)
	}
	fmt.Fprintf(&b, "\nEND:VCARD")
	return b.String(), nil
}

// QR creates a QR code of the VCard
func (v *VCard) QR(x, y int) (barcode.Barcode, error) {
	if err := v.Validate(); err != nil {
		return nil, err
	}

	// Create the barcode
	vcard, err := v.Generate()
	if err != nil {
		return nil, err
	}
	qrCode, err := qr.Encode(vcard, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}

	// Scale the barcode to 200x200 pixels
	return barcode.Scale(qrCode, x, y)
}

// QRPng creates a png file containing a QR code of the VCard
func (v *VCard) QRPng(x, y int, filename string) error {
	qrCode, err := v.QR(x, y)
	if err != nil {
		return err
	}
	// create the output file
	file, _ := os.Create(filename)
	defer file.Close()

	// encode the barcode as png
	return png.Encode(file, qrCode)
}

// VersionError is used to to return an error when the user has selected a wrong version
type VersionError struct{}

// Error implements the error interface
func (err *VersionError) Error() string {
	var v []string
	for k := range versions {
		v = append(v, k)
	}

	return fmt.Sprintf("invalid version, supported versions: %s", strings.Join(v, ", "))
}

// New returns a new VCard based upon the provided version and fields,
// returns an error if an invalid version is provided or if validation
// of the vcard fails
func New(v string, fields ...FieldFormatter) (*VCard, error) {
	_, ok := versions[v]
	if !ok {
		return nil, &VersionError{}
	}

	card := VCard{
		Version: v,
		Fields:  fields,
	}

	if err := card.Validate(); err != nil {
		return nil, err
	}

	return &card, nil
}
