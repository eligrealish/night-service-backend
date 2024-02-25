package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"night-service-backend/domain/services"
)

// TODO reviewing threading of local scope vs global e.g should i use VAR
var (
	policyService services.PolicyService
)

type PolicyEndpoint struct {
	PolicyService services.PolicyService
}

// TODO review use of pointers here
// TODO is this a constructor?
func NewPolicyEndpoint() *PolicyEndpoint {
	policyService = *services.NewPolicyService()
	policyEndpoint := &PolicyEndpoint{policyService}
	return policyEndpoint
}

// the gin context is passed implictly when the router is registered
func (e PolicyEndpoint) GetPolicy(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, policyService.GetPolicy())
}
