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

// NewMetroclusterDiagnosticsCreateParams creates a new MetroclusterDiagnosticsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMetroclusterDiagnosticsCreateParams() *MetroclusterDiagnosticsCreateParams {
	return &MetroclusterDiagnosticsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMetroclusterDiagnosticsCreateParamsWithTimeout creates a new MetroclusterDiagnosticsCreateParams object
// with the ability to set a timeout on a request.
func NewMetroclusterDiagnosticsCreateParamsWithTimeout(timeout time.Duration) *MetroclusterDiagnosticsCreateParams {
	return &MetroclusterDiagnosticsCreateParams{
		timeout: timeout,
	}
}

// NewMetroclusterDiagnosticsCreateParamsWithContext creates a new MetroclusterDiagnosticsCreateParams object
// with the ability to set a context for a request.
func NewMetroclusterDiagnosticsCreateParamsWithContext(ctx context.Context) *MetroclusterDiagnosticsCreateParams {
	return &MetroclusterDiagnosticsCreateParams{
		Context: ctx,
	}
}

// NewMetroclusterDiagnosticsCreateParamsWithHTTPClient creates a new MetroclusterDiagnosticsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewMetroclusterDiagnosticsCreateParamsWithHTTPClient(client *http.Client) *MetroclusterDiagnosticsCreateParams {
	return &MetroclusterDiagnosticsCreateParams{
		HTTPClient: client,
	}
}

/* MetroclusterDiagnosticsCreateParams contains all the parameters to send to the API endpoint
   for the metrocluster diagnostics create operation.

   Typically these are written to a http.Request.
*/
type MetroclusterDiagnosticsCreateParams struct {

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeoutQueryParameter *int64

	/* Schedule.

	   Shows the minutes of every hour when a job runs. Setting this parameter schedules the periodic job to be run to perform MetroCluster diagnostic.
	*/
	ScheduleQueryParameter *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the metrocluster diagnostics create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetroclusterDiagnosticsCreateParams) WithDefaults() *MetroclusterDiagnosticsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the metrocluster diagnostics create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetroclusterDiagnosticsCreateParams) SetDefaults() {
	var (
		returnTimeoutQueryParameterDefault = int64(0)
	)

	val := MetroclusterDiagnosticsCreateParams{
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the metrocluster diagnostics create params
func (o *MetroclusterDiagnosticsCreateParams) WithTimeout(timeout time.Duration) *MetroclusterDiagnosticsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the metrocluster diagnostics create params
func (o *MetroclusterDiagnosticsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the metrocluster diagnostics create params
func (o *MetroclusterDiagnosticsCreateParams) WithContext(ctx context.Context) *MetroclusterDiagnosticsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the metrocluster diagnostics create params
func (o *MetroclusterDiagnosticsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the metrocluster diagnostics create params
func (o *MetroclusterDiagnosticsCreateParams) WithHTTPClient(client *http.Client) *MetroclusterDiagnosticsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the metrocluster diagnostics create params
func (o *MetroclusterDiagnosticsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the metrocluster diagnostics create params
func (o *MetroclusterDiagnosticsCreateParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *MetroclusterDiagnosticsCreateParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the metrocluster diagnostics create params
func (o *MetroclusterDiagnosticsCreateParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithScheduleQueryParameter adds the schedule to the metrocluster diagnostics create params
func (o *MetroclusterDiagnosticsCreateParams) WithScheduleQueryParameter(schedule *int64) *MetroclusterDiagnosticsCreateParams {
	o.SetScheduleQueryParameter(schedule)
	return o
}

// SetScheduleQueryParameter adds the schedule to the metrocluster diagnostics create params
func (o *MetroclusterDiagnosticsCreateParams) SetScheduleQueryParameter(schedule *int64) {
	o.ScheduleQueryParameter = schedule
}

// WriteToRequest writes these params to a swagger request
func (o *MetroclusterDiagnosticsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.ScheduleQueryParameter != nil {

		// query param schedule
		var qrSchedule int64

		if o.ScheduleQueryParameter != nil {
			qrSchedule = *o.ScheduleQueryParameter
		}
		qSchedule := swag.FormatInt64(qrSchedule)
		if qSchedule != "" {

			if err := r.SetQueryParam("schedule", qSchedule); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
