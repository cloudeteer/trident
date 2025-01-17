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
)

// NewWebGetParams creates a new WebGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWebGetParams() *WebGetParams {
	return &WebGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWebGetParamsWithTimeout creates a new WebGetParams object
// with the ability to set a timeout on a request.
func NewWebGetParamsWithTimeout(timeout time.Duration) *WebGetParams {
	return &WebGetParams{
		timeout: timeout,
	}
}

// NewWebGetParamsWithContext creates a new WebGetParams object
// with the ability to set a context for a request.
func NewWebGetParamsWithContext(ctx context.Context) *WebGetParams {
	return &WebGetParams{
		Context: ctx,
	}
}

// NewWebGetParamsWithHTTPClient creates a new WebGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewWebGetParamsWithHTTPClient(client *http.Client) *WebGetParams {
	return &WebGetParams{
		HTTPClient: client,
	}
}

/* WebGetParams contains all the parameters to send to the API endpoint
   for the web get operation.

   Typically these are written to a http.Request.
*/
type WebGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the web get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WebGetParams) WithDefaults() *WebGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the web get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WebGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the web get params
func (o *WebGetParams) WithTimeout(timeout time.Duration) *WebGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the web get params
func (o *WebGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the web get params
func (o *WebGetParams) WithContext(ctx context.Context) *WebGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the web get params
func (o *WebGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the web get params
func (o *WebGetParams) WithHTTPClient(client *http.Client) *WebGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the web get params
func (o *WebGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the web get params
func (o *WebGetParams) WithFieldsQueryParameter(fields []string) *WebGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the web get params
func (o *WebGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WriteToRequest writes these params to a swagger request
func (o *WebGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamWebGet binds the parameter fields
func (o *WebGetParams) bindParamFields(formats strfmt.Registry) []string {
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
