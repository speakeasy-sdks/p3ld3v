// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/p3ld3v/v3/pkg/models/shared"
	"net/http"
)

type PlaceOrderRawResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// successful operation
	Order *shared.Order
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PlaceOrderRawResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PlaceOrderRawResponse) GetOrder() *shared.Order {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *PlaceOrderRawResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PlaceOrderRawResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
