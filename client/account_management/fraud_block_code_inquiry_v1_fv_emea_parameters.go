// Code generated by go-swagger; DO NOT EDIT.

package account_management

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

// NewFraudBlockCodeInquiryV1FvEmeaParams creates a new FraudBlockCodeInquiryV1FvEmeaParams object
// with the default values initialized.
func NewFraudBlockCodeInquiryV1FvEmeaParams() *FraudBlockCodeInquiryV1FvEmeaParams {
	var (
		appCodeDefault       = string("Please reach out to FirstVision API team to get the right values for your application. ")
		authorizationDefault = string("Bearer <token>")
		contentTypeDefault   = string("application/json")
	)
	return &FraudBlockCodeInquiryV1FvEmeaParams{
		AppCode:       &appCodeDefault,
		Authorization: authorizationDefault,
		ContentType:   &contentTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewFraudBlockCodeInquiryV1FvEmeaParamsWithTimeout creates a new FraudBlockCodeInquiryV1FvEmeaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFraudBlockCodeInquiryV1FvEmeaParamsWithTimeout(timeout time.Duration) *FraudBlockCodeInquiryV1FvEmeaParams {
	var (
		appCodeDefault       = string("Please reach out to FirstVision API team to get the right values for your application. ")
		authorizationDefault = string("Bearer <token>")
		contentTypeDefault   = string("application/json")
	)
	return &FraudBlockCodeInquiryV1FvEmeaParams{
		AppCode:       &appCodeDefault,
		Authorization: authorizationDefault,
		ContentType:   &contentTypeDefault,

		timeout: timeout,
	}
}

// NewFraudBlockCodeInquiryV1FvEmeaParamsWithContext creates a new FraudBlockCodeInquiryV1FvEmeaParams object
// with the default values initialized, and the ability to set a context for a request
func NewFraudBlockCodeInquiryV1FvEmeaParamsWithContext(ctx context.Context) *FraudBlockCodeInquiryV1FvEmeaParams {
	var (
		appCodeDefault       = string("Please reach out to FirstVision API team to get the right values for your application. ")
		authorizationDefault = string("Bearer <token>")
		contentTypeDefault   = string("application/json")
	)
	return &FraudBlockCodeInquiryV1FvEmeaParams{
		AppCode:       &appCodeDefault,
		Authorization: authorizationDefault,
		ContentType:   &contentTypeDefault,

		Context: ctx,
	}
}

// NewFraudBlockCodeInquiryV1FvEmeaParamsWithHTTPClient creates a new FraudBlockCodeInquiryV1FvEmeaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFraudBlockCodeInquiryV1FvEmeaParamsWithHTTPClient(client *http.Client) *FraudBlockCodeInquiryV1FvEmeaParams {
	var (
		appCodeDefault       = string("Please reach out to FirstVision API team to get the right values for your application. ")
		authorizationDefault = string("Bearer <token>")
		contentTypeDefault   = string("application/json")
	)
	return &FraudBlockCodeInquiryV1FvEmeaParams{
		AppCode:       &appCodeDefault,
		Authorization: authorizationDefault,
		ContentType:   &contentTypeDefault,
		HTTPClient:    client,
	}
}

/*FraudBlockCodeInquiryV1FvEmeaParams contains all the parameters to send to the API endpoint
for the fraud block code inquiry v1 fv emea operation typically these are written to a http.Request
*/
type FraudBlockCodeInquiryV1FvEmeaParams struct {

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
	Body *models.FraudBlockCodeInquiryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the fraud block code inquiry v1 fv emea params
func (o *FraudBlockCodeInquiryV1FvEmeaParams) WithTimeout(timeout time.Duration) *FraudBlockCodeInquiryV1FvEmeaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fraud block code inquiry v1 fv emea params
func (o *FraudBlockCodeInquiryV1FvEmeaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fraud block code inquiry v1 fv emea params
func (o *FraudBlockCodeInquiryV1FvEmeaParams) WithContext(ctx context.Context) *FraudBlockCodeInquiryV1FvEmeaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fraud block code inquiry v1 fv emea params
func (o *FraudBlockCodeInquiryV1FvEmeaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fraud block code inquiry v1 fv emea params
func (o *FraudBlockCodeInquiryV1FvEmeaParams) WithHTTPClient(client *http.Client) *FraudBlockCodeInquiryV1FvEmeaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fraud block code inquiry v1 fv emea params
func (o *FraudBlockCodeInquiryV1FvEmeaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppCode adds the appCode to the fraud block code inquiry v1 fv emea params
func (o *FraudBlockCodeInquiryV1FvEmeaParams) WithAppCode(appCode *string) *FraudBlockCodeInquiryV1FvEmeaParams {
	o.SetAppCode(appCode)
	return o
}

// SetAppCode adds the appCode to the fraud block code inquiry v1 fv emea params
func (o *FraudBlockCodeInquiryV1FvEmeaParams) SetAppCode(appCode *string) {
	o.AppCode = appCode
}

// WithAuthorization adds the authorization to the fraud block code inquiry v1 fv emea params
func (o *FraudBlockCodeInquiryV1FvEmeaParams) WithAuthorization(authorization string) *FraudBlockCodeInquiryV1FvEmeaParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the fraud block code inquiry v1 fv emea params
func (o *FraudBlockCodeInquiryV1FvEmeaParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithContentType adds the contentType to the fraud block code inquiry v1 fv emea params
func (o *FraudBlockCodeInquiryV1FvEmeaParams) WithContentType(contentType *string) *FraudBlockCodeInquiryV1FvEmeaParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the fraud block code inquiry v1 fv emea params
func (o *FraudBlockCodeInquiryV1FvEmeaParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithBody adds the body to the fraud block code inquiry v1 fv emea params
func (o *FraudBlockCodeInquiryV1FvEmeaParams) WithBody(body *models.FraudBlockCodeInquiryRequest) *FraudBlockCodeInquiryV1FvEmeaParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the fraud block code inquiry v1 fv emea params
func (o *FraudBlockCodeInquiryV1FvEmeaParams) SetBody(body *models.FraudBlockCodeInquiryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *FraudBlockCodeInquiryV1FvEmeaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
