// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeletePetRequest struct {
	APIKey *string `header:"style=simple,explode=false,name=api_key"`
	// Pet id to delete
	PetID int64 `pathParam:"style=simple,explode=false,name=petId"`
}

func (o *DeletePetRequest) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}

func (o *DeletePetRequest) GetPetID() int64 {
	if o == nil {
		return 0
	}
	return o.PetID
}

type DeletePetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeletePetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeletePetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeletePetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
