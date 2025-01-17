// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// NewIgroupNestedCreateParams creates a new IgroupNestedCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIgroupNestedCreateParams() *IgroupNestedCreateParams {
	return &IgroupNestedCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIgroupNestedCreateParamsWithTimeout creates a new IgroupNestedCreateParams object
// with the ability to set a timeout on a request.
func NewIgroupNestedCreateParamsWithTimeout(timeout time.Duration) *IgroupNestedCreateParams {
	return &IgroupNestedCreateParams{
		timeout: timeout,
	}
}

// NewIgroupNestedCreateParamsWithContext creates a new IgroupNestedCreateParams object
// with the ability to set a context for a request.
func NewIgroupNestedCreateParamsWithContext(ctx context.Context) *IgroupNestedCreateParams {
	return &IgroupNestedCreateParams{
		Context: ctx,
	}
}

// NewIgroupNestedCreateParamsWithHTTPClient creates a new IgroupNestedCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIgroupNestedCreateParamsWithHTTPClient(client *http.Client) *IgroupNestedCreateParams {
	return &IgroupNestedCreateParams{
		HTTPClient: client,
	}
}

/* IgroupNestedCreateParams contains all the parameters to send to the API endpoint
   for the igroup nested create operation.

   Typically these are written to a http.Request.
*/
type IgroupNestedCreateParams struct {

	/* IgroupUUID.

	   The unique identifier of the parent initiator group.

	*/
	IgroupUUIDPathParameter string

	/* Info.

	   The properties of the nested initiator group to add to the initiator group.

	*/
	Info *models.IgroupNested

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecordsQueryParameter *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the igroup nested create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IgroupNestedCreateParams) WithDefaults() *IgroupNestedCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the igroup nested create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IgroupNestedCreateParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(false)
	)

	val := IgroupNestedCreateParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the igroup nested create params
func (o *IgroupNestedCreateParams) WithTimeout(timeout time.Duration) *IgroupNestedCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the igroup nested create params
func (o *IgroupNestedCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the igroup nested create params
func (o *IgroupNestedCreateParams) WithContext(ctx context.Context) *IgroupNestedCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the igroup nested create params
func (o *IgroupNestedCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the igroup nested create params
func (o *IgroupNestedCreateParams) WithHTTPClient(client *http.Client) *IgroupNestedCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the igroup nested create params
func (o *IgroupNestedCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIgroupUUIDPathParameter adds the igroupUUID to the igroup nested create params
func (o *IgroupNestedCreateParams) WithIgroupUUIDPathParameter(igroupUUID string) *IgroupNestedCreateParams {
	o.SetIgroupUUIDPathParameter(igroupUUID)
	return o
}

// SetIgroupUUIDPathParameter adds the igroupUuid to the igroup nested create params
func (o *IgroupNestedCreateParams) SetIgroupUUIDPathParameter(igroupUUID string) {
	o.IgroupUUIDPathParameter = igroupUUID
}

// WithInfo adds the info to the igroup nested create params
func (o *IgroupNestedCreateParams) WithInfo(info *models.IgroupNested) *IgroupNestedCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the igroup nested create params
func (o *IgroupNestedCreateParams) SetInfo(info *models.IgroupNested) {
	o.Info = info
}

// WithReturnRecordsQueryParameter adds the returnRecords to the igroup nested create params
func (o *IgroupNestedCreateParams) WithReturnRecordsQueryParameter(returnRecords *bool) *IgroupNestedCreateParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the igroup nested create params
func (o *IgroupNestedCreateParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *IgroupNestedCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param igroup.uuid
	if err := r.SetPathParam("igroup.uuid", o.IgroupUUIDPathParameter); err != nil {
		return err
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.ReturnRecordsQueryParameter != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecordsQueryParameter != nil {
			qrReturnRecords = *o.ReturnRecordsQueryParameter
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
