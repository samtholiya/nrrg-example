package route

import (
	"github.com/gin-gonic/gin"
	"github.com/samtholiya/tempServer/config"
	"github.com/samtholiya/tempServer/route/student"
)

//Route every route should implement this
type Route interface {

	//SetMessagingQueue set the messaging queue to be used
	SetMessagingQueue(messagingQueue config.MessagingQueue, defaultQueue string)

	//SetCacheDatabase set the caching database
	SetCacheDatabase(cache config.CacheDatabase)

	//Register registers
	Register(router *gin.RouterGroup)
}

//GetStudentRoute get the route module for Student
func GetStudentRoute() Route {
	return &student.Route{}
}
