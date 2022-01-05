package response

type CustomerLoginResponse struct {
	Token    string `json:"token"`
	UserId   int    `json:"userId"`
	Email    string `json:"email"`
	Fullname string `json:"email"`
}
