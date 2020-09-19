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

// NewAccountDetailInquiryV3FvEmeaParams creates a new AccountDetailInquiryV3FvEmeaParams object
// with the default values initialized.
func NewAccountDetailInquiryV3FvEmeaParams() *AccountDetailInquiryV3FvEmeaParams {
	var (
		appCodeDefault       = string("Please reach out to FirstVision API team to get the right values for your application. ")
		authorizationDefault = string("Bearer <token>")
		contentTypeDefault   = string("application/json")
	)
	return &AccountDetailInquiryV3FvEmeaParams{
		AppCode:       &appCodeDefault,
		Authorization: authorizationDefault,
		ContentType:   &contentTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewAccountDetailInquiryV3FvEmeaParamsWithTimeout creates a new AccountDetailInquiryV3FvEmeaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAccountDetailInquiryV3FvEmeaParamsWithTimeout(timeout time.Duration) *AccountDetailInquiryV3FvEmeaParams {
	var (
		appCodeDefault       = string("Please reach out to FirstVision API team to get the right values for your application. ")
		authorizationDefault = string("Bearer <token>")
		contentTypeDefault   = string("application/json")
	)
	return &AccountDetailInquiryV3FvEmeaParams{
		AppCode:       &appCodeDefault,
		Authorization: authorizationDefault,
		ContentType:   &contentTypeDefault,

		timeout: timeout,
	}
}

// NewAccountDetailInquiryV3FvEmeaParamsWithContext creates a new AccountDetailInquiryV3FvEmeaParams object
// with the default values initialized, and the ability to set a context for a request
func NewAccountDetailInquiryV3FvEmeaParamsWithContext(ctx context.Context) *AccountDetailInquiryV3FvEmeaParams {
	var (
		appCodeDefault       = string("Please reach out to FirstVision API team to get the right values for your application. ")
		authorizationDefault = string("Bearer <token>")
		contentTypeDefault   = string("application/json")
	)
	return &AccountDetailInquiryV3FvEmeaParams{
		AppCode:       &appCodeDefault,
		Authorization: authorizationDefault,
		ContentType:   &contentTypeDefault,

		Context: ctx,
	}
}

// NewAccountDetailInquiryV3FvEmeaParamsWithHTTPClient creates a new AccountDetailInquiryV3FvEmeaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAccountDetailInquiryV3FvEmeaParamsWithHTTPClient(client *http.Client) *AccountDetailInquiryV3FvEmeaParams {
	var (
		appCodeDefault       = string("Please reach out to FirstVision API team to get the right values for your application. ")
		authorizationDefault = string("Bearer <token>")
		contentTypeDefault   = string("application/json")
	)
	return &AccountDetailInquiryV3FvEmeaParams{
		AppCode:       &appCodeDefault,
		Authorization: authorizationDefault,
		ContentType:   &contentTypeDefault,
		HTTPClient:    client,
	}
}

/*AccountDetailInquiryV3FvEmeaParams contains all the parameters to send to the API endpoint
for the account detail inquiry v3 fv emea operation typically these are written to a http.Request
*/
type AccountDetailInquiryV3FvEmeaParams struct {

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
	Body *models.AccountDetailInquiryRequest3

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the account detail inquiry v3 fv emea params
func (o *AccountDetailInquiryV3FvEmeaParams) WithTimeout(timeout time.Duration) *AccountDetailInquiryV3FvEmeaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account detail inquiry v3 fv emea params
func (o *AccountDetailInquiryV3FvEmeaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account detail inquiry v3 fv emea params
func (o *AccountDetailInquiryV3FvEmeaParams) WithContext(ctx context.Context) *AccountDetailInquiryV3FvEmeaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account detail inquiry v3 fv emea params
func (o *AccountDetailInquiryV3FvEmeaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account detail inquiry v3 fv emea params
func (o *AccountDetailInquiryV3FvEmeaParams) WithHTTPClient(client *http.Client) *AccountDetailInquiryV3FvEmeaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account detail inquiry v3 fv emea params
func (o *AccountDetailInquiryV3FvEmeaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppCode adds the appCode to the account detail inquiry v3 fv emea params
func (o *AccountDetailInquiryV3FvEmeaParams) WithAppCode(appCode *string) *AccountDetailInquiryV3FvEmeaParams {
	o.SetAppCode(appCode)
	return o
}

// SetAppCode adds the appCode to the account detail inquiry v3 fv emea params
func (o *AccountDetailInquiryV3FvEmeaParams) SetAppCode(appCode *string) {
	o.AppCode = appCode
}

// WithAuthorization adds the authorization to the account detail inquiry v3 fv emea params
func (o *AccountDetailInquiryV3FvEmeaParams) WithAuthorization(authorization string) *AccountDetailInquiryV3FvEmeaParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the account detail inquiry v3 fv emea params
func (o *AccountDetailInquiryV3FvEmeaParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithContentType adds the contentType to the account detail inquiry v3 fv emea params
func (o *AccountDetailInquiryV3FvEmeaParams) WithContentType(contentType *string) *AccountDetailInquiryV3FvEmeaParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the account detail inquiry v3 fv emea params
func (o *AccountDetailInquiryV3FvEmeaParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithBody adds the body to the account detail inquiry v3 fv emea params
func (o *AccountDetailInquiryV3FvEmeaParams) WithBody(body *models.AccountDetailInquiryRequest3) *AccountDetailInquiryV3FvEmeaParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the account detail inquiry v3 fv emea params
func (o *AccountDetailInquiryV3FvEmeaParams) SetBody(body *models.AccountDetailInquiryRequest3) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AccountDetailInquiryV3FvEmeaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
