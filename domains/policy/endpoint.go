package policy

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// TODO reviewing threading of local scope vs global e.g should i use VAR
var (
	policyService PolicyService
)

type PolicyEndpoint struct {
	PolicyService PolicyService
}

// TODO review use of pointers here
// TODO is this a constructor?
func NewPolicyEndpoint() *PolicyEndpoint {
	policyService = *NewPolicyService()
	policyEndpoint := &PolicyEndpoint{policyService}
	return policyEndpoint
}

// the gin context is passed implictly when the router is registered
func (e PolicyEndpoint) GetPolicy(context *gin.Context) {
	context.JSON(http.StatusOK, policyService.GetPolicy())
}

func (e PolicyEndpoint) GetTest(context *gin.Context) {
	context.JSON(http.StatusOK, "\"{status\":\"test\"}")
}
