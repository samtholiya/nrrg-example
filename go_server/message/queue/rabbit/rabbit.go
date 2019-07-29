package rabbit

import (
	"bytes"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

//RabbitMq struct holding a variables required for sending and receiving messages
type RabbitMq struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   *amqp.Queue
}

//Connect params to connect to the rabbitmq
func (r *RabbitMq) Connect(host, port, username, password string) {
	var buffer bytes.Buffer
	buffer.WriteString("amqp://")
	buffer.WriteString(username)
	buffer.WriteString(":")
	buffer.WriteString(password)
	buffer.WriteString("@")
	buffer.WriteString(host)
	buffer.WriteString(":")
	buffer.WriteString(port)
	buffer.WriteString("/")

	conn, err := amqp.Dial(buffer.String())
	failOnError(err, "Failed to connect to RabbitMQ")
	r.conn = conn
	ch, err := r.conn.Channel()
	failOnError(err, "Failed to open a channel")
	r.channel = ch
}

//Close closes the channel and connection
func (r *RabbitMq) Close() {
	r.channel.Close()
	r.conn.Close()
}

//SetDefaultQueueName creates and sets a default queue in rabbitmq
func (r *RabbitMq) SetDefaultQueueName(name string) {
	q, err := r.channel.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to create queue")
	r.queue = &q
}

//PublishMessage publishes message and returns error if any
func (r *RabbitMq) PublishMessage(message string) error {
	return r.publishMessage(message, r.queue)
}

func (r *RabbitMq) publishMessage(message string, queue *amqp.Queue) error {
	fmt.Println(message)
	return r.channel.Publish(
		"",           // exchange
		r.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

//ReceiveString returns a string in the queue. (Blocking call)
func (r *RabbitMq) ReceiveString() (string, error) {
	receive, err := r.receiveMessageFromQueue(r.queue)
	if err != nil {
		return "", err
	}
	obj := <-receive
	return string(obj.Body), nil
}

func (r *RabbitMq) receiveMessageFromQueue(queue *amqp.Queue) (<-chan amqp.Delivery, error) {
	receive, err := r.channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	return receive, err
}

//PublishMessageToQueue message to the queue specified
func (r *RabbitMq) PublishMessageToQueue(message, queueName string) error {
	q, err := r.channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to create queue")
	return r.channel.Publish(
		"",     // exchange
		q.Name, // routing key
		true,   // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

//ReceiveStringFromQueue Receive from the queue specified get the message in string format (Blocking)
func (r *RabbitMq) ReceiveStringFromQueue(queueName string) (string, error) {
	queue, err := r.channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return "", err
	}
	receive, err := r.receiveMessageFromQueue(&queue)
	obj := <-receive
	return string(obj.Body), nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
