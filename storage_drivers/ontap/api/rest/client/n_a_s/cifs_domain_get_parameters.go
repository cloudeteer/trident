// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewCifsDomainGetParams creates a new CifsDomainGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsDomainGetParams() *CifsDomainGetParams {
	return &CifsDomainGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsDomainGetParamsWithTimeout creates a new CifsDomainGetParams object
// with the ability to set a timeout on a request.
func NewCifsDomainGetParamsWithTimeout(timeout time.Duration) *CifsDomainGetParams {
	return &CifsDomainGetParams{
		timeout: timeout,
	}
}

// NewCifsDomainGetParamsWithContext creates a new CifsDomainGetParams object
// with the ability to set a context for a request.
func NewCifsDomainGetParamsWithContext(ctx context.Context) *CifsDomainGetParams {
	return &CifsDomainGetParams{
		Context: ctx,
	}
}

// NewCifsDomainGetParamsWithHTTPClient creates a new CifsDomainGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsDomainGetParamsWithHTTPClient(client *http.Client) *CifsDomainGetParams {
	return &CifsDomainGetParams{
		HTTPClient: client,
	}
}

/* CifsDomainGetParams contains all the parameters to send to the API endpoint
   for the cifs domain get operation.

   Typically these are written to a http.Request.
*/
type CifsDomainGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* RediscoverTrusts.

	   Force the discovery of trusted domains.
	*/
	RediscoverTrustsQueryParameter *bool

	/* ResetDiscoveredServers.

	   Force a rediscovery.
	*/
	ResetDiscoveredServersQueryParameter *bool

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cifs domain get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsDomainGetParams) WithDefaults() *CifsDomainGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs domain get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsDomainGetParams) SetDefaults() {
	var (
		rediscoverTrustsQueryParameterDefault = bool(false)

		resetDiscoveredServersQueryParameterDefault = bool(false)
	)

	val := CifsDomainGetParams{
		RediscoverTrustsQueryParameter:       &rediscoverTrustsQueryParameterDefault,
		ResetDiscoveredServersQueryParameter: &resetDiscoveredServersQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cifs domain get params
func (o *CifsDomainGetParams) WithTimeout(timeout time.Duration) *CifsDomainGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs domain get params
func (o *CifsDomainGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs domain get params
func (o *CifsDomainGetParams) WithContext(ctx context.Context) *CifsDomainGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs domain get params
func (o *CifsDomainGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs domain get params
func (o *CifsDomainGetParams) WithHTTPClient(client *http.Client) *CifsDomainGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs domain get params
func (o *CifsDomainGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the cifs domain get params
func (o *CifsDomainGetParams) WithFieldsQueryParameter(fields []string) *CifsDomainGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the cifs domain get params
func (o *CifsDomainGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithRediscoverTrustsQueryParameter adds the rediscoverTrusts to the cifs domain get params
func (o *CifsDomainGetParams) WithRediscoverTrustsQueryParameter(rediscoverTrusts *bool) *CifsDomainGetParams {
	o.SetRediscoverTrustsQueryParameter(rediscoverTrusts)
	return o
}

// SetRediscoverTrustsQueryParameter adds the rediscoverTrusts to the cifs domain get params
func (o *CifsDomainGetParams) SetRediscoverTrustsQueryParameter(rediscoverTrusts *bool) {
	o.RediscoverTrustsQueryParameter = rediscoverTrusts
}

// WithResetDiscoveredServersQueryParameter adds the resetDiscoveredServers to the cifs domain get params
func (o *CifsDomainGetParams) WithResetDiscoveredServersQueryParameter(resetDiscoveredServers *bool) *CifsDomainGetParams {
	o.SetResetDiscoveredServersQueryParameter(resetDiscoveredServers)
	return o
}

// SetResetDiscoveredServersQueryParameter adds the resetDiscoveredServers to the cifs domain get params
func (o *CifsDomainGetParams) SetResetDiscoveredServersQueryParameter(resetDiscoveredServers *bool) {
	o.ResetDiscoveredServersQueryParameter = resetDiscoveredServers
}

// WithSVMUUIDPathParameter adds the svmUUID to the cifs domain get params
func (o *CifsDomainGetParams) WithSVMUUIDPathParameter(svmUUID string) *CifsDomainGetParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the cifs domain get params
func (o *CifsDomainGetParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *CifsDomainGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.RediscoverTrustsQueryParameter != nil {

		// query param rediscover_trusts
		var qrRediscoverTrusts bool

		if o.RediscoverTrustsQueryParameter != nil {
			qrRediscoverTrusts = *o.RediscoverTrustsQueryParameter
		}
		qRediscoverTrusts := swag.FormatBool(qrRediscoverTrusts)
		if qRediscoverTrusts != "" {

			if err := r.SetQueryParam("rediscover_trusts", qRediscoverTrusts); err != nil {
				return err
			}
		}
	}

	if o.ResetDiscoveredServersQueryParameter != nil {

		// query param reset_discovered_servers
		var qrResetDiscoveredServers bool

		if o.ResetDiscoveredServersQueryParameter != nil {
			qrResetDiscoveredServers = *o.ResetDiscoveredServersQueryParameter
		}
		qResetDiscoveredServers := swag.FormatBool(qrResetDiscoveredServers)
		if qResetDiscoveredServers != "" {

			if err := r.SetQueryParam("reset_discovered_servers", qResetDiscoveredServers); err != nil {
				return err
			}
		}
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCifsDomainGet binds the parameter fields
func (o *CifsDomainGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}