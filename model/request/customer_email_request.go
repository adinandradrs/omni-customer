package request

type CustomerEmailRequest struct {
	Email string `binding:"required"`
}
