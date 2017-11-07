package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Action Action object
// swagger:model action
type Action struct {

	// Command to execute
	// Required: true
	Command *string `json:"command"`

	// Command options object
	Options interface{} `json:"options,omitempty"`
}

// Validate validates this action
func (m *Action) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommand(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var actionTypeCommandPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cancel"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		actionTypeCommandPropEnum = append(actionTypeCommandPropEnum, v)
	}
}

const (
	// ActionCommandCancel captures enum value "cancel"
	ActionCommandCancel string = "cancel"
)

// prop value enum
func (m *Action) validateCommandEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, actionTypeCommandPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Action) validateCommand(formats strfmt.Registry) error {

	if err := validate.Required("command", "body", m.Command); err != nil {
		return err
	}

	// value enum
	if err := m.validateCommandEnum("command", "body", *m.Command); err != nil {
		return err
	}

	return nil
}