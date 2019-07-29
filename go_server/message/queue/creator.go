package queue

import "github.com/samtholiya/tempServer/message/queue/rabbit"

//MessagingQueue Every messaging queue should have
type MessagingQueue interface {
	//Connect connect to messaging queue using host, port, username and password
	Connect(host, port, username, password string)

	//Close close the connection
	Close()

	//SetDefaultQueueName a default queue with the speceified queue
	SetDefaultQueueName(name string)

	//Publish message to the queue created
	PublishMessage(message string) error

	//ReceiveString get the message in string format (Blocking)
	ReceiveString() (string, error)

	//PublishMessageToQueue message to the queue specified
	PublishMessageToQueue(message, queueName string) error

	//ReceiveStringFromQueue Receive from the queue specified get the message in string format (Blocking)
	ReceiveStringFromQueue(queueName string) (string, error)
}

//New returns and instance of rabbitMq
func New() MessagingQueue {
	return &rabbit.RabbitMq{}
}
