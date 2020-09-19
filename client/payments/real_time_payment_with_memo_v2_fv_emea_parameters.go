// Code generated by go-swagger; DO NOT EDIT.

package payments

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

	"github.com/tooolbox/firstvision/models"
)

// NewRealTimePaymentWithMemoV2FvEmeaParams creates a new RealTimePaymentWithMemoV2FvEmeaParams object
// with the default values initialized.
func NewRealTimePaymentWithMemoV2FvEmeaParams() *RealTimePaymentWithMemoV2FvEmeaParams {
	var (
		appCodeDefault       = string("Please reach out to FirstVision API team to get the right values for your application. ")
		authorizationDefault = string("Bearer <token>")
		contentTypeDefault   = string("application/json")
	)
	return &RealTimePaymentWithMemoV2FvEmeaParams{
		AppCode:       &appCodeDefault,
		Authorization: authorizationDefault,
		ContentType:   &contentTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewRealTimePaymentWithMemoV2FvEmeaParamsWithTimeout creates a new RealTimePaymentWithMemoV2FvEmeaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRealTimePaymentWithMemoV2FvEmeaParamsWithTimeout(timeout time.Duration) *RealTimePaymentWithMemoV2FvEmeaParams {
	var (
		appCodeDefault       = string("Please reach out to FirstVision API team to get the right values for your application. ")
		authorizationDefault = string("Bearer <token>")
		contentTypeDefault   = string("application/json")
	)
	return &RealTimePaymentWithMemoV2FvEmeaParams{
		AppCode:       &appCodeDefault,
		Authorization: authorizationDefault,
		ContentType:   &contentTypeDefault,

		timeout: timeout,
	}
}

// NewRealTimePaymentWithMemoV2FvEmeaParamsWithContext creates a new RealTimePaymentWithMemoV2FvEmeaParams object
// with the default values initialized, and the ability to set a context for a request
func NewRealTimePaymentWithMemoV2FvEmeaParamsWithContext(ctx context.Context) *RealTimePaymentWithMemoV2FvEmeaParams {
	var (
		appCodeDefault       = string("Please reach out to FirstVision API team to get the right values for your application. ")
		authorizationDefault = string("Bearer <token>")
		contentTypeDefault   = string("application/json")
	)
	return &RealTimePaymentWithMemoV2FvEmeaParams{
		AppCode:       &appCodeDefault,
		Authorization: authorizationDefault,
		ContentType:   &contentTypeDefault,

		Context: ctx,
	}
}

// NewRealTimePaymentWithMemoV2FvEmeaParamsWithHTTPClient creates a new RealTimePaymentWithMemoV2FvEmeaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRealTimePaymentWithMemoV2FvEmeaParamsWithHTTPClient(client *http.Client) *RealTimePaymentWithMemoV2FvEmeaParams {
	var (
		appCodeDefault       = string("Please reach out to FirstVision API team to get the right values for your application. ")
		authorizationDefault = string("Bearer <token>")
		contentTypeDefault   = string("application/json")
	)
	return &RealTimePaymentWithMemoV2FvEmeaParams{
		AppCode:       &appCodeDefault,
		Authorization: authorizationDefault,
		ContentType:   &contentTypeDefault,
		HTTPClient:    client,
	}
}

/*RealTimePaymentWithMemoV2FvEmeaParams contains all the parameters to send to the API endpoint
for the real time payment with memo v2 fv emea operation typically these are written to a http.Request
*/
type RealTimePaymentWithMemoV2FvEmeaParams struct {

	/*AppCode
	  Application code identifying the calling application within a client system

	*/
	AppCode *string
	/*Authorization
	  OAuth2.0 access token (Bearer token) that you get from security API

	*/
	Authorization string
	/*ContentType
	  Request content type

	*/
	ContentType *string
	/*Body*/
	Body *models.RealTimePaymentWithMemoRequest2

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the real time payment with memo v2 fv emea params
func (o *RealTimePaymentWithMemoV2FvEmeaParams) WithTimeout(timeout time.Duration) *RealTimePaymentWithMemoV2FvEmeaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the real time payment with memo v2 fv emea params
func (o *RealTimePaymentWithMemoV2FvEmeaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the real time payment with memo v2 fv emea params
func (o *RealTimePaymentWithMemoV2FvEmeaParams) WithContext(ctx context.Context) *RealTimePaymentWithMemoV2FvEmeaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the real time payment with memo v2 fv emea params
func (o *RealTimePaymentWithMemoV2FvEmeaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the real time payment with memo v2 fv emea params
func (o *RealTimePaymentWithMemoV2FvEmeaParams) WithHTTPClient(client *http.Client) *RealTimePaymentWithMemoV2FvEmeaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the real time payment with memo v2 fv emea params
func (o *RealTimePaymentWithMemoV2FvEmeaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppCode adds the appCode to the real time payment with memo v2 fv emea params
func (o *RealTimePaymentWithMemoV2FvEmeaParams) WithAppCode(appCode *string) *RealTimePaymentWithMemoV2FvEmeaParams {
	o.SetAppCode(appCode)
	return o
}

// SetAppCode adds the appCode to the real time payment with memo v2 fv emea params
func (o *RealTimePaymentWithMemoV2FvEmeaParams) SetAppCode(appCode *string) {
	o.AppCode = appCode
}

// WithAuthorization adds the authorization to the real time payment with memo v2 fv emea params
func (o *RealTimePaymentWithMemoV2FvEmeaParams) WithAuthorization(authorization string) *RealTimePaymentWithMemoV2FvEmeaParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the real time payment with memo v2 fv emea params
func (o *RealTimePaymentWithMemoV2FvEmeaParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithContentType adds the contentType to the real time payment with memo v2 fv emea params
func (o *RealTimePaymentWithMemoV2FvEmeaParams) WithContentType(contentType *string) *RealTimePaymentWithMemoV2FvEmeaParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the real time payment with memo v2 fv emea params
func (o *RealTimePaymentWithMemoV2FvEmeaParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithBody adds the body to the real time payment with memo v2 fv emea params
func (o *RealTimePaymentWithMemoV2FvEmeaParams) WithBody(body *models.RealTimePaymentWithMemoRequest2) *RealTimePaymentWithMemoV2FvEmeaParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the real time payment with memo v2 fv emea params
func (o *RealTimePaymentWithMemoV2FvEmeaParams) SetBody(body *models.RealTimePaymentWithMemoRequest2) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RealTimePaymentWithMemoV2FvEmeaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppCode != nil {

		// header param AppCode
		if err := r.SetHeaderParam("AppCode", *o.AppCode); err != nil {
			return err
		}

	}

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.ContentType != nil {

		// header param Content-Type
		if err := r.SetHeaderParam("Content-Type", *o.ContentType); err != nil {
			return err
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
