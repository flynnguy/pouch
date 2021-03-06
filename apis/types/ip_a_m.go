// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IPAM represents IP Address Management
// swagger:model IPAM

type IPAM struct {

	// config
	Config []IPAMConfig `json:"Config"`

	// driver
	Driver string `json:"Driver,omitempty"`

	// options
	Options map[string]string `json:"Options,omitempty"`
}

/* polymorph IPAM Config false */

/* polymorph IPAM Driver false */

/* polymorph IPAM Options false */

// Validate validates this IP a m
func (m *IPAM) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPAM) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Config) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPAM) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPAM) UnmarshalBinary(b []byte) error {
	var res IPAM
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
