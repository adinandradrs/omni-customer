package request

type CustomerSigninRequest struct {
	Email    string `binding:"required"`
	Password string
	Otp      string
}
