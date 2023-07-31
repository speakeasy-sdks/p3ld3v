// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/p3ld3v/pkg/models/shared"
	"net/http"
)

type GetUserByNameRequest struct {
	// The name that needs to be fetched. Use user1 for testing.
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

func (o *GetUserByNameRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type GetUserByNameResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// successful operation
	User *shared.User
}

func (o *GetUserByNameResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *GetUserByNameResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUserByNameResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUserByNameResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetUserByNameResponse) GetUser() *shared.User {
	if o == nil {
		return nil
	}
	return o.User
}
