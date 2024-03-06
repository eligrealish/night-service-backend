package utils

import (
	"github.com/gin-gonic/gin"
)

type RequestUtils struct {
}

func RequestUtilsNewUtils() *RequestUtils {
	instance := &RequestUtils{}
	return instance
}

func (u RequestUtils) CheckRequestKeyPresent(context *gin.Context, paramName string) bool {
	value, ok := context.GetQuery(paramName)
	return ok && value != ""
}
