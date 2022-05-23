package myservice

import (
	"encoding/json"

	"github.com/hooklift/gowsdl/soap"
)

func (xdt *DateTime) MarshalJSON() (result []byte, err error) {
	return (*soap.XSDDateTime)(xdt).ToGoTime().MarshalJSON()
}

func (arr *ArrayOfPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(arr.Post)
}
