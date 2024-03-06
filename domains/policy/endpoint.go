package policy

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	policyService PolicyService
)

type PolicyEndpoint struct {
	policyService PolicyService
}

func NewPolicyEndpoint() *PolicyEndpoint {
	policyService = *NewPolicyService()
	policyEndpoint := &PolicyEndpoint{policyService}
	return policyEndpoint
}

// the gin context is passed implictly when the router is registered
func (e PolicyEndpoint) GetPolicy(context *gin.Context) {
	context.JSON(http.StatusOK, policyService.GetPolicyHandler())
}
