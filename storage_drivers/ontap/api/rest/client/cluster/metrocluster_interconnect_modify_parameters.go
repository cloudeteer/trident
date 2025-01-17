// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewMetroclusterInterconnectModifyParams creates a new MetroclusterInterconnectModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMetroclusterInterconnectModifyParams() *MetroclusterInterconnectModifyParams {
	return &MetroclusterInterconnectModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMetroclusterInterconnectModifyParamsWithTimeout creates a new MetroclusterInterconnectModifyParams object
// with the ability to set a timeout on a request.
func NewMetroclusterInterconnectModifyParamsWithTimeout(timeout time.Duration) *MetroclusterInterconnectModifyParams {
	return &MetroclusterInterconnectModifyParams{
		timeout: timeout,
	}
}

// NewMetroclusterInterconnectModifyParamsWithContext creates a new MetroclusterInterconnectModifyParams object
// with the ability to set a context for a request.
func NewMetroclusterInterconnectModifyParamsWithContext(ctx context.Context) *MetroclusterInterconnectModifyParams {
	return &MetroclusterInterconnectModifyParams{
		Context: ctx,
	}
}

// NewMetroclusterInterconnectModifyParamsWithHTTPClient creates a new MetroclusterInterconnectModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewMetroclusterInterconnectModifyParamsWithHTTPClient(client *http.Client) *MetroclusterInterconnectModifyParams {
	return &MetroclusterInterconnectModifyParams{
		HTTPClient: client,
	}
}

/* MetroclusterInterconnectModifyParams contains all the parameters to send to the API endpoint
   for the metrocluster interconnect modify operation.

   Typically these are written to a http.Request.
*/
type MetroclusterInterconnectModifyParams struct {

	/* Adapter.

	   Interconnect adapter
	*/
	AdapterPathParameter string

	/* Info.

	   MetroCluster interconnect interface information
	*/
	Info *models.MetroclusterInterconnect

	/* NodeUUID.

	   Node UUID
	*/
	NodeUUIDPathParameter string

	/* PartnerType.

	   DR Partner type
	*/
	PartnerTypePathParameter string

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeoutQueryParameter *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the metrocluster interconnect modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetroclusterInterconnectModifyParams) WithDefaults() *MetroclusterInterconnectModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the metrocluster interconnect modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetroclusterInterconnectModifyParams) SetDefaults() {
	var (
		returnTimeoutQueryParameterDefault = int64(0)
	)

	val := MetroclusterInterconnectModifyParams{
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) WithTimeout(timeout time.Duration) *MetroclusterInterconnectModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) WithContext(ctx context.Context) *MetroclusterInterconnectModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) WithHTTPClient(client *http.Client) *MetroclusterInterconnectModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdapterPathParameter adds the adapter to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) WithAdapterPathParameter(adapter string) *MetroclusterInterconnectModifyParams {
	o.SetAdapterPathParameter(adapter)
	return o
}

// SetAdapterPathParameter adds the adapter to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) SetAdapterPathParameter(adapter string) {
	o.AdapterPathParameter = adapter
}

// WithInfo adds the info to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) WithInfo(info *models.MetroclusterInterconnect) *MetroclusterInterconnectModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) SetInfo(info *models.MetroclusterInterconnect) {
	o.Info = info
}

// WithNodeUUIDPathParameter adds the nodeUUID to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) WithNodeUUIDPathParameter(nodeUUID string) *MetroclusterInterconnectModifyParams {
	o.SetNodeUUIDPathParameter(nodeUUID)
	return o
}

// SetNodeUUIDPathParameter adds the nodeUuid to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) SetNodeUUIDPathParameter(nodeUUID string) {
	o.NodeUUIDPathParameter = nodeUUID
}

// WithPartnerTypePathParameter adds the partnerType to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) WithPartnerTypePathParameter(partnerType string) *MetroclusterInterconnectModifyParams {
	o.SetPartnerTypePathParameter(partnerType)
	return o
}

// SetPartnerTypePathParameter adds the partnerType to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) SetPartnerTypePathParameter(partnerType string) {
	o.PartnerTypePathParameter = partnerType
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *MetroclusterInterconnectModifyParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *MetroclusterInterconnectModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param adapter
	if err := r.SetPathParam("adapter", o.AdapterPathParameter); err != nil {
		return err
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param node.uuid
	if err := r.SetPathParam("node.uuid", o.NodeUUIDPathParameter); err != nil {
		return err
	}

	// path param partner_type
	if err := r.SetPathParam("partner_type", o.PartnerTypePathParameter); err != nil {
		return err
	}

	if o.ReturnTimeoutQueryParameter != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeoutQueryParameter != nil {
			qrReturnTimeout = *o.ReturnTimeoutQueryParameter
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
