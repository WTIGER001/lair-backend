package internal

import (
	"errors"

	"github.com/wtiger001/lair-backend/models"
)

// RAMDatastore - Holds any information for the RAM store
type RAMDatastore struct {
	Workspaces []*models.Workspace
	Statuses   []*models.WorkspaceStatus
}

/*
	Connect() error
	Close() error
	GetWorkspaces() (*[]models.Workspac, error)
	Retrieve(id string) (*models.Workspace, error)
	Update(workspace *Workspace) error
	Delete(id string) error
*/

// NewRAMDatastore - Creates a new RAM Datastore
func NewRAMDatastore() *RAMDatastore {
	ds := new(RAMDatastore)
	return ds
}

// Connect - Trival Connect
func (ds *RAMDatastore) Connect() error {
	ds.Workspaces = make([]*models.Workspace, 0)
	ds.Statuses = make([]*models.WorkspaceStatus, 0)

	return nil
}

// Close - Trival Close
func (ds *RAMDatastore) Close() error {
	return nil
}

// GetWorkspaces - Gets the workspaces from an in memory array
func (ds *RAMDatastore) GetWorkspaces() ([]*models.Workspace, error) {
	return ds.Workspaces, nil
}

// Retrieve - Gets from array or returns error
func (ds *RAMDatastore) Retrieve(id string) (*models.Workspace, error) {
	index := ds.IndexOf(id)
	if index >= 0 {
		return ds.Workspaces[index], nil
	}

	return nil, errors.New("No Workspace found")
}

// Update - Replace or add
func (ds *RAMDatastore) Update(workspace *models.Workspace) error {
	index := ds.IndexOf(workspace.ID)
	if index >= 0 {
		ds.Workspaces[index] = workspace
	} else {
		ds.Workspaces = append(ds.Workspaces, workspace)
	}
	return nil
}

// Delete - Deletes
func (ds *RAMDatastore) Delete(id string) error {
	i := ds.IndexOf(id)
	if i >= 0 {

		// Remove the element at index i from a.
		copy(ds.Workspaces[i:], ds.Workspaces[i+1:])         // Shift a[i+1:] left one index.
		ds.Workspaces = ds.Workspaces[:len(ds.Workspaces)-1] // Truncate slice.
	}
	return nil
}

// IndexOf - Finds the index
func (ds *RAMDatastore) IndexOf(id string) int {
	for i := range ds.Workspaces {
		if ds.Workspaces[i].ID == id {
			return i
		}
	}
	return -1
}
