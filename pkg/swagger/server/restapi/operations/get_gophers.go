// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetGophersHandlerFunc turns a function with the right signature into a get gophers handler
type GetGophersHandlerFunc func(GetGophersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetGophersHandlerFunc) Handle(params GetGophersParams) middleware.Responder {
	return fn(params)
}

// GetGophersHandler interface for that can handle valid get gophers params
type GetGophersHandler interface {
	Handle(GetGophersParams) middleware.Responder
}

// NewGetGophers creates a new http.Handler for the get gophers operation
func NewGetGophers(ctx *middleware.Context, handler GetGophersHandler) *GetGophers {
	return &GetGophers{Context: ctx, Handler: handler}
}

/*
	GetGophers swagger:route GET /gophers getGophers

List Gophers
*/
type GetGophers struct {
	Context *middleware.Context
	Handler GetGophersHandler
}

func (o *GetGophers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetGophersParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
