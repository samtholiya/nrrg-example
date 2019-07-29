package student

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samtholiya/tempServer/config"
)

//Route collection of all student routes to be registered on gin router
type Route struct {
	cache          config.CacheDatabase
	messagingQueue config.MessagingQueue
	defaultQueue   string
}

//SetMessagingQueue set the messaging queue to be used
func (r *Route) SetMessagingQueue(messagingQueue config.MessagingQueue, defaultQueue string) {
	r.messagingQueue = messagingQueue
	r.defaultQueue = defaultQueue
}

//SetCacheDatabase set the caching database
func (r *Route) SetCacheDatabase(cache config.CacheDatabase) {
	r.cache = cache
}

//Register registers
func (r *Route) Register(router *gin.RouterGroup) {
	router.POST("/", r.postStudentData)
}

func (r *Route) postStudentData(c *gin.Context) {
	id, _ := uuid.NewRandom()
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, _ := json.Marshal(student)
	r.cache.Set(id.String(), string(data))
	if err := r.messagingQueue.PublishMessageToQueue(id.String(), r.defaultQueue); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(201, gin.H{
		"message": "successfully saved to redis",
	})
}
