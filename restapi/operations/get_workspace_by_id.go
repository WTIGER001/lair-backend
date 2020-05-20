// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetWorkspaceByIDHandlerFunc turns a function with the right signature into a get workspace by Id handler
type GetWorkspaceByIDHandlerFunc func(GetWorkspaceByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetWorkspaceByIDHandlerFunc) Handle(params GetWorkspaceByIDParams) middleware.Responder {
	return fn(params)
}

// GetWorkspaceByIDHandler interface for that can handle valid get workspace by Id params
type GetWorkspaceByIDHandler interface {
	Handle(GetWorkspaceByIDParams) middleware.Responder
}

// NewGetWorkspaceByID creates a new http.Handler for the get workspace by Id operation
func NewGetWorkspaceByID(ctx *middleware.Context, handler GetWorkspaceByIDHandler) *GetWorkspaceByID {
	return &GetWorkspaceByID{Context: ctx, Handler: handler}
}

/*GetWorkspaceByID swagger:route GET /workspaces/{id} getWorkspaceById

Returns a workspace

*/
type GetWorkspaceByID struct {
	Context *middleware.Context
	Handler GetWorkspaceByIDHandler
}

func (o *GetWorkspaceByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetWorkspaceByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}