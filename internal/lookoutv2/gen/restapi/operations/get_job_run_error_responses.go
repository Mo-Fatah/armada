// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/G-Research/armada/internal/lookoutv2/gen/models"
)

// GetJobRunErrorOKCode is the HTTP code returned for type GetJobRunErrorOK
const GetJobRunErrorOKCode int = 200

/*GetJobRunErrorOK Returns error for specific job run (if present)

swagger:response getJobRunErrorOK
*/
type GetJobRunErrorOK struct {

	/*
	  In: Body
	*/
	Payload *GetJobRunErrorOKBody `json:"body,omitempty"`
}

// NewGetJobRunErrorOK creates GetJobRunErrorOK with default headers values
func NewGetJobRunErrorOK() *GetJobRunErrorOK {

	return &GetJobRunErrorOK{}
}

// WithPayload adds the payload to the get job run error o k response
func (o *GetJobRunErrorOK) WithPayload(payload *GetJobRunErrorOKBody) *GetJobRunErrorOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get job run error o k response
func (o *GetJobRunErrorOK) SetPayload(payload *GetJobRunErrorOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetJobRunErrorOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetJobRunErrorBadRequestCode is the HTTP code returned for type GetJobRunErrorBadRequest
const GetJobRunErrorBadRequestCode int = 400

/*GetJobRunErrorBadRequest Error response

swagger:response getJobRunErrorBadRequest
*/
type GetJobRunErrorBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetJobRunErrorBadRequest creates GetJobRunErrorBadRequest with default headers values
func NewGetJobRunErrorBadRequest() *GetJobRunErrorBadRequest {

	return &GetJobRunErrorBadRequest{}
}

// WithPayload adds the payload to the get job run error bad request response
func (o *GetJobRunErrorBadRequest) WithPayload(payload *models.Error) *GetJobRunErrorBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get job run error bad request response
func (o *GetJobRunErrorBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetJobRunErrorBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetJobRunErrorDefault Error response

swagger:response getJobRunErrorDefault
*/
type GetJobRunErrorDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetJobRunErrorDefault creates GetJobRunErrorDefault with default headers values
func NewGetJobRunErrorDefault(code int) *GetJobRunErrorDefault {
	if code <= 0 {
		code = 500
	}

	return &GetJobRunErrorDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get job run error default response
func (o *GetJobRunErrorDefault) WithStatusCode(code int) *GetJobRunErrorDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get job run error default response
func (o *GetJobRunErrorDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get job run error default response
func (o *GetJobRunErrorDefault) WithPayload(payload *models.Error) *GetJobRunErrorDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get job run error default response
func (o *GetJobRunErrorDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetJobRunErrorDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}