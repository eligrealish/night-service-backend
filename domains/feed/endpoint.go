package feed

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	feedService FeedService
)

type FeedEndpoint struct {
	feedService FeedService
}

func NewFeedEndpoint() *FeedEndpoint {
	feedService = *NewFeedService()
	feedEndpoint := &FeedEndpoint{feedService}
	return feedEndpoint
}

// the gin context is passed implictly when the router is registered
func (f FeedEndpoint) GetFeed(context *gin.Context) {
	context.JSON(http.StatusOK, feedService.GetFeedHandler(context))
}
