package internal

import (
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/wtiger001/lair-backend/restapi/operations"
)

// DS - Data store to use
var DS Datastore

// GetWorkspaces - Gets the workspaces
func GetWorkspaces(operations.GetWorkspacesParams) middleware.Responder {
	model, err := DS.GetWorkspaces()

	if err != nil {
		return operations.NewGetWorkspacesInternalServerError()
	}

	return operations.NewGetWorkspacesOK().WithPayload(model)
}

// GetWorkspaceByID - Retrieves a single workspace
func GetWorkspaceByID(params operations.GetWorkspaceByIDParams) middleware.Responder {
	model, err := DS.Retrieve(params.ID)

	if err != nil {
		return operations.NewGetWorkspacesInternalServerError()
	}

	if model == nil {
		return operations.NewLaunchWorkpaceByIDNotFound()
	}

	return operations.NewGetWorkspaceByIDOK().WithPayload(model)
}

// SaveWorkspace - Retrieves a single workspace
func SaveWorkspace(params operations.PostWorkpaceByIDParams) middleware.Responder {
	err := DS.Update(params.Workspace)

	if err != nil {
		return operations.NewPostWorkpaceByIDInternalServerError()
	}

	return operations.NewPostWorkpaceByIDOK()
}

// DeleteWorkspace - Deletes a workspace
func DeleteWorkspace(params operations.DeleteWorkpaceByIDParams) middleware.Responder {
	err := DS.Delete(params.ID)

	if err != nil {
		return operations.NewDeleteWorkpaceByIDInternalServerError()
	}

	return operations.NewDeleteWorkpaceByIDOK()
}
