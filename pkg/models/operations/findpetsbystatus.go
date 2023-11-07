// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/models/shared"
	"github.com/speakeasy-sdks/p3ld3v/v2/pkg/utils"
	"net/http"
)

// Status values that need to be considered for filter
type Status string

const (
	StatusAvailable Status = "available"
	StatusPending   Status = "pending"
	StatusSold      Status = "sold"
)

func (e Status) ToPointer() *Status {
	return &e
}

func (e *Status) UnmarshalJSON(data []byte) error {
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
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

type FindPetsByStatusRequest struct {
	// Status values that need to be considered for filter
	Status *Status `default:"available" queryParam:"style=form,explode=true,name=status"`
}

func (f FindPetsByStatusRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FindPetsByStatusRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *FindPetsByStatusRequest) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

type FindPetsByStatusResponse struct {
	// successful operation
	TwoHundredApplicationJSONClasses []shared.Pet
	Body                             []byte
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *FindPetsByStatusResponse) GetTwoHundredApplicationJSONClasses() []shared.Pet {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONClasses
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
