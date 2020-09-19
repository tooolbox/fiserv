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

// NewMoveFundsV1FvEmeaParams creates a new MoveFundsV1FvEmeaParams object
// with the default values initialized.
func NewMoveFundsV1FvEmeaParams() *MoveFundsV1FvEmeaParams {
	var (
		appCodeDefault       = string("Please reach out to FirstVision API team to get the right values for your application. ")
		authorizationDefault = string("Bearer <token>")
		contentTypeDefault   = string("application/json")
	)
	return &MoveFundsV1FvEmeaParams{
		AppCode:       &appCodeDefault,
		Authorization: authorizationDefault,
		ContentType:   &contentTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewMoveFundsV1FvEmeaParamsWithTimeout creates a new MoveFundsV1FvEmeaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMoveFundsV1FvEmeaParamsWithTimeout(timeout time.Duration) *MoveFundsV1FvEmeaParams {
	var (
		appCodeDefault       = string("Please reach out to FirstVision API team to get the right values for your application. ")
		authorizationDefault = string("Bearer <token>")
		contentTypeDefault   = string("application/json")
	)
	return &MoveFundsV1FvEmeaParams{
		AppCode:       &appCodeDefault,
		Authorization: authorizationDefault,
		ContentType:   &contentTypeDefault,

		timeout: timeout,
	}
}

// NewMoveFundsV1FvEmeaParamsWithContext creates a new MoveFundsV1FvEmeaParams object
// with the default values initialized, and the ability to set a context for a request
func NewMoveFundsV1FvEmeaParamsWithContext(ctx context.Context) *MoveFundsV1FvEmeaParams {
	var (
		appCodeDefault       = string("Please reach out to FirstVision API team to get the right values for your application. ")
		authorizationDefault = string("Bearer <token>")
		contentTypeDefault   = string("application/json")
	)
	return &MoveFundsV1FvEmeaParams{
		AppCode:       &appCodeDefault,
		Authorization: authorizationDefault,
		ContentType:   &contentTypeDefault,

		Context: ctx,
	}
}

// NewMoveFundsV1FvEmeaParamsWithHTTPClient creates a new MoveFundsV1FvEmeaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMoveFundsV1FvEmeaParamsWithHTTPClient(client *http.Client) *MoveFundsV1FvEmeaParams {
	var (
		appCodeDefault       = string("Please reach out to FirstVision API team to get the right values for your application. ")
		authorizationDefault = string("Bearer <token>")
		contentTypeDefault   = string("application/json")
	)
	return &MoveFundsV1FvEmeaParams{
		AppCode:       &appCodeDefault,
		Authorization: authorizationDefault,
		ContentType:   &contentTypeDefault,
		HTTPClient:    client,
	}
}

/*MoveFundsV1FvEmeaParams contains all the parameters to send to the API endpoint
for the move funds v1 fv emea operation typically these are written to a http.Request
*/
type MoveFundsV1FvEmeaParams struct {

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
	Body *models.MoveFundsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the move funds v1 fv emea params
func (o *MoveFundsV1FvEmeaParams) WithTimeout(timeout time.Duration) *MoveFundsV1FvEmeaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the move funds v1 fv emea params
func (o *MoveFundsV1FvEmeaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the move funds v1 fv emea params
func (o *MoveFundsV1FvEmeaParams) WithContext(ctx context.Context) *MoveFundsV1FvEmeaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the move funds v1 fv emea params
func (o *MoveFundsV1FvEmeaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the move funds v1 fv emea params
func (o *MoveFundsV1FvEmeaParams) WithHTTPClient(client *http.Client) *MoveFundsV1FvEmeaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the move funds v1 fv emea params
func (o *MoveFundsV1FvEmeaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppCode adds the appCode to the move funds v1 fv emea params
func (o *MoveFundsV1FvEmeaParams) WithAppCode(appCode *string) *MoveFundsV1FvEmeaParams {
	o.SetAppCode(appCode)
	return o
}

// SetAppCode adds the appCode to the move funds v1 fv emea params
func (o *MoveFundsV1FvEmeaParams) SetAppCode(appCode *string) {
	o.AppCode = appCode
}

// WithAuthorization adds the authorization to the move funds v1 fv emea params
func (o *MoveFundsV1FvEmeaParams) WithAuthorization(authorization string) *MoveFundsV1FvEmeaParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the move funds v1 fv emea params
func (o *MoveFundsV1FvEmeaParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithContentType adds the contentType to the move funds v1 fv emea params
func (o *MoveFundsV1FvEmeaParams) WithContentType(contentType *string) *MoveFundsV1FvEmeaParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the move funds v1 fv emea params
func (o *MoveFundsV1FvEmeaParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithBody adds the body to the move funds v1 fv emea params
func (o *MoveFundsV1FvEmeaParams) WithBody(body *models.MoveFundsRequest) *MoveFundsV1FvEmeaParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the move funds v1 fv emea params
func (o *MoveFundsV1FvEmeaParams) SetBody(body *models.MoveFundsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *MoveFundsV1FvEmeaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
