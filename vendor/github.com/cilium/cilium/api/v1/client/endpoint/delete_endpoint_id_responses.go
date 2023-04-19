// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// DeleteEndpointIDReader is a Reader for the DeleteEndpointID structure.
type DeleteEndpointIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEndpointIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEndpointIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 206:
		result := NewDeleteEndpointIDErrors()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteEndpointIDInvalid()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteEndpointIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteEndpointIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteEndpointIDOK creates a DeleteEndpointIDOK with default headers values
func NewDeleteEndpointIDOK() *DeleteEndpointIDOK {
	return &DeleteEndpointIDOK{}
}

/*
DeleteEndpointIDOK describes a response with status code 200, with default header values.

Success
*/
type DeleteEndpointIDOK struct {
}

// IsSuccess returns true when this delete endpoint Id o k response has a 2xx status code
func (o *DeleteEndpointIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete endpoint Id o k response has a 3xx status code
func (o *DeleteEndpointIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete endpoint Id o k response has a 4xx status code
func (o *DeleteEndpointIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete endpoint Id o k response has a 5xx status code
func (o *DeleteEndpointIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete endpoint Id o k response a status code equal to that given
func (o *DeleteEndpointIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteEndpointIDOK) Error() string {
	return fmt.Sprintf("[DELETE /endpoint/{id}][%d] deleteEndpointIdOK ", 200)
}

func (o *DeleteEndpointIDOK) String() string {
	return fmt.Sprintf("[DELETE /endpoint/{id}][%d] deleteEndpointIdOK ", 200)
}

func (o *DeleteEndpointIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEndpointIDErrors creates a DeleteEndpointIDErrors with default headers values
func NewDeleteEndpointIDErrors() *DeleteEndpointIDErrors {
	return &DeleteEndpointIDErrors{}
}

/*
DeleteEndpointIDErrors describes a response with status code 206, with default header values.

Deleted with a number of errors encountered
*/
type DeleteEndpointIDErrors struct {
	Payload int64
}

// IsSuccess returns true when this delete endpoint Id errors response has a 2xx status code
func (o *DeleteEndpointIDErrors) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete endpoint Id errors response has a 3xx status code
func (o *DeleteEndpointIDErrors) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete endpoint Id errors response has a 4xx status code
func (o *DeleteEndpointIDErrors) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete endpoint Id errors response has a 5xx status code
func (o *DeleteEndpointIDErrors) IsServerError() bool {
	return false
}

// IsCode returns true when this delete endpoint Id errors response a status code equal to that given
func (o *DeleteEndpointIDErrors) IsCode(code int) bool {
	return code == 206
}

func (o *DeleteEndpointIDErrors) Error() string {
	return fmt.Sprintf("[DELETE /endpoint/{id}][%d] deleteEndpointIdErrors  %+v", 206, o.Payload)
}

func (o *DeleteEndpointIDErrors) String() string {
	return fmt.Sprintf("[DELETE /endpoint/{id}][%d] deleteEndpointIdErrors  %+v", 206, o.Payload)
}

func (o *DeleteEndpointIDErrors) GetPayload() int64 {
	return o.Payload
}

func (o *DeleteEndpointIDErrors) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEndpointIDInvalid creates a DeleteEndpointIDInvalid with default headers values
func NewDeleteEndpointIDInvalid() *DeleteEndpointIDInvalid {
	return &DeleteEndpointIDInvalid{}
}

/*
	DeleteEndpointIDInvalid describes a response with status code 400, with default header values.

	Invalid endpoint ID format for specified type. Details in error

message
*/
type DeleteEndpointIDInvalid struct {
	Payload models.Error
}

// IsSuccess returns true when this delete endpoint Id invalid response has a 2xx status code
func (o *DeleteEndpointIDInvalid) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete endpoint Id invalid response has a 3xx status code
func (o *DeleteEndpointIDInvalid) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete endpoint Id invalid response has a 4xx status code
func (o *DeleteEndpointIDInvalid) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete endpoint Id invalid response has a 5xx status code
func (o *DeleteEndpointIDInvalid) IsServerError() bool {
	return false
}

// IsCode returns true when this delete endpoint Id invalid response a status code equal to that given
func (o *DeleteEndpointIDInvalid) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteEndpointIDInvalid) Error() string {
	return fmt.Sprintf("[DELETE /endpoint/{id}][%d] deleteEndpointIdInvalid  %+v", 400, o.Payload)
}

func (o *DeleteEndpointIDInvalid) String() string {
	return fmt.Sprintf("[DELETE /endpoint/{id}][%d] deleteEndpointIdInvalid  %+v", 400, o.Payload)
}

func (o *DeleteEndpointIDInvalid) GetPayload() models.Error {
	return o.Payload
}

func (o *DeleteEndpointIDInvalid) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEndpointIDNotFound creates a DeleteEndpointIDNotFound with default headers values
func NewDeleteEndpointIDNotFound() *DeleteEndpointIDNotFound {
	return &DeleteEndpointIDNotFound{}
}

/*
DeleteEndpointIDNotFound describes a response with status code 404, with default header values.

Endpoint not found
*/
type DeleteEndpointIDNotFound struct {
}

// IsSuccess returns true when this delete endpoint Id not found response has a 2xx status code
func (o *DeleteEndpointIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete endpoint Id not found response has a 3xx status code
func (o *DeleteEndpointIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete endpoint Id not found response has a 4xx status code
func (o *DeleteEndpointIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete endpoint Id not found response has a 5xx status code
func (o *DeleteEndpointIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete endpoint Id not found response a status code equal to that given
func (o *DeleteEndpointIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteEndpointIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /endpoint/{id}][%d] deleteEndpointIdNotFound ", 404)
}

func (o *DeleteEndpointIDNotFound) String() string {
	return fmt.Sprintf("[DELETE /endpoint/{id}][%d] deleteEndpointIdNotFound ", 404)
}

func (o *DeleteEndpointIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEndpointIDTooManyRequests creates a DeleteEndpointIDTooManyRequests with default headers values
func NewDeleteEndpointIDTooManyRequests() *DeleteEndpointIDTooManyRequests {
	return &DeleteEndpointIDTooManyRequests{}
}

/*
DeleteEndpointIDTooManyRequests describes a response with status code 429, with default header values.

Rate-limiting too many requests in the given time frame
*/
type DeleteEndpointIDTooManyRequests struct {
}

// IsSuccess returns true when this delete endpoint Id too many requests response has a 2xx status code
func (o *DeleteEndpointIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete endpoint Id too many requests response has a 3xx status code
func (o *DeleteEndpointIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete endpoint Id too many requests response has a 4xx status code
func (o *DeleteEndpointIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete endpoint Id too many requests response has a 5xx status code
func (o *DeleteEndpointIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete endpoint Id too many requests response a status code equal to that given
func (o *DeleteEndpointIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteEndpointIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /endpoint/{id}][%d] deleteEndpointIdTooManyRequests ", 429)
}

func (o *DeleteEndpointIDTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /endpoint/{id}][%d] deleteEndpointIdTooManyRequests ", 429)
}

func (o *DeleteEndpointIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
