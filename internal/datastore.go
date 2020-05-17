package internal

import (
	"github.com/wtiger001/lair-backend/models"
)

// Datastore - The storage interface to abstract how models are stored in the database
type Datastore interface {
	Connect() error

	Close() error

	GetWorkspaces() ([]*models.Workspace, error)

	Retrieve(id string) (*models.Workspace, error)

	Update(workspace *models.Workspace) error

	Delete(id string) error
}