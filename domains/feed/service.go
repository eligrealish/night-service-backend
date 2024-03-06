package feed

import (
	"github.com/gin-gonic/gin"
	"log"
	"night-service-backend/utils"
)

var (
	requestUtils utils.RequestUtils
)

type FeedService struct {
}

func NewFeedService() *FeedService {
	feedService := &FeedService{}
	requestUtils = *utils.RequestUtilsNewUtils()
	return feedService
}

func (e FeedService) GetFeedHandler(context *gin.Context) Feed {
	log.Println("Event Handler")
	// Parse query string parameters
	if requestUtils.CheckRequestKeyPresent(context, "startDate") && requestUtils.CheckRequestKeyPresent(context, "endDate") {

		log.Println("dates passed")
		return Feed{}
	}
	log.Println("dates not passed")
	return Feed{}
}
func GetFeedByLocation() {

}

func GetFeedByLocationAndDate() {

}
