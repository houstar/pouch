// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProcessConfig ExecProcessConfig holds information about the exec process.
// swagger:model ProcessConfig
type ProcessConfig struct {

	// arguments
	// Required: true
	Arguments []string `json:"arguments"`

	// entrypoint
	// Required: true
	Entrypoint string `json:"entrypoint"`

	// privileged
	// Required: true
	Privileged bool `json:"privileged"`

	// tty
	// Required: true
	Tty bool `json:"tty"`

	// user
	// Required: true
	User string `json:"user"`
}

// Validate validates this process config
func (m *ProcessConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArguments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntrypoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivileged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessConfig) validateArguments(formats strfmt.Registry) error {

	if err := validate.Required("arguments", "body", m.Arguments); err != nil {
		return err
	}

	return nil
}

func (m *ProcessConfig) validateEntrypoint(formats strfmt.Registry) error {

	if err := validate.RequiredString("entrypoint", "body", string(m.Entrypoint)); err != nil {
		return err
	}

	return nil
}

func (m *ProcessConfig) validatePrivileged(formats strfmt.Registry) error {

	if err := validate.Required("privileged", "body", bool(m.Privileged)); err != nil {
		return err
	}

	return nil
}

func (m *ProcessConfig) validateTty(formats strfmt.Registry) error {

	if err := validate.Required("tty", "body", bool(m.Tty)); err != nil {
		return err
	}

	return nil
}

func (m *ProcessConfig) validateUser(formats strfmt.Registry) error {

	if err := validate.RequiredString("user", "body", string(m.User)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProcessConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessConfig) UnmarshalBinary(b []byte) error {
	var res ProcessConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
