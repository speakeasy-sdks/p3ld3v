// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/p3ld3v/v3/pkg/models/shared"
	"net/http"
)

type UpdatePetJSONResponse struct {
	Body []byte
	// HTTP response content type for this operation
	ContentType string
	// Successful operation
	Pet *shared.Pet
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdatePetJSONResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *UpdatePetJSONResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdatePetJSONResponse) GetPet() *shared.Pet {
	if o == nil {
		return nil
	}
	return o.Pet
}

func (o *UpdatePetJSONResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdatePetJSONResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
