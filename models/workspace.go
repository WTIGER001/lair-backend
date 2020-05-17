// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Workspace Workspace
//
// swagger:model Workspace
type Workspace struct {

	// Amount of CPU being requested to run this algorithm
	Cpus float64 `json:"cpus,omitempty"`

	// Workspace ID
	Description string `json:"description,omitempty"`

	// Full docker image name. Including the repo
	DockerImage string `json:"dockerImage,omitempty"`

	// Workspace ID
	ID string `json:"id,omitempty"`

	// Image of the workspace
	Image string `json:"image,omitempty"`

	// Amount of Memory in MBs being requested
	Memory float64 `json:"memory,omitempty"`

	// Workspace ID
	Name string `json:"name,omitempty"`

	// If the workpsace is shared for others to clone
	Shared bool `json:"shared,omitempty"`

	// Timeout, in minutes, that a workspace will be kept active, default 30 min
	Timeout int64 `json:"timeout,omitempty"`

	// Location of the workspace volume
	WorkpsaceLocation string `json:"workpsaceLocation,omitempty"`
}

// Validate validates this workspace
func (m *Workspace) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Workspace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Workspace) UnmarshalBinary(b []byte) error {
	var res Workspace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
