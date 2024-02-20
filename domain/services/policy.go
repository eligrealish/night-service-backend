package services

import "fmt"

//TODO hardcode and don't use model for testing

type PolicyService struct {
}

func NewPolicyService() *PolicyService {
	instance := &PolicyService{}
	return instance
}

func (s PolicyService) GetPolicy() {
	fmt.Println("logging Policy Instance")
}
