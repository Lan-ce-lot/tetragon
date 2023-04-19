// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetMapNameEventsParams creates a new GetMapNameEventsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMapNameEventsParams() *GetMapNameEventsParams {
	return &GetMapNameEventsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMapNameEventsParamsWithTimeout creates a new GetMapNameEventsParams object
// with the ability to set a timeout on a request.
func NewGetMapNameEventsParamsWithTimeout(timeout time.Duration) *GetMapNameEventsParams {
	return &GetMapNameEventsParams{
		timeout: timeout,
	}
}

// NewGetMapNameEventsParamsWithContext creates a new GetMapNameEventsParams object
// with the ability to set a context for a request.
func NewGetMapNameEventsParamsWithContext(ctx context.Context) *GetMapNameEventsParams {
	return &GetMapNameEventsParams{
		Context: ctx,
	}
}

// NewGetMapNameEventsParamsWithHTTPClient creates a new GetMapNameEventsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMapNameEventsParamsWithHTTPClient(client *http.Client) *GetMapNameEventsParams {
	return &GetMapNameEventsParams{
		HTTPClient: client,
	}
}

/*
GetMapNameEventsParams contains all the parameters to send to the API endpoint

	for the get map name events operation.

	Typically these are written to a http.Request.
*/
type GetMapNameEventsParams struct {

	/* Follow.

	   Whether to follow streamed requests
	*/
	Follow *bool

	/* Name.

	   Name of map
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get map name events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMapNameEventsParams) WithDefaults() *GetMapNameEventsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get map name events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMapNameEventsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get map name events params
func (o *GetMapNameEventsParams) WithTimeout(timeout time.Duration) *GetMapNameEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get map name events params
func (o *GetMapNameEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get map name events params
func (o *GetMapNameEventsParams) WithContext(ctx context.Context) *GetMapNameEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get map name events params
func (o *GetMapNameEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get map name events params
func (o *GetMapNameEventsParams) WithHTTPClient(client *http.Client) *GetMapNameEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get map name events params
func (o *GetMapNameEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFollow adds the follow to the get map name events params
func (o *GetMapNameEventsParams) WithFollow(follow *bool) *GetMapNameEventsParams {
	o.SetFollow(follow)
	return o
}

// SetFollow adds the follow to the get map name events params
func (o *GetMapNameEventsParams) SetFollow(follow *bool) {
	o.Follow = follow
}

// WithName adds the name to the get map name events params
func (o *GetMapNameEventsParams) WithName(name string) *GetMapNameEventsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get map name events params
func (o *GetMapNameEventsParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetMapNameEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Follow != nil {

		// query param follow
		var qrFollow bool

		if o.Follow != nil {
			qrFollow = *o.Follow
		}
		qFollow := swag.FormatBool(qrFollow)
		if qFollow != "" {

			if err := r.SetQueryParam("follow", qFollow); err != nil {
				return err
			}
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
