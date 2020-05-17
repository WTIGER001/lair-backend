// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/wtiger001/lair-backend/models"
)

// LaunchWorkpaceByIDOKCode is the HTTP code returned for type LaunchWorkpaceByIDOK
const LaunchWorkpaceByIDOKCode int = 200

/*LaunchWorkpaceByIDOK OK

swagger:response launchWorkpaceByIdOK
*/
type LaunchWorkpaceByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.WorkspaceStatus `json:"body,omitempty"`
}

// NewLaunchWorkpaceByIDOK creates LaunchWorkpaceByIDOK with default headers values
func NewLaunchWorkpaceByIDOK() *LaunchWorkpaceByIDOK {

	return &LaunchWorkpaceByIDOK{}
}

// WithPayload adds the payload to the launch workpace by Id o k response
func (o *LaunchWorkpaceByIDOK) WithPayload(payload *models.WorkspaceStatus) *LaunchWorkpaceByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the launch workpace by Id o k response
func (o *LaunchWorkpaceByIDOK) SetPayload(payload *models.WorkspaceStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LaunchWorkpaceByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LaunchWorkpaceByIDBadRequestCode is the HTTP code returned for type LaunchWorkpaceByIDBadRequest
const LaunchWorkpaceByIDBadRequestCode int = 400

/*LaunchWorkpaceByIDBadRequest Invalid Parameters

swagger:response launchWorkpaceByIdBadRequest
*/
type LaunchWorkpaceByIDBadRequest struct {
}

// NewLaunchWorkpaceByIDBadRequest creates LaunchWorkpaceByIDBadRequest with default headers values
func NewLaunchWorkpaceByIDBadRequest() *LaunchWorkpaceByIDBadRequest {

	return &LaunchWorkpaceByIDBadRequest{}
}

// WriteResponse to the client
func (o *LaunchWorkpaceByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// LaunchWorkpaceByIDUnauthorizedCode is the HTTP code returned for type LaunchWorkpaceByIDUnauthorized
const LaunchWorkpaceByIDUnauthorizedCode int = 401

/*LaunchWorkpaceByIDUnauthorized Unauthenticated

swagger:response launchWorkpaceByIdUnauthorized
*/
type LaunchWorkpaceByIDUnauthorized struct {
}

// NewLaunchWorkpaceByIDUnauthorized creates LaunchWorkpaceByIDUnauthorized with default headers values
func NewLaunchWorkpaceByIDUnauthorized() *LaunchWorkpaceByIDUnauthorized {

	return &LaunchWorkpaceByIDUnauthorized{}
}

// WriteResponse to the client
func (o *LaunchWorkpaceByIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// LaunchWorkpaceByIDForbiddenCode is the HTTP code returned for type LaunchWorkpaceByIDForbidden
const LaunchWorkpaceByIDForbiddenCode int = 403

/*LaunchWorkpaceByIDForbidden User is not authorized to perform this action

swagger:response launchWorkpaceByIdForbidden
*/
type LaunchWorkpaceByIDForbidden struct {
}

// NewLaunchWorkpaceByIDForbidden creates LaunchWorkpaceByIDForbidden with default headers values
func NewLaunchWorkpaceByIDForbidden() *LaunchWorkpaceByIDForbidden {

	return &LaunchWorkpaceByIDForbidden{}
}

// WriteResponse to the client
func (o *LaunchWorkpaceByIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// LaunchWorkpaceByIDNotFoundCode is the HTTP code returned for type LaunchWorkpaceByIDNotFound
const LaunchWorkpaceByIDNotFoundCode int = 404

/*LaunchWorkpaceByIDNotFound Workspace not found

swagger:response launchWorkpaceByIdNotFound
*/
type LaunchWorkpaceByIDNotFound struct {
}

// NewLaunchWorkpaceByIDNotFound creates LaunchWorkpaceByIDNotFound with default headers values
func NewLaunchWorkpaceByIDNotFound() *LaunchWorkpaceByIDNotFound {

	return &LaunchWorkpaceByIDNotFound{}
}

// WriteResponse to the client
func (o *LaunchWorkpaceByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// LaunchWorkpaceByIDInternalServerErrorCode is the HTTP code returned for type LaunchWorkpaceByIDInternalServerError
const LaunchWorkpaceByIDInternalServerErrorCode int = 500

/*LaunchWorkpaceByIDInternalServerError Internal Server Error

swagger:response launchWorkpaceByIdInternalServerError
*/
type LaunchWorkpaceByIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewLaunchWorkpaceByIDInternalServerError creates LaunchWorkpaceByIDInternalServerError with default headers values
func NewLaunchWorkpaceByIDInternalServerError() *LaunchWorkpaceByIDInternalServerError {

	return &LaunchWorkpaceByIDInternalServerError{}
}

// WithPayload adds the payload to the launch workpace by Id internal server error response
func (o *LaunchWorkpaceByIDInternalServerError) WithPayload(payload string) *LaunchWorkpaceByIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the launch workpace by Id internal server error response
func (o *LaunchWorkpaceByIDInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LaunchWorkpaceByIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}