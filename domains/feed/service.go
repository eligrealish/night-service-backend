package feed

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"night-service-backend/utils"
	"time"
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
	location, err := context.GetQuery("location")
	if err {
		log.Println(err)
		// should write 500 here
	}
	// Parse query string parameters
	if requestUtils.CheckRequestKeyPresent(context, "startDate") && requestUtils.CheckRequestKeyPresent(context, "endDate") {
		startDate, _ := context.GetQuery("startDate")
		endDate, _ := context.GetQuery("endDate")
		GetFeedByLocationAndDate(startDate, endDate, location)
		log.Println("dates passed")
		return Feed{}
	}
	log.Println("dates not passed")
	GetFeedByLocation(location)
	return Feed{}
}
func GetFeedByLocation(location string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("Night_Service").Collection("Event")
}

func GetFeedByLocationAndDate(startDate string, endDate string, location string) {

}
