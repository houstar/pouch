// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// BlkioStatEntry BlkioStatEntry is one small entity to store a piece of Blkio stats
// swagger:model BlkioStatEntry
type BlkioStatEntry struct {

	// major
	Major uint64 `json:"major,omitempty"`

	// minor
	Minor uint64 `json:"minor,omitempty"`

	// op
	Op string `json:"op,omitempty"`

	// value
	Value uint64 `json:"value,omitempty"`
}

// Validate validates this blkio stat entry
func (m *BlkioStatEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BlkioStatEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlkioStatEntry) UnmarshalBinary(b []byte) error {
	var res BlkioStatEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
