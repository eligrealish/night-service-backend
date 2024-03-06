package feed

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	feedService FeedService
)

type FeedEndpoint struct {
	FeedService FeedService
}

func NewFeedEndpoint() *FeedEndpoint {
	feedService = *NewFeedService()
	feedEndpoint := &FeedEndpoint{feedService}
	return feedEndpoint
}

func (f FolicyEndpoint) GetPolicy(context *gin.Context) {
	context.JSON(http.StatusOK, feedService.GetFeedByLocation())
}
