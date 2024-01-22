/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2023-01-29 15:32:47
 * @modify date 2023-01-29 15:32:47
 * @desc [description]
 */
package config

import (
	"context"
	"fmt"
	"os"

	"github.com/samsul96maarif/auth-service/lib/queue"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	WorkerTopic       = "WORKER-TOPIC"
	ImporterTopic     = "IMPORTER"
	KeyWorkerImporter = "WORKER-IMPORTER"
)

type ConsumerConfig struct {
	Topic       string
	PollTimeout int
}

func NewConsumer(cfg ConsumerConfig) (qc queue.Consumer, err error) {
	groupID := fmt.Sprintf("app.%s.%s", os.Getenv("ENV"), cfg.Topic)
	var c *kafka.Consumer
	protocol := os.Getenv("KAFKA_SECURITY_PROTOCOL")
	if protocol != "" {
		c, err = kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers":                   os.Getenv("KAFKA_BROKERS"),
			"security.protocol":                   protocol,
			"sasl.mechanisms":                     "PLAIN",
			"sasl.username":                       os.Getenv("KAFKA_USERNAME"),
			"sasl.password":                       os.Getenv("KAFKA_PASSWORD"),
			"enable.ssl.certificate.verification": "false",
			"enable.auto.commit":                  true,
			"auto.offset.reset":                   "latest",
			"max.poll.interval.ms":                cfg.PollTimeout,
			"auto.commit.interval.ms":             2000,
			"group.id":                            groupID,
			"enable.auto.offset.store":            false,
		})
	} else {
		c, err = kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers":        os.Getenv("KAFKA_BROKERS"),
			"enable.auto.commit":       true,
			"auto.offset.reset":        "latest",
			"max.poll.interval.ms":     cfg.PollTimeout,
			"auto.commit.interval.ms":  2000,
			"group.id":                 groupID,
			"enable.auto.offset.store": false,
		})
	}

	if err == nil {
		qc = &queue.QueueConsumer{
			Consumer:       c,
			Topic:          cfg.Topic,
			WorkerHandlers: make(map[string]func(ctx context.Context, payload []byte, key []byte) error),
		}
	}
	return
}

func NewPublisher() (q queue.Publisher, err error) {
	var producer *kafka.Producer
	protocol := os.Getenv("KAFKA_SECURITY_PROTOCOL")
	if protocol != "" {
		producer, err = kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers":                   os.Getenv("KAFKA_BROKERS"),
			"acks":                                "all",
			"security.protocol":                   protocol,
			"sasl.mechanisms":                     "PLAIN",
			"sasl.username":                       os.Getenv("KAFKA_USERNAME"),
			"sasl.password":                       os.Getenv("KAFKA_PASSWORD"),
			"enable.ssl.certificate.verification": "false",
		})
	} else {
		producer, err = kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": os.Getenv("KAFKA_BROKERS"),
			"acks":              "all",
		})
	}
	if err != nil {
		return nil, err
	}
	kafkaPublisher := queue.QueuePublisher{producer}
	q = &kafkaPublisher
	return
}
