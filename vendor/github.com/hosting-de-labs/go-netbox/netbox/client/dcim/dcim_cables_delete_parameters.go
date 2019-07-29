// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDcimCablesDeleteParams creates a new DcimCablesDeleteParams object
// with the default values initialized.
func NewDcimCablesDeleteParams() *DcimCablesDeleteParams {
	var ()
	return &DcimCablesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimCablesDeleteParamsWithTimeout creates a new DcimCablesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimCablesDeleteParamsWithTimeout(timeout time.Duration) *DcimCablesDeleteParams {
	var ()
	return &DcimCablesDeleteParams{

		timeout: timeout,
	}
}

// NewDcimCablesDeleteParamsWithContext creates a new DcimCablesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimCablesDeleteParamsWithContext(ctx context.Context) *DcimCablesDeleteParams {
	var ()
	return &DcimCablesDeleteParams{

		Context: ctx,
	}
}

// NewDcimCablesDeleteParamsWithHTTPClient creates a new DcimCablesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimCablesDeleteParamsWithHTTPClient(client *http.Client) *DcimCablesDeleteParams {
	var ()
	return &DcimCablesDeleteParams{
		HTTPClient: client,
	}
}

/*DcimCablesDeleteParams contains all the parameters to send to the API endpoint
for the dcim cables delete operation typically these are written to a http.Request
*/
type DcimCablesDeleteParams struct {

	/*ID
	  A unique integer value identifying this cable.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim cables delete params
func (o *DcimCablesDeleteParams) WithTimeout(timeout time.Duration) *DcimCablesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim cables delete params
func (o *DcimCablesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim cables delete params
func (o *DcimCablesDeleteParams) WithContext(ctx context.Context) *DcimCablesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim cables delete params
func (o *DcimCablesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim cables delete params
func (o *DcimCablesDeleteParams) WithHTTPClient(client *http.Client) *DcimCablesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim cables delete params
func (o *DcimCablesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim cables delete params
func (o *DcimCablesDeleteParams) WithID(id int64) *DcimCablesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim cables delete params
func (o *DcimCablesDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimCablesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}