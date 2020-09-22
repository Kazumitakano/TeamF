// Code generated by go-swagger; DO NOT EDIT.

package scenepicks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/shortintern2020-C-cryptograph/TeamF/server/gen/models"
)

// GetCommentByIDHandlerFunc turns a function with the right signature into a get comment by Id handler
type GetCommentByIDHandlerFunc func(GetCommentByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCommentByIDHandlerFunc) Handle(params GetCommentByIDParams) middleware.Responder {
	return fn(params)
}

// GetCommentByIDHandler interface for that can handle valid get comment by Id params
type GetCommentByIDHandler interface {
	Handle(GetCommentByIDParams) middleware.Responder
}

// NewGetCommentByID creates a new http.Handler for the get comment by Id operation
func NewGetCommentByID(ctx *middleware.Context, handler GetCommentByIDHandler) *GetCommentByID {
	return &GetCommentByID{Context: ctx, Handler: handler}
}

/*GetCommentByID swagger:route GET /dialog/{id} getCommentById

GetCommentByID get comment by Id API

*/
type GetCommentByID struct {
	Context *middleware.Context
	Handler GetCommentByIDHandler
}

func (o *GetCommentByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetCommentByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetCommentByIDOKBody get comment by ID o k body
//
// swagger:model GetCommentByIDOKBody
type GetCommentByIDOKBody struct {

	// comments
	Comments []*models.Comment `json:"comments"`

	// message
	Message string `json:"message,omitempty"`

	// tags
	Tags []*models.Tag `json:"tags"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (o *GetCommentByIDOKBody) UnmarshalJSON(data []byte) error {
	var props struct {

		// comments
		Comments []*models.Comment `json:"comments"`

		// message
		Message string `json:"message,omitempty"`

		// tags
		Tags []*models.Tag `json:"tags"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	o.Comments = props.Comments
	o.Message = props.Message
	o.Tags = props.Tags
	return nil
}

// Validate validates this get comment by ID o k body
func (o *GetCommentByIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComments(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCommentByIDOKBody) validateComments(formats strfmt.Registry) error {

	if swag.IsZero(o.Comments) { // not required
		return nil
	}

	for i := 0; i < len(o.Comments); i++ {
		if swag.IsZero(o.Comments[i]) { // not required
			continue
		}

		if o.Comments[i] != nil {
			if err := o.Comments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getCommentByIdOK" + "." + "comments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetCommentByIDOKBody) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(o.Tags) { // not required
		return nil
	}

	for i := 0; i < len(o.Tags); i++ {
		if swag.IsZero(o.Tags[i]) { // not required
			continue
		}

		if o.Tags[i] != nil {
			if err := o.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getCommentByIdOK" + "." + "tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCommentByIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCommentByIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetCommentByIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
