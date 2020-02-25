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

package ipam

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

// NewIPAMServicesListParams creates a new IPAMServicesListParams object
// with the default values initialized.
func NewIPAMServicesListParams() *IPAMServicesListParams {
	var ()
	return &IPAMServicesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMServicesListParamsWithTimeout creates a new IPAMServicesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMServicesListParamsWithTimeout(timeout time.Duration) *IPAMServicesListParams {
	var ()
	return &IPAMServicesListParams{

		timeout: timeout,
	}
}

// NewIPAMServicesListParamsWithContext creates a new IPAMServicesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMServicesListParamsWithContext(ctx context.Context) *IPAMServicesListParams {
	var ()
	return &IPAMServicesListParams{

		Context: ctx,
	}
}

// NewIPAMServicesListParamsWithHTTPClient creates a new IPAMServicesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMServicesListParamsWithHTTPClient(client *http.Client) *IPAMServicesListParams {
	var ()
	return &IPAMServicesListParams{
		HTTPClient: client,
	}
}

/*IPAMServicesListParams contains all the parameters to send to the API endpoint
for the ipam services list operation typically these are written to a http.Request
*/
type IPAMServicesListParams struct {

	/*Created*/
	Created *string
	/*CreatedGte*/
	CreatedGte *string
	/*CreatedLte*/
	CreatedLte *string
	/*Device*/
	Device *string
	/*DeviceID*/
	DeviceID *int64
	/*ID*/
	ID *int64
	/*LastUpdated*/
	LastUpdated *string
	/*LastUpdatedGte*/
	LastUpdatedGte *string
	/*LastUpdatedLte*/
	LastUpdatedLte *string
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
	/*Port*/
	Port *string
	/*Protocol*/
	Protocol *string
	/*Q*/
	Q *string
	/*Tag*/
	Tag *string
	/*VirtualMachine*/
	VirtualMachine *string
	/*VirtualMachineID*/
	VirtualMachineID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam services list params
func (o *IPAMServicesListParams) WithTimeout(timeout time.Duration) *IPAMServicesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam services list params
func (o *IPAMServicesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam services list params
func (o *IPAMServicesListParams) WithContext(ctx context.Context) *IPAMServicesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam services list params
func (o *IPAMServicesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam services list params
func (o *IPAMServicesListParams) WithHTTPClient(client *http.Client) *IPAMServicesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam services list params
func (o *IPAMServicesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the ipam services list params
func (o *IPAMServicesListParams) WithCreated(created *string) *IPAMServicesListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the ipam services list params
func (o *IPAMServicesListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the ipam services list params
func (o *IPAMServicesListParams) WithCreatedGte(createdGte *string) *IPAMServicesListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the ipam services list params
func (o *IPAMServicesListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the ipam services list params
func (o *IPAMServicesListParams) WithCreatedLte(createdLte *string) *IPAMServicesListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the ipam services list params
func (o *IPAMServicesListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithDevice adds the device to the ipam services list params
func (o *IPAMServicesListParams) WithDevice(device *string) *IPAMServicesListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the ipam services list params
func (o *IPAMServicesListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the ipam services list params
func (o *IPAMServicesListParams) WithDeviceID(deviceID *int64) *IPAMServicesListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the ipam services list params
func (o *IPAMServicesListParams) SetDeviceID(deviceID *int64) {
	o.DeviceID = deviceID
}

// WithID adds the id to the ipam services list params
func (o *IPAMServicesListParams) WithID(id *int64) *IPAMServicesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam services list params
func (o *IPAMServicesListParams) SetID(id *int64) {
	o.ID = id
}

// WithLastUpdated adds the lastUpdated to the ipam services list params
func (o *IPAMServicesListParams) WithLastUpdated(lastUpdated *string) *IPAMServicesListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the ipam services list params
func (o *IPAMServicesListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the ipam services list params
func (o *IPAMServicesListParams) WithLastUpdatedGte(lastUpdatedGte *string) *IPAMServicesListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the ipam services list params
func (o *IPAMServicesListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the ipam services list params
func (o *IPAMServicesListParams) WithLastUpdatedLte(lastUpdatedLte *string) *IPAMServicesListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the ipam services list params
func (o *IPAMServicesListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the ipam services list params
func (o *IPAMServicesListParams) WithLimit(limit *int64) *IPAMServicesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam services list params
func (o *IPAMServicesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the ipam services list params
func (o *IPAMServicesListParams) WithName(name *string) *IPAMServicesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ipam services list params
func (o *IPAMServicesListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the ipam services list params
func (o *IPAMServicesListParams) WithOffset(offset *int64) *IPAMServicesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam services list params
func (o *IPAMServicesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPort adds the port to the ipam services list params
func (o *IPAMServicesListParams) WithPort(port *string) *IPAMServicesListParams {
	o.SetPort(port)
	return o
}

// SetPort adds the port to the ipam services list params
func (o *IPAMServicesListParams) SetPort(port *string) {
	o.Port = port
}

// WithProtocol adds the protocol to the ipam services list params
func (o *IPAMServicesListParams) WithProtocol(protocol *string) *IPAMServicesListParams {
	o.SetProtocol(protocol)
	return o
}

// SetProtocol adds the protocol to the ipam services list params
func (o *IPAMServicesListParams) SetProtocol(protocol *string) {
	o.Protocol = protocol
}

// WithQ adds the q to the ipam services list params
func (o *IPAMServicesListParams) WithQ(q *string) *IPAMServicesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam services list params
func (o *IPAMServicesListParams) SetQ(q *string) {
	o.Q = q
}

// WithTag adds the tag to the ipam services list params
func (o *IPAMServicesListParams) WithTag(tag *string) *IPAMServicesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the ipam services list params
func (o *IPAMServicesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithVirtualMachine adds the virtualMachine to the ipam services list params
func (o *IPAMServicesListParams) WithVirtualMachine(virtualMachine *string) *IPAMServicesListParams {
	o.SetVirtualMachine(virtualMachine)
	return o
}

// SetVirtualMachine adds the virtualMachine to the ipam services list params
func (o *IPAMServicesListParams) SetVirtualMachine(virtualMachine *string) {
	o.VirtualMachine = virtualMachine
}

// WithVirtualMachineID adds the virtualMachineID to the ipam services list params
func (o *IPAMServicesListParams) WithVirtualMachineID(virtualMachineID *int64) *IPAMServicesListParams {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the ipam services list params
func (o *IPAMServicesListParams) SetVirtualMachineID(virtualMachineID *int64) {
	o.VirtualMachineID = virtualMachineID
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMServicesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Created != nil {

		// query param created
		var qrCreated string
		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {
			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}

	}

	if o.CreatedGte != nil {

		// query param created__gte
		var qrCreatedGte string
		if o.CreatedGte != nil {
			qrCreatedGte = *o.CreatedGte
		}
		qCreatedGte := qrCreatedGte
		if qCreatedGte != "" {
			if err := r.SetQueryParam("created__gte", qCreatedGte); err != nil {
				return err
			}
		}

	}

	if o.CreatedLte != nil {

		// query param created__lte
		var qrCreatedLte string
		if o.CreatedLte != nil {
			qrCreatedLte = *o.CreatedLte
		}
		qCreatedLte := qrCreatedLte
		if qCreatedLte != "" {
			if err := r.SetQueryParam("created__lte", qCreatedLte); err != nil {
				return err
			}
		}

	}

	if o.Device != nil {

		// query param device
		var qrDevice string
		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {
			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}

	}

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID int64
		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := swag.FormatInt64(qrDeviceID)
		if qDeviceID != "" {
			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
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

	if o.LastUpdated != nil {

		// query param last_updated
		var qrLastUpdated string
		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := qrLastUpdated
		if qLastUpdated != "" {
			if err := r.SetQueryParam("last_updated", qLastUpdated); err != nil {
				return err
			}
		}

	}

	if o.LastUpdatedGte != nil {

		// query param last_updated__gte
		var qrLastUpdatedGte string
		if o.LastUpdatedGte != nil {
			qrLastUpdatedGte = *o.LastUpdatedGte
		}
		qLastUpdatedGte := qrLastUpdatedGte
		if qLastUpdatedGte != "" {
			if err := r.SetQueryParam("last_updated__gte", qLastUpdatedGte); err != nil {
				return err
			}
		}

	}

	if o.LastUpdatedLte != nil {

		// query param last_updated__lte
		var qrLastUpdatedLte string
		if o.LastUpdatedLte != nil {
			qrLastUpdatedLte = *o.LastUpdatedLte
		}
		qLastUpdatedLte := qrLastUpdatedLte
		if qLastUpdatedLte != "" {
			if err := r.SetQueryParam("last_updated__lte", qLastUpdatedLte); err != nil {
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

	if o.Port != nil {

		// query param port
		var qrPort string
		if o.Port != nil {
			qrPort = *o.Port
		}
		qPort := qrPort
		if qPort != "" {
			if err := r.SetQueryParam("port", qPort); err != nil {
				return err
			}
		}

	}

	if o.Protocol != nil {

		// query param protocol
		var qrProtocol string
		if o.Protocol != nil {
			qrProtocol = *o.Protocol
		}
		qProtocol := qrProtocol
		if qProtocol != "" {
			if err := r.SetQueryParam("protocol", qProtocol); err != nil {
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

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
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
