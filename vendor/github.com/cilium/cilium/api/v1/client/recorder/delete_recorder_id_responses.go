// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package recorder

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// DeleteRecorderIDReader is a Reader for the DeleteRecorderID structure.
type DeleteRecorderIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRecorderIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRecorderIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteRecorderIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRecorderIDFailure()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRecorderIDOK creates a DeleteRecorderIDOK with default headers values
func NewDeleteRecorderIDOK() *DeleteRecorderIDOK {
	return &DeleteRecorderIDOK{}
}

/*
DeleteRecorderIDOK describes a response with status code 200, with default header values.

Success
*/
type DeleteRecorderIDOK struct {
}

// IsSuccess returns true when this delete recorder Id o k response has a 2xx status code
func (o *DeleteRecorderIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete recorder Id o k response has a 3xx status code
func (o *DeleteRecorderIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete recorder Id o k response has a 4xx status code
func (o *DeleteRecorderIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete recorder Id o k response has a 5xx status code
func (o *DeleteRecorderIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete recorder Id o k response a status code equal to that given
func (o *DeleteRecorderIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteRecorderIDOK) Error() string {
	return fmt.Sprintf("[DELETE /recorder/{id}][%d] deleteRecorderIdOK ", 200)
}

func (o *DeleteRecorderIDOK) String() string {
	return fmt.Sprintf("[DELETE /recorder/{id}][%d] deleteRecorderIdOK ", 200)
}

func (o *DeleteRecorderIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRecorderIDNotFound creates a DeleteRecorderIDNotFound with default headers values
func NewDeleteRecorderIDNotFound() *DeleteRecorderIDNotFound {
	return &DeleteRecorderIDNotFound{}
}

/*
DeleteRecorderIDNotFound describes a response with status code 404, with default header values.

Recorder not found
*/
type DeleteRecorderIDNotFound struct {
}

// IsSuccess returns true when this delete recorder Id not found response has a 2xx status code
func (o *DeleteRecorderIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete recorder Id not found response has a 3xx status code
func (o *DeleteRecorderIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete recorder Id not found response has a 4xx status code
func (o *DeleteRecorderIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete recorder Id not found response has a 5xx status code
func (o *DeleteRecorderIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete recorder Id not found response a status code equal to that given
func (o *DeleteRecorderIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteRecorderIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /recorder/{id}][%d] deleteRecorderIdNotFound ", 404)
}

func (o *DeleteRecorderIDNotFound) String() string {
	return fmt.Sprintf("[DELETE /recorder/{id}][%d] deleteRecorderIdNotFound ", 404)
}

func (o *DeleteRecorderIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRecorderIDFailure creates a DeleteRecorderIDFailure with default headers values
func NewDeleteRecorderIDFailure() *DeleteRecorderIDFailure {
	return &DeleteRecorderIDFailure{}
}

/*
DeleteRecorderIDFailure describes a response with status code 500, with default header values.

Recorder deletion failed
*/
type DeleteRecorderIDFailure struct {
	Payload models.Error
}

// IsSuccess returns true when this delete recorder Id failure response has a 2xx status code
func (o *DeleteRecorderIDFailure) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete recorder Id failure response has a 3xx status code
func (o *DeleteRecorderIDFailure) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete recorder Id failure response has a 4xx status code
func (o *DeleteRecorderIDFailure) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete recorder Id failure response has a 5xx status code
func (o *DeleteRecorderIDFailure) IsServerError() bool {
	return true
}

// IsCode returns true when this delete recorder Id failure response a status code equal to that given
func (o *DeleteRecorderIDFailure) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteRecorderIDFailure) Error() string {
	return fmt.Sprintf("[DELETE /recorder/{id}][%d] deleteRecorderIdFailure  %+v", 500, o.Payload)
}

func (o *DeleteRecorderIDFailure) String() string {
	return fmt.Sprintf("[DELETE /recorder/{id}][%d] deleteRecorderIdFailure  %+v", 500, o.Payload)
}

func (o *DeleteRecorderIDFailure) GetPayload() models.Error {
	return o.Payload
}

func (o *DeleteRecorderIDFailure) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
