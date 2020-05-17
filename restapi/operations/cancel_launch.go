// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CancelLaunchHandlerFunc turns a function with the right signature into a cancel launch handler
type CancelLaunchHandlerFunc func(CancelLaunchParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CancelLaunchHandlerFunc) Handle(params CancelLaunchParams) middleware.Responder {
	return fn(params)
}

// CancelLaunchHandler interface for that can handle valid cancel launch params
type CancelLaunchHandler interface {
	Handle(CancelLaunchParams) middleware.Responder
}

// NewCancelLaunch creates a new http.Handler for the cancel launch operation
func NewCancelLaunch(ctx *middleware.Context, handler CancelLaunchHandler) *CancelLaunch {
	return &CancelLaunch{Context: ctx, Handler: handler}
}

/*CancelLaunch swagger:route DELETE /workspaces/{id}/launch cancelLaunch

Stops a launched workspace

*/
type CancelLaunch struct {
	Context *middleware.Context
	Handler CancelLaunchHandler
}

func (o *CancelLaunch) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCancelLaunchParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
