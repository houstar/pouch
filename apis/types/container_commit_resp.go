// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ContainerCommitResp response of commit container for the remote API: POST /commit
// swagger:model ContainerCommitResp
type ContainerCommitResp struct {

	// ID uniquely identifies an image committed by a container
	ID string `json:"Id,omitempty"`
}

// Validate validates this container commit resp
func (m *ContainerCommitResp) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContainerCommitResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerCommitResp) UnmarshalBinary(b []byte) error {
	var res ContainerCommitResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
