package model

type CustomerActivationRequest struct {
	Code    string `json:"code"`
	PhoneNo string `json:"phoneNo"`
}
