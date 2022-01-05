package request

type CustomerActivationRequest struct {
	ActivationId string `binding:"required"`
}
