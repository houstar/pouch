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

// ContainerState container state
// swagger:model ContainerState
type ContainerState struct {

	// Whether this container is dead.
	// Required: true
	Dead bool `json:"Dead"`

	// The error message of this container
	// Required: true
	Error string `json:"Error"`

	// The last exit code of this container
	// Required: true
	ExitCode int64 `json:"ExitCode"`

	// The time when this container last exited.
	// Required: true
	FinishedAt string `json:"FinishedAt"`

	// Whether this container has been killed because it ran out of memory.
	// Required: true
	OOMKilled bool `json:"OOMKilled"`

	// Whether this container is paused.
	// Required: true
	Paused bool `json:"Paused"`

	// The process ID of this container
	// Required: true
	Pid int64 `json:"Pid"`

	// Whether this container is restarting.
	// Required: true
	Restarting bool `json:"Restarting"`

	// Whether this container is running.
	//
	// Note that a running container can be _paused_. The `Running` and `Paused`
	// booleans are not mutually exclusive:
	//
	// When pausing a container (on Linux), the cgroups freezer is used to suspend
	// all processes in the container. Freezing the process requires the process to
	// be running. As a result, paused containers are both `Running` _and_ `Paused`.
	//
	// Use the `Status` field instead to determine if a container's state is "running".
	//
	// Required: true
	Running bool `json:"Running"`

	// The time when this container was last started.
	// Required: true
	StartedAt string `json:"StartedAt"`

	// status
	// Required: true
	Status Status `json:"Status"`
}

// Validate validates this container state
func (m *ContainerState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDead(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExitCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinishedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOOMKilled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaused(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestarting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunning(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerState) validateDead(formats strfmt.Registry) error {

	if err := validate.Required("Dead", "body", bool(m.Dead)); err != nil {
		return err
	}

	return nil
}

func (m *ContainerState) validateError(formats strfmt.Registry) error {

	if err := validate.RequiredString("Error", "body", string(m.Error)); err != nil {
		return err
	}

	return nil
}

func (m *ContainerState) validateExitCode(formats strfmt.Registry) error {

	if err := validate.Required("ExitCode", "body", int64(m.ExitCode)); err != nil {
		return err
	}

	return nil
}

func (m *ContainerState) validateFinishedAt(formats strfmt.Registry) error {

	if err := validate.RequiredString("FinishedAt", "body", string(m.FinishedAt)); err != nil {
		return err
	}

	return nil
}

func (m *ContainerState) validateOOMKilled(formats strfmt.Registry) error {

	if err := validate.Required("OOMKilled", "body", bool(m.OOMKilled)); err != nil {
		return err
	}

	return nil
}

func (m *ContainerState) validatePaused(formats strfmt.Registry) error {

	if err := validate.Required("Paused", "body", bool(m.Paused)); err != nil {
		return err
	}

	return nil
}

func (m *ContainerState) validatePid(formats strfmt.Registry) error {

	if err := validate.Required("Pid", "body", int64(m.Pid)); err != nil {
		return err
	}

	return nil
}

func (m *ContainerState) validateRestarting(formats strfmt.Registry) error {

	if err := validate.Required("Restarting", "body", bool(m.Restarting)); err != nil {
		return err
	}

	return nil
}

func (m *ContainerState) validateRunning(formats strfmt.Registry) error {

	if err := validate.Required("Running", "body", bool(m.Running)); err != nil {
		return err
	}

	return nil
}

func (m *ContainerState) validateStartedAt(formats strfmt.Registry) error {

	if err := validate.RequiredString("StartedAt", "body", string(m.StartedAt)); err != nil {
		return err
	}

	return nil
}

func (m *ContainerState) validateStatus(formats strfmt.Registry) error {

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerState) UnmarshalBinary(b []byte) error {
	var res ContainerState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
