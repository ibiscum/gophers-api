// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetGopherRandomHandlerFunc turns a function with the right signature into a get gopher random handler
type GetGopherRandomHandlerFunc func(GetGopherRandomParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetGopherRandomHandlerFunc) Handle(params GetGopherRandomParams) middleware.Responder {
	return fn(params)
}

// GetGopherRandomHandler interface for that can handle valid get gopher random params
type GetGopherRandomHandler interface {
	Handle(GetGopherRandomParams) middleware.Responder
}

// NewGetGopherRandom creates a new http.Handler for the get gopher random operation
func NewGetGopherRandom(ctx *middleware.Context, handler GetGopherRandomHandler) *GetGopherRandom {
	return &GetGopherRandom{Context: ctx, Handler: handler}
}

/*
	GetGopherRandom swagger:route GET /gopher/random getGopherRandom

Return a random Gopher Image
*/
type GetGopherRandom struct {
	Context *middleware.Context
	Handler GetGopherRandomHandler
}

func (o *GetGopherRandom) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetGopherRandomParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}