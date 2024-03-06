package utils

import (
	"github.com/gin-gonic/gin"
)

type Utils struct {
}

func NewUtils() *Utils {
	instance := &Utils{}
	return instance
}

func checkRequestKeyPresent(context *gin.Context, paramName string) (string, bool) {
	value, ok := context.GetQuery(paramName)
	return value, ok && value != ""
}
