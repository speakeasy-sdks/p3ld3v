// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/p3ld3v/v3/pkg/models/shared"
	"net/http"
)

type GetOrderByIDRequest struct {
	// ID of order that needs to be fetched
	OrderID int64 `pathParam:"style=simple,explode=false,name=orderId"`
}

func (o *GetOrderByIDRequest) GetOrderID() int64 {
	if o == nil {
		return 0
	}
	return o.OrderID
}

type GetOrderByIDResponse struct {
	Body []byte
	// HTTP response content type for this operation
	ContentType string
	// successful operation
	Order *shared.Order
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetOrderByIDResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *GetOrderByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetOrderByIDResponse) GetOrder() *shared.Order {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *GetOrderByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetOrderByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
