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

// ExecCreateConfig is a small subset of the Config struct that holds the configuration.
// swagger:model ExecCreateConfig
type ExecCreateConfig struct {

	// Attach the standard error
	AttachStderr bool `json:"AttachStderr,omitempty"`

	// Attach the standard input, makes possible user interaction
	AttachStdin bool `json:"AttachStdin,omitempty"`

	// Attach the standard output
	AttachStdout bool `json:"AttachStdout,omitempty"`

	// Execution commands and args
	// Required: true
	// Min Items: 1
	Cmd []string `json:"Cmd"`

	// Execute in detach mode
	Detach bool `json:"Detach,omitempty"`

	// Escape keys for detach
	DetachKeys string `json:"DetachKeys,omitempty"`

	// envs for exec command in container
	Env []string `json:"Env"`

	// Is the container in privileged mode
	Privileged bool `json:"Privileged,omitempty"`

	// Attach standard streams to a tty
	Tty bool `json:"Tty,omitempty"`

	// User that will run the command
	User string `json:"User,omitempty"`
}

// Validate validates this exec create config
func (m *ExecCreateConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCmd(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExecCreateConfig) validateCmd(formats strfmt.Registry) error {

	if err := validate.Required("Cmd", "body", m.Cmd); err != nil {
		return err
	}

	iCmdSize := int64(len(m.Cmd))

	if err := validate.MinItems("Cmd", "body", iCmdSize, 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExecCreateConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExecCreateConfig) UnmarshalBinary(b []byte) error {
	var res ExecCreateConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
