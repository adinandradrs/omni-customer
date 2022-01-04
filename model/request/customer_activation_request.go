package request

type CustomerActivationRequest struct {
	ActivationId string `json:"activationId" binding:"required"`
}
