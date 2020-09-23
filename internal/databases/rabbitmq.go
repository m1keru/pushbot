package databases

import (
	"encoding/json"
	"time"

	"github.com/m1keru/pushbot/internal/interfaces"
	"github.com/m1keru/pushbot/internal/logging"
	"github.com/streadway/amqp"
)

// RabbitMQ -- databases
type RabbitMQ struct {
	Amqp         string `yaml:"amqp"`
	FailTimeout  int    `yaml:"fail_timeout"`
	Schema       string `yaml:"schema,omitempty"`
	ExchangeName string `yaml:"exchange_name"`
	QueueName    string `yaml:"queue_name"`
	Channel      *amqp.Channel
	Queue        *amqp.Queue
}

var conn *amqp.Connection
var ch *amqp.Channel

// Init -- init
func (db *RabbitMQ) Init() *RabbitMQ {
	var err error

	conn, err = amqp.Dial(db.Amqp)
	logging.CheckError("Failed to connect to RabbitMQ\n", err)

	db.Channel, err = conn.Channel()
	logging.CheckError("Failed to open a channel", err)
	queue, err := db.Channel.QueueDeclare(
		db.QueueName, // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	db.Queue = &queue
	logging.CheckError("Failed to declare queue", err)

	err = db.Channel.ExchangeDeclare(
		db.ExchangeName,
		"direct",
		true,  // durable
		false, // auto-deleted
		false, // internal
		false, // noWait
		nil,   // arguments
	)
	logging.CheckError("Failed to declare the Exchange", err)

	err = db.Channel.QueueBind(
		db.QueueName,    // queue name
		db.QueueName,    // routing key
		db.ExchangeName, // exchange
		false,
		nil)

	logging.CheckError("Failed to bind queue to Exchange", err)

	logging.Log("rabbitmq: connection done")
	return db
}

//Publish -- Publish
func (db *RabbitMQ) Publish(msg interfaces.HookMessage) {
	payload, err := json.Marshal(msg)
	logging.CheckError("Failed to Marshall", err)
	err = db.Channel.Publish(
		db.ExchangeName, // exchange
		db.QueueName,    // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Transient,
			ContentType:  "application/json",
			Body:         payload,
			Timestamp:    time.Now(),
		})

	logging.CheckError("Failed to declare the Exchange", err)
	logging.Logf("rabbitmq: message: {{{ %+v }}} published: OK", msg)
}

// Consume -- Consume
func (db *RabbitMQ) Consume() {
	msgs, err := db.Channel.Consume(
		db.QueueName, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	logging.CheckError("Failed to register a consumer", err)

	go func() {
		for d := range msgs {
			logging.Logf("Received a message: %s", d.Body)
		}
	}()
}
