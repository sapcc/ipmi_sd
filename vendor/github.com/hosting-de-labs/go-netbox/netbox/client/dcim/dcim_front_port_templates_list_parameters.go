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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDcimFrontPortTemplatesListParams creates a new DcimFrontPortTemplatesListParams object
// with the default values initialized.
func NewDcimFrontPortTemplatesListParams() *DcimFrontPortTemplatesListParams {
	var ()
	return &DcimFrontPortTemplatesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortTemplatesListParamsWithTimeout creates a new DcimFrontPortTemplatesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimFrontPortTemplatesListParamsWithTimeout(timeout time.Duration) *DcimFrontPortTemplatesListParams {
	var ()
	return &DcimFrontPortTemplatesListParams{

		timeout: timeout,
	}
}

// NewDcimFrontPortTemplatesListParamsWithContext creates a new DcimFrontPortTemplatesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimFrontPortTemplatesListParamsWithContext(ctx context.Context) *DcimFrontPortTemplatesListParams {
	var ()
	return &DcimFrontPortTemplatesListParams{

		Context: ctx,
	}
}

// NewDcimFrontPortTemplatesListParamsWithHTTPClient creates a new DcimFrontPortTemplatesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimFrontPortTemplatesListParamsWithHTTPClient(client *http.Client) *DcimFrontPortTemplatesListParams {
	var ()
	return &DcimFrontPortTemplatesListParams{
		HTTPClient: client,
	}
}

/*DcimFrontPortTemplatesListParams contains all the parameters to send to the API endpoint
for the dcim front port templates list operation typically these are written to a http.Request
*/
type DcimFrontPortTemplatesListParams struct {

	/*DevicetypeID*/
	DevicetypeID *int64
	/*ID*/
	ID *int64
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*Type*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithTimeout(timeout time.Duration) *DcimFrontPortTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithContext(ctx context.Context) *DcimFrontPortTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithHTTPClient(client *http.Client) *DcimFrontPortTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevicetypeID adds the devicetypeID to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithDevicetypeID(devicetypeID *int64) *DcimFrontPortTemplatesListParams {
	o.SetDevicetypeID(devicetypeID)
	return o
}

// SetDevicetypeID adds the devicetypeId to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetDevicetypeID(devicetypeID *int64) {
	o.DevicetypeID = devicetypeID
}

// WithID adds the id to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithID(id *int64) *DcimFrontPortTemplatesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetID(id *int64) {
	o.ID = id
}

// WithLimit adds the limit to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithLimit(limit *int64) *DcimFrontPortTemplatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithName(name *string) *DcimFrontPortTemplatesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithOffset(offset *int64) *DcimFrontPortTemplatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithQ(q *string) *DcimFrontPortTemplatesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetQ(q *string) {
	o.Q = q
}

// WithType adds the typeVar to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithType(typeVar *string) *DcimFrontPortTemplatesListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DevicetypeID != nil {

		// query param devicetype_id
		var qrDevicetypeID int64
		if o.DevicetypeID != nil {
			qrDevicetypeID = *o.DevicetypeID
		}
		qDevicetypeID := swag.FormatInt64(qrDevicetypeID)
		if qDevicetypeID != "" {
			if err := r.SetQueryParam("devicetype_id", qDevicetypeID); err != nil {
				return err
			}
		}

	}

	if o.ID != nil {

		// query param id
		var qrID int64
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := swag.FormatInt64(qrID)
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
