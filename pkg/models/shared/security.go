// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Security struct {
	PetstoreAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

func (o *Security) GetPetstoreAuth() string {
	if o == nil {
		return ""
	}
	return o.PetstoreAuth
}