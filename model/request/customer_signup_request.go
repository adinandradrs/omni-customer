package request

type CustomerSignupRequest struct {
	Email    string `binding:"required"`
	Fullname string `binding:"required"`
	Password string `binding:"required"`
}
