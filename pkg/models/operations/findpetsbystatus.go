// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/shared"
	"net/http"
)

type FindPetsByStatusSecurity struct {
	PetstoreAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

func (o *FindPetsByStatusSecurity) GetPetstoreAuth() string {
	if o == nil {
		return ""
	}
	return o.PetstoreAuth
}

// FindPetsByStatusStatus - Status values that need to be considered for filter
type FindPetsByStatusStatus string

const (
	FindPetsByStatusStatusAvailable FindPetsByStatusStatus = "available"
	FindPetsByStatusStatusPending   FindPetsByStatusStatus = "pending"
	FindPetsByStatusStatusSold      FindPetsByStatusStatus = "sold"
)

func (e FindPetsByStatusStatus) ToPointer() *FindPetsByStatusStatus {
	return &e
}

func (e *FindPetsByStatusStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "available":
		fallthrough
	case "pending":
		fallthrough
	case "sold":
		*e = FindPetsByStatusStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindPetsByStatusStatus: %v", v)
	}
}

type FindPetsByStatusRequest struct {
	// Status values that need to be considered for filter
	Status *FindPetsByStatusStatus `queryParam:"style=form,explode=true,name=status"`
}

func (o *FindPetsByStatusRequest) GetStatus() *FindPetsByStatusStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type FindPetsByStatusResponse struct {
	Body        []byte
	ContentType string
	// successful operation
	Pets        []shared.Pet
	StatusCode  int
	RawResponse *http.Response
}

func (o *FindPetsByStatusResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *FindPetsByStatusResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *FindPetsByStatusResponse) GetPets() []shared.Pet {
	if o == nil {
		return nil
	}
	return o.Pets
}

func (o *FindPetsByStatusResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *FindPetsByStatusResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
