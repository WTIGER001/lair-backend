// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteWorkpaceByIDHandlerFunc turns a function with the right signature into a delete workpace by Id handler
type DeleteWorkpaceByIDHandlerFunc func(DeleteWorkpaceByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteWorkpaceByIDHandlerFunc) Handle(params DeleteWorkpaceByIDParams) middleware.Responder {
	return fn(params)
}

// DeleteWorkpaceByIDHandler interface for that can handle valid delete workpace by Id params
type DeleteWorkpaceByIDHandler interface {
	Handle(DeleteWorkpaceByIDParams) middleware.Responder
}

// NewDeleteWorkpaceByID creates a new http.Handler for the delete workpace by Id operation
func NewDeleteWorkpaceByID(ctx *middleware.Context, handler DeleteWorkpaceByIDHandler) *DeleteWorkpaceByID {
	return &DeleteWorkpaceByID{Context: ctx, Handler: handler}
}

/*DeleteWorkpaceByID swagger:route DELETE /workspaces/{id} deleteWorkpaceById

Deletes a workspace

*/
type DeleteWorkpaceByID struct {
	Context *middleware.Context
	Handler DeleteWorkpaceByIDHandler
}

func (o *DeleteWorkpaceByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteWorkpaceByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}