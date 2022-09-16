// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SharedProcessorPoolPlacementGroup shared processor pool placement group
//
// swagger:model SharedProcessorPoolPlacementGroup
type SharedProcessorPoolPlacementGroup struct {

	// The id of the Shared Processor Pool Placement Group
	// Required: true
	ID *string `json:"id"`

	// The name of the Shared Processor Pool Placement Group
	// Required: true
	Name *string `json:"name"`

	// The Shared Processor Pool Placement Group policy
	// Required: true
	Policy *string `json:"policy"`
}

// Validate validates this shared processor pool placement group
func (m *SharedProcessorPoolPlacementGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SharedProcessorPoolPlacementGroup) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SharedProcessorPoolPlacementGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SharedProcessorPoolPlacementGroup) validatePolicy(formats strfmt.Registry) error {

	if err := validate.Required("policy", "body", m.Policy); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this shared processor pool placement group based on context it is used
func (m *SharedProcessorPoolPlacementGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SharedProcessorPoolPlacementGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SharedProcessorPoolPlacementGroup) UnmarshalBinary(b []byte) error {
	var res SharedProcessorPoolPlacementGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}