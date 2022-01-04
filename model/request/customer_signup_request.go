package request

type CustomerSignupRequest struct {
	Email    string `json:"email" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Password string `json:"password" binding:"required"`
}
