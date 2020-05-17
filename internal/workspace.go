package internal

import (
	"github.com/wtiger001/lair-backend/restapi/operations"
	middleware "github.com/go-openapi/runtime/middleware"

)

// DS - Data store to use
var DS Datastore

// GetWorkspaces - Gets the workspaces
func GetWorkspaces(operations.GetWorkspacesParams) middleware.Responder {
	model, err := DS.GetWorkspaces()
	
	if(err != nil) {
		return operations.NewGetWorkspacesInternalServerError()
	}

	payload := operations.GetWorkspacesOKBody{}
	payload.User = ""
	payload.Workspaces = model

	return operations.NewGetWorkspacesOK().WithPayload(&payload)
}

// GetWorkspaceByID - Retrieves a single workspace
func GetWorkspaceByID(params operations.GetWorkspacesIDParams) middleware.Responder {
	model, err := DS.Retrieve(params.ID)

	if(err != nil) {
		return operations.NewGetWorkspacesInternalServerError()
	}

	if (model == nil) {
		return operations.NewLaunchWorkpaceByIDNotFound()
	}

	return operations.NewGetWorkspacesIDOK().WithPayload(model)
}


// SaveWorkspace - Retrieves a single workspace
func SaveWorkspace(params operations.PostWorkpaceByIDParams) middleware.Responder {
	 err := DS.Update(params.Workspace)

	if(err != nil) {
		return operations.NewPostWorkpaceByIDInternalServerError()
	}

	return operations.NewPostWorkpaceByIDOK()
}
