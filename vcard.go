package vcard

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
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

func (v *VCard) isSupportedField(i int) bool {
	for _, ver := range v.Fields[i].Versions() {
		if ver == v.Version {
			return true
		}
	}

	return false
}

// Validate checks whether a VCard is valid by checking if all required fields are set for the version
// and if provided fields are supported by the required version
func (v *VCard) Validate() error {
	for i := range v.Fields {
		if !v.isSupportedField(i) {
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

// String implements the Stringer interface
func (v *VCard) String() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "BEGIN:VCARD\nVERSION:%s", v.Version)
	for i := range v.Fields {
		fmt.Fprintf(&b, "\n%s", v.Fields[i])
	}
	fmt.Fprintf(&b, "\nEND:VCARD")
	return b.String()
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
