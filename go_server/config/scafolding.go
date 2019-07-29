package config

//CacheDatabase every cache database should implement
type CacheDatabase interface {

	//Set sets the string value in redis
	Set(key, value string) error

	//Get gets the string corresponding to the key provided
	Get(key string) (string, error)
}

//MessagingQueue Every messaging queue should have
type MessagingQueue interface {

	//PublishMessageToQueue message to the queue specified
	PublishMessageToQueue(message, queueName string) error

	//ReceiveStringFromQueue Receive from the queue specified get the message in string format (Blocking)
	ReceiveStringFromQueue(queueName string) (string, error)
}
