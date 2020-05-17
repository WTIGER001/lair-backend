// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/wtiger001/lair-backend/models"
)

// GetWorkpaceLaunchStatusOKCode is the HTTP code returned for type GetWorkpaceLaunchStatusOK
const GetWorkpaceLaunchStatusOKCode int = 200

/*GetWorkpaceLaunchStatusOK OK

swagger:response getWorkpaceLaunchStatusOK
*/
type GetWorkpaceLaunchStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.WorkspaceStatus `json:"body,omitempty"`
}

// NewGetWorkpaceLaunchStatusOK creates GetWorkpaceLaunchStatusOK with default headers values
func NewGetWorkpaceLaunchStatusOK() *GetWorkpaceLaunchStatusOK {

	return &GetWorkpaceLaunchStatusOK{}
}

// WithPayload adds the payload to the get workpace launch status o k response
func (o *GetWorkpaceLaunchStatusOK) WithPayload(payload *models.WorkspaceStatus) *GetWorkpaceLaunchStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get workpace launch status o k response
func (o *GetWorkpaceLaunchStatusOK) SetPayload(payload *models.WorkspaceStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWorkpaceLaunchStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetWorkpaceLaunchStatusBadRequestCode is the HTTP code returned for type GetWorkpaceLaunchStatusBadRequest
const GetWorkpaceLaunchStatusBadRequestCode int = 400

/*GetWorkpaceLaunchStatusBadRequest Invalid Parameters

swagger:response getWorkpaceLaunchStatusBadRequest
*/
type GetWorkpaceLaunchStatusBadRequest struct {
}

// NewGetWorkpaceLaunchStatusBadRequest creates GetWorkpaceLaunchStatusBadRequest with default headers values
func NewGetWorkpaceLaunchStatusBadRequest() *GetWorkpaceLaunchStatusBadRequest {

	return &GetWorkpaceLaunchStatusBadRequest{}
}

// WriteResponse to the client
func (o *GetWorkpaceLaunchStatusBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetWorkpaceLaunchStatusUnauthorizedCode is the HTTP code returned for type GetWorkpaceLaunchStatusUnauthorized
const GetWorkpaceLaunchStatusUnauthorizedCode int = 401

/*GetWorkpaceLaunchStatusUnauthorized Unauthenticated

swagger:response getWorkpaceLaunchStatusUnauthorized
*/
type GetWorkpaceLaunchStatusUnauthorized struct {
}

// NewGetWorkpaceLaunchStatusUnauthorized creates GetWorkpaceLaunchStatusUnauthorized with default headers values
func NewGetWorkpaceLaunchStatusUnauthorized() *GetWorkpaceLaunchStatusUnauthorized {

	return &GetWorkpaceLaunchStatusUnauthorized{}
}

// WriteResponse to the client
func (o *GetWorkpaceLaunchStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetWorkpaceLaunchStatusForbiddenCode is the HTTP code returned for type GetWorkpaceLaunchStatusForbidden
const GetWorkpaceLaunchStatusForbiddenCode int = 403

/*GetWorkpaceLaunchStatusForbidden User is not authorized to perform this action

swagger:response getWorkpaceLaunchStatusForbidden
*/
type GetWorkpaceLaunchStatusForbidden struct {
}

// NewGetWorkpaceLaunchStatusForbidden creates GetWorkpaceLaunchStatusForbidden with default headers values
func NewGetWorkpaceLaunchStatusForbidden() *GetWorkpaceLaunchStatusForbidden {

	return &GetWorkpaceLaunchStatusForbidden{}
}

// WriteResponse to the client
func (o *GetWorkpaceLaunchStatusForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetWorkpaceLaunchStatusNotFoundCode is the HTTP code returned for type GetWorkpaceLaunchStatusNotFound
const GetWorkpaceLaunchStatusNotFoundCode int = 404

/*GetWorkpaceLaunchStatusNotFound Workspace not found

swagger:response getWorkpaceLaunchStatusNotFound
*/
type GetWorkpaceLaunchStatusNotFound struct {
}

// NewGetWorkpaceLaunchStatusNotFound creates GetWorkpaceLaunchStatusNotFound with default headers values
func NewGetWorkpaceLaunchStatusNotFound() *GetWorkpaceLaunchStatusNotFound {

	return &GetWorkpaceLaunchStatusNotFound{}
}

// WriteResponse to the client
func (o *GetWorkpaceLaunchStatusNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetWorkpaceLaunchStatusInternalServerErrorCode is the HTTP code returned for type GetWorkpaceLaunchStatusInternalServerError
const GetWorkpaceLaunchStatusInternalServerErrorCode int = 500

/*GetWorkpaceLaunchStatusInternalServerError Internal Server Error

swagger:response getWorkpaceLaunchStatusInternalServerError
*/
type GetWorkpaceLaunchStatusInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetWorkpaceLaunchStatusInternalServerError creates GetWorkpaceLaunchStatusInternalServerError with default headers values
func NewGetWorkpaceLaunchStatusInternalServerError() *GetWorkpaceLaunchStatusInternalServerError {

	return &GetWorkpaceLaunchStatusInternalServerError{}
}

// WithPayload adds the payload to the get workpace launch status internal server error response
func (o *GetWorkpaceLaunchStatusInternalServerError) WithPayload(payload string) *GetWorkpaceLaunchStatusInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get workpace launch status internal server error response
func (o *GetWorkpaceLaunchStatusInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWorkpaceLaunchStatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}