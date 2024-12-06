// Code generated by go-swagger; DO NOT EDIT.

package version_service

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

// NewVersionServiceMetadataParams creates a new VersionServiceMetadataParams object
// with the default values initialized.
func NewVersionServiceMetadataParams() *VersionServiceMetadataParams {
	var ()
	return &VersionServiceMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVersionServiceMetadataParamsWithTimeout creates a new VersionServiceMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVersionServiceMetadataParamsWithTimeout(timeout time.Duration) *VersionServiceMetadataParams {
	var ()
	return &VersionServiceMetadataParams{

		timeout: timeout,
	}
}

// NewVersionServiceMetadataParamsWithContext creates a new VersionServiceMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewVersionServiceMetadataParamsWithContext(ctx context.Context) *VersionServiceMetadataParams {
	var ()
	return &VersionServiceMetadataParams{

		Context: ctx,
	}
}

// NewVersionServiceMetadataParamsWithHTTPClient creates a new VersionServiceMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVersionServiceMetadataParamsWithHTTPClient(client *http.Client) *VersionServiceMetadataParams {
	var ()
	return &VersionServiceMetadataParams{
		HTTPClient: client,
	}
}

/*VersionServiceMetadataParams contains all the parameters to send to the API endpoint
for the version service metadata operation typically these are written to a http.Request
*/
type VersionServiceMetadataParams struct {

	/*Product*/
	Product string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the version service metadata params
func (o *VersionServiceMetadataParams) WithTimeout(timeout time.Duration) *VersionServiceMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the version service metadata params
func (o *VersionServiceMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the version service metadata params
func (o *VersionServiceMetadataParams) WithContext(ctx context.Context) *VersionServiceMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the version service metadata params
func (o *VersionServiceMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the version service metadata params
func (o *VersionServiceMetadataParams) WithHTTPClient(client *http.Client) *VersionServiceMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the version service metadata params
func (o *VersionServiceMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProduct adds the product to the version service metadata params
func (o *VersionServiceMetadataParams) WithProduct(product string) *VersionServiceMetadataParams {
	o.SetProduct(product)
	return o
}

// SetProduct adds the product to the version service metadata params
func (o *VersionServiceMetadataParams) SetProduct(product string) {
	o.Product = product
}

// WriteToRequest writes these params to a swagger request
func (o *VersionServiceMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param product
	if err := r.SetPathParam("product", o.Product); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
