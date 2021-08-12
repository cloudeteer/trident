// Code generated by go-swagger; DO NOT EDIT.

package networking

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
)

// NewHTTPProxyDeleteParams creates a new HTTPProxyDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHTTPProxyDeleteParams() *HTTPProxyDeleteParams {
	return &HTTPProxyDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHTTPProxyDeleteParamsWithTimeout creates a new HTTPProxyDeleteParams object
// with the ability to set a timeout on a request.
func NewHTTPProxyDeleteParamsWithTimeout(timeout time.Duration) *HTTPProxyDeleteParams {
	return &HTTPProxyDeleteParams{
		timeout: timeout,
	}
}

// NewHTTPProxyDeleteParamsWithContext creates a new HTTPProxyDeleteParams object
// with the ability to set a context for a request.
func NewHTTPProxyDeleteParamsWithContext(ctx context.Context) *HTTPProxyDeleteParams {
	return &HTTPProxyDeleteParams{
		Context: ctx,
	}
}

// NewHTTPProxyDeleteParamsWithHTTPClient creates a new HTTPProxyDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewHTTPProxyDeleteParamsWithHTTPClient(client *http.Client) *HTTPProxyDeleteParams {
	return &HTTPProxyDeleteParams{
		HTTPClient: client,
	}
}

/* HTTPProxyDeleteParams contains all the parameters to send to the API endpoint
   for the http proxy delete operation.

   Typically these are written to a http.Request.
*/
type HTTPProxyDeleteParams struct {

	/* UUID.

	   HTTP proxy UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the http proxy delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HTTPProxyDeleteParams) WithDefaults() *HTTPProxyDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the http proxy delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HTTPProxyDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the http proxy delete params
func (o *HTTPProxyDeleteParams) WithTimeout(timeout time.Duration) *HTTPProxyDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the http proxy delete params
func (o *HTTPProxyDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the http proxy delete params
func (o *HTTPProxyDeleteParams) WithContext(ctx context.Context) *HTTPProxyDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the http proxy delete params
func (o *HTTPProxyDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the http proxy delete params
func (o *HTTPProxyDeleteParams) WithHTTPClient(client *http.Client) *HTTPProxyDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the http proxy delete params
func (o *HTTPProxyDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUIDPathParameter adds the uuid to the http proxy delete params
func (o *HTTPProxyDeleteParams) WithUUIDPathParameter(uuid string) *HTTPProxyDeleteParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the http proxy delete params
func (o *HTTPProxyDeleteParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *HTTPProxyDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}