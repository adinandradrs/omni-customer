package model

type CustomerRegisterRequest struct {
	PhoneNo  string `json:"phone"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type CustomerRegisterResponse struct {
	RegistrationId string `json:"registrationId"`
	Email          string `json:"email"`
}
