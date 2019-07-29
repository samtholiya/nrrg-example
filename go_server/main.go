package main

import (
	"os"

	"github.com/samtholiya/tempServer/route"

	"github.com/gin-gonic/gin"
	"github.com/samtholiya/tempServer/cache"
	"github.com/samtholiya/tempServer/message/queue"
)

func main() {
	r := gin.Default()

	messaging := queue.New()
	messaging.Connect(os.Getenv("RABBITMQ_HOST"), os.Getenv("RABBITMQ_PORT"), os.Getenv("RABBITMQ_USERNAME"), os.Getenv("RABBITMQ_PASSWORD"))

	redisClient := cache.New()
	redisClient.SetConfig(os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PASSWORD"))

	messaging.SetDefaultQueueName("db")

	v1 := r.Group("/api/v1")
	student := route.GetStudentRoute()
	student.SetMessagingQueue(messaging, "db.student")
	student.SetCacheDatabase(redisClient)
	student.Register(v1.Group("/student"))
	r.Run()
}
