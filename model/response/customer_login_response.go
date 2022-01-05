package response

type CustomerLoginResponse struct {
	Token    string
	UserId   int
	Email    string
	Fullname string
}
