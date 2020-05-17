// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/wtiger001/lair-backend/models"
)

// PutWorkpaceByIDOKCode is the HTTP code returned for type PutWorkpaceByIDOK
const PutWorkpaceByIDOKCode int = 200

/*PutWorkpaceByIDOK OK

swagger:response putWorkpaceByIdOK
*/
type PutWorkpaceByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Workspace `json:"body,omitempty"`
}

// NewPutWorkpaceByIDOK creates PutWorkpaceByIDOK with default headers values
func NewPutWorkpaceByIDOK() *PutWorkpaceByIDOK {

	return &PutWorkpaceByIDOK{}
}

// WithPayload adds the payload to the put workpace by Id o k response
func (o *PutWorkpaceByIDOK) WithPayload(payload *models.Workspace) *PutWorkpaceByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put workpace by Id o k response
func (o *PutWorkpaceByIDOK) SetPayload(payload *models.Workspace) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutWorkpaceByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutWorkpaceByIDBadRequestCode is the HTTP code returned for type PutWorkpaceByIDBadRequest
const PutWorkpaceByIDBadRequestCode int = 400

/*PutWorkpaceByIDBadRequest Invalid Parameters

swagger:response putWorkpaceByIdBadRequest
*/
type PutWorkpaceByIDBadRequest struct {
}

// NewPutWorkpaceByIDBadRequest creates PutWorkpaceByIDBadRequest with default headers values
func NewPutWorkpaceByIDBadRequest() *PutWorkpaceByIDBadRequest {

	return &PutWorkpaceByIDBadRequest{}
}

// WriteResponse to the client
func (o *PutWorkpaceByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PutWorkpaceByIDUnauthorizedCode is the HTTP code returned for type PutWorkpaceByIDUnauthorized
const PutWorkpaceByIDUnauthorizedCode int = 401

/*PutWorkpaceByIDUnauthorized Unauthenticated

swagger:response putWorkpaceByIdUnauthorized
*/
type PutWorkpaceByIDUnauthorized struct {
}

// NewPutWorkpaceByIDUnauthorized creates PutWorkpaceByIDUnauthorized with default headers values
func NewPutWorkpaceByIDUnauthorized() *PutWorkpaceByIDUnauthorized {

	return &PutWorkpaceByIDUnauthorized{}
}

// WriteResponse to the client
func (o *PutWorkpaceByIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// PutWorkpaceByIDForbiddenCode is the HTTP code returned for type PutWorkpaceByIDForbidden
const PutWorkpaceByIDForbiddenCode int = 403

/*PutWorkpaceByIDForbidden User is not authorized to perform this action

swagger:response putWorkpaceByIdForbidden
*/
type PutWorkpaceByIDForbidden struct {
}

// NewPutWorkpaceByIDForbidden creates PutWorkpaceByIDForbidden with default headers values
func NewPutWorkpaceByIDForbidden() *PutWorkpaceByIDForbidden {

	return &PutWorkpaceByIDForbidden{}
}

// WriteResponse to the client
func (o *PutWorkpaceByIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// PutWorkpaceByIDNotFoundCode is the HTTP code returned for type PutWorkpaceByIDNotFound
const PutWorkpaceByIDNotFoundCode int = 404

/*PutWorkpaceByIDNotFound Workspace not found

swagger:response putWorkpaceByIdNotFound
*/
type PutWorkpaceByIDNotFound struct {
}

// NewPutWorkpaceByIDNotFound creates PutWorkpaceByIDNotFound with default headers values
func NewPutWorkpaceByIDNotFound() *PutWorkpaceByIDNotFound {

	return &PutWorkpaceByIDNotFound{}
}

// WriteResponse to the client
func (o *PutWorkpaceByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// PutWorkpaceByIDInternalServerErrorCode is the HTTP code returned for type PutWorkpaceByIDInternalServerError
const PutWorkpaceByIDInternalServerErrorCode int = 500

/*PutWorkpaceByIDInternalServerError Internal Server Error

swagger:response putWorkpaceByIdInternalServerError
*/
type PutWorkpaceByIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPutWorkpaceByIDInternalServerError creates PutWorkpaceByIDInternalServerError with default headers values
func NewPutWorkpaceByIDInternalServerError() *PutWorkpaceByIDInternalServerError {

	return &PutWorkpaceByIDInternalServerError{}
}

// WithPayload adds the payload to the put workpace by Id internal server error response
func (o *PutWorkpaceByIDInternalServerError) WithPayload(payload string) *PutWorkpaceByIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put workpace by Id internal server error response
func (o *PutWorkpaceByIDInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutWorkpaceByIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
