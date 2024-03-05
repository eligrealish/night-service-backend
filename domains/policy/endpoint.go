package policy

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	policyService PolicyService
)

type PolicyEndpoint struct {
	PolicyService PolicyService
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

func (e PolicyEndpoint) GetTest(context *gin.Context) {
	context.JSON(http.StatusOK, "\"{status\":\"test\"}")
}
