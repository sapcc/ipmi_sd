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

package virtualization

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

// NewVirtualizationInterfacesListParams creates a new VirtualizationInterfacesListParams object
// with the default values initialized.
func NewVirtualizationInterfacesListParams() *VirtualizationInterfacesListParams {
	var ()
	return &VirtualizationInterfacesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationInterfacesListParamsWithTimeout creates a new VirtualizationInterfacesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationInterfacesListParamsWithTimeout(timeout time.Duration) *VirtualizationInterfacesListParams {
	var ()
	return &VirtualizationInterfacesListParams{

		timeout: timeout,
	}
}

// NewVirtualizationInterfacesListParamsWithContext creates a new VirtualizationInterfacesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationInterfacesListParamsWithContext(ctx context.Context) *VirtualizationInterfacesListParams {
	var ()
	return &VirtualizationInterfacesListParams{

		Context: ctx,
	}
}

// NewVirtualizationInterfacesListParamsWithHTTPClient creates a new VirtualizationInterfacesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationInterfacesListParamsWithHTTPClient(client *http.Client) *VirtualizationInterfacesListParams {
	var ()
	return &VirtualizationInterfacesListParams{
		HTTPClient: client,
	}
}

/*VirtualizationInterfacesListParams contains all the parameters to send to the API endpoint
for the virtualization interfaces list operation typically these are written to a http.Request
*/
type VirtualizationInterfacesListParams struct {

	/*Enabled*/
	Enabled *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*MacAddress*/
	MacAddress *string
	/*Mtu*/
	Mtu *int64
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*VirtualMachine*/
	VirtualMachine *string
	/*VirtualMachineID*/
	VirtualMachineID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithTimeout(timeout time.Duration) *VirtualizationInterfacesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithContext(ctx context.Context) *VirtualizationInterfacesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithHTTPClient(client *http.Client) *VirtualizationInterfacesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnabled adds the enabled to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithEnabled(enabled *string) *VirtualizationInterfacesListParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetEnabled(enabled *string) {
	o.Enabled = enabled
}

// WithLimit adds the limit to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithLimit(limit *int64) *VirtualizationInterfacesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMacAddress adds the macAddress to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithMacAddress(macAddress *string) *VirtualizationInterfacesListParams {
	o.SetMacAddress(macAddress)
	return o
}

// SetMacAddress adds the macAddress to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetMacAddress(macAddress *string) {
	o.MacAddress = macAddress
}

// WithMtu adds the mtu to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithMtu(mtu *int64) *VirtualizationInterfacesListParams {
	o.SetMtu(mtu)
	return o
}

// SetMtu adds the mtu to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetMtu(mtu *int64) {
	o.Mtu = mtu
}

// WithName adds the name to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithName(name *string) *VirtualizationInterfacesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithOffset(offset *int64) *VirtualizationInterfacesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithQ(q *string) *VirtualizationInterfacesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetQ(q *string) {
	o.Q = q
}

// WithVirtualMachine adds the virtualMachine to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithVirtualMachine(virtualMachine *string) *VirtualizationInterfacesListParams {
	o.SetVirtualMachine(virtualMachine)
	return o
}

// SetVirtualMachine adds the virtualMachine to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetVirtualMachine(virtualMachine *string) {
	o.VirtualMachine = virtualMachine
}

// WithVirtualMachineID adds the virtualMachineID to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithVirtualMachineID(virtualMachineID *int64) *VirtualizationInterfacesListParams {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetVirtualMachineID(virtualMachineID *int64) {
	o.VirtualMachineID = virtualMachineID
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationInterfacesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Enabled != nil {

		// query param enabled
		var qrEnabled string
		if o.Enabled != nil {
			qrEnabled = *o.Enabled
		}
		qEnabled := qrEnabled
		if qEnabled != "" {
			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
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

	if o.MacAddress != nil {

		// query param mac_address
		var qrMacAddress string
		if o.MacAddress != nil {
			qrMacAddress = *o.MacAddress
		}
		qMacAddress := qrMacAddress
		if qMacAddress != "" {
			if err := r.SetQueryParam("mac_address", qMacAddress); err != nil {
				return err
			}
		}

	}

	if o.Mtu != nil {

		// query param mtu
		var qrMtu int64
		if o.Mtu != nil {
			qrMtu = *o.Mtu
		}
		qMtu := swag.FormatInt64(qrMtu)
		if qMtu != "" {
			if err := r.SetQueryParam("mtu", qMtu); err != nil {
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

	if o.VirtualMachine != nil {

		// query param virtual_machine
		var qrVirtualMachine string
		if o.VirtualMachine != nil {
			qrVirtualMachine = *o.VirtualMachine
		}
		qVirtualMachine := qrVirtualMachine
		if qVirtualMachine != "" {
			if err := r.SetQueryParam("virtual_machine", qVirtualMachine); err != nil {
				return err
			}
		}

	}

	if o.VirtualMachineID != nil {

		// query param virtual_machine_id
		var qrVirtualMachineID int64
		if o.VirtualMachineID != nil {
			qrVirtualMachineID = *o.VirtualMachineID
		}
		qVirtualMachineID := swag.FormatInt64(qrVirtualMachineID)
		if qVirtualMachineID != "" {
			if err := r.SetQueryParam("virtual_machine_id", qVirtualMachineID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
