// Code generated by go-swagger; DO NOT EDIT.

package account_maintenance

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

// NewDirectDebitUpdateV2FvEmeaParams creates a new DirectDebitUpdateV2FvEmeaParams object
// with the default values initialized.
func NewDirectDebitUpdateV2FvEmeaParams() *DirectDebitUpdateV2FvEmeaParams {
	var (
		appCodeDefault       = string("Please reach out to FirstVision API team to get the right values for your application. ")
		authorizationDefault = string("Bearer <token>")
		contentTypeDefault   = string("application/json")
	)
	return &DirectDebitUpdateV2FvEmeaParams{
		AppCode:       &appCodeDefault,
		Authorization: authorizationDefault,
		ContentType:   &contentTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDirectDebitUpdateV2FvEmeaParamsWithTimeout creates a new DirectDebitUpdateV2FvEmeaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDirectDebitUpdateV2FvEmeaParamsWithTimeout(timeout time.Duration) *DirectDebitUpdateV2FvEmeaParams {
	var (
		appCodeDefault       = string("Please reach out to FirstVision API team to get the right values for your application. ")
		authorizationDefault = string("Bearer <token>")
		contentTypeDefault   = string("application/json")
	)
	return &DirectDebitUpdateV2FvEmeaParams{
		AppCode:       &appCodeDefault,
		Authorization: authorizationDefault,
		ContentType:   &contentTypeDefault,

		timeout: timeout,
	}
}

// NewDirectDebitUpdateV2FvEmeaParamsWithContext creates a new DirectDebitUpdateV2FvEmeaParams object
// with the default values initialized, and the ability to set a context for a request
func NewDirectDebitUpdateV2FvEmeaParamsWithContext(ctx context.Context) *DirectDebitUpdateV2FvEmeaParams {
	var (
		appCodeDefault       = string("Please reach out to FirstVision API team to get the right values for your application. ")
		authorizationDefault = string("Bearer <token>")
		contentTypeDefault   = string("application/json")
	)
	return &DirectDebitUpdateV2FvEmeaParams{
		AppCode:       &appCodeDefault,
		Authorization: authorizationDefault,
		ContentType:   &contentTypeDefault,

		Context: ctx,
	}
}

// NewDirectDebitUpdateV2FvEmeaParamsWithHTTPClient creates a new DirectDebitUpdateV2FvEmeaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDirectDebitUpdateV2FvEmeaParamsWithHTTPClient(client *http.Client) *DirectDebitUpdateV2FvEmeaParams {
	var (
		appCodeDefault       = string("Please reach out to FirstVision API team to get the right values for your application. ")
		authorizationDefault = string("Bearer <token>")
		contentTypeDefault   = string("application/json")
	)
	return &DirectDebitUpdateV2FvEmeaParams{
		AppCode:       &appCodeDefault,
		Authorization: authorizationDefault,
		ContentType:   &contentTypeDefault,
		HTTPClient:    client,
	}
}

/*DirectDebitUpdateV2FvEmeaParams contains all the parameters to send to the API endpoint
for the direct debit update v2 fv emea operation typically these are written to a http.Request
*/
type DirectDebitUpdateV2FvEmeaParams struct {

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
	Body *models.DirectDebitUpdateRequest2

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the direct debit update v2 fv emea params
func (o *DirectDebitUpdateV2FvEmeaParams) WithTimeout(timeout time.Duration) *DirectDebitUpdateV2FvEmeaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the direct debit update v2 fv emea params
func (o *DirectDebitUpdateV2FvEmeaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the direct debit update v2 fv emea params
func (o *DirectDebitUpdateV2FvEmeaParams) WithContext(ctx context.Context) *DirectDebitUpdateV2FvEmeaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the direct debit update v2 fv emea params
func (o *DirectDebitUpdateV2FvEmeaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the direct debit update v2 fv emea params
func (o *DirectDebitUpdateV2FvEmeaParams) WithHTTPClient(client *http.Client) *DirectDebitUpdateV2FvEmeaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the direct debit update v2 fv emea params
func (o *DirectDebitUpdateV2FvEmeaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppCode adds the appCode to the direct debit update v2 fv emea params
func (o *DirectDebitUpdateV2FvEmeaParams) WithAppCode(appCode *string) *DirectDebitUpdateV2FvEmeaParams {
	o.SetAppCode(appCode)
	return o
}

// SetAppCode adds the appCode to the direct debit update v2 fv emea params
func (o *DirectDebitUpdateV2FvEmeaParams) SetAppCode(appCode *string) {
	o.AppCode = appCode
}

// WithAuthorization adds the authorization to the direct debit update v2 fv emea params
func (o *DirectDebitUpdateV2FvEmeaParams) WithAuthorization(authorization string) *DirectDebitUpdateV2FvEmeaParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the direct debit update v2 fv emea params
func (o *DirectDebitUpdateV2FvEmeaParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithContentType adds the contentType to the direct debit update v2 fv emea params
func (o *DirectDebitUpdateV2FvEmeaParams) WithContentType(contentType *string) *DirectDebitUpdateV2FvEmeaParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the direct debit update v2 fv emea params
func (o *DirectDebitUpdateV2FvEmeaParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithBody adds the body to the direct debit update v2 fv emea params
func (o *DirectDebitUpdateV2FvEmeaParams) WithBody(body *models.DirectDebitUpdateRequest2) *DirectDebitUpdateV2FvEmeaParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the direct debit update v2 fv emea params
func (o *DirectDebitUpdateV2FvEmeaParams) SetBody(body *models.DirectDebitUpdateRequest2) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DirectDebitUpdateV2FvEmeaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
