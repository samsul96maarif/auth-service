/*
 * Author: Samsul Ma'arif <samsulma828@gmail.com>
 * Copyright (c) 2023.
 */

package queue

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/samsul96maarif/auth-service/lib/logger"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type QueueConsumer struct {
	Consumer       *kafka.Consumer
	WorkerHandlers map[string]func(ctx context.Context, payload []byte, key []byte) error
	Topic          string
	PollTimeout    int
}

func (q *QueueConsumer) RegisterWorker(key string, handler func(ctx context.Context, payload, key []byte) error) {
	q.WorkerHandlers[key] = handler
}

func (q *QueueConsumer) Start(ctx context.Context) error {
	err := q.Consumer.Subscribe(q.Topic, nil)
	if err != nil {
		logger.Error(context.Background(), "error Start consumer", map[string]interface{}{
			"error": err.Error(),
		})
		return err
	}
	logger.Info(context.Background(), "consumer activated topic: "+q.Topic, nil)

	run := true
	for run {
		select {
		default:
			event := q.Consumer.Poll(q.PollTimeout)
			if event == nil {
				continue
			}
			ctx := context.Background()
			switch ev := event.(type) {
			case *kafka.Message:
				logger.Info(ctx, fmt.Sprintf("Processing message"), map[string]interface{}{
					"tags": []string{"kafka", "worker", "process_message", q.Topic},
				})
				handler, ok := q.WorkerHandlers[string(ev.Key)]
				if ok {
					err = handler(ctx, ev.Value, ev.Key)
					if err != nil {
						logger.Error(ctx, err.Error(), map[string]interface{}{
							"error": err,
							"tags":  []string{"kafka", "worker", "store_offset", q.Topic},
						})
						run = false
					}
				}
			case kafka.Error:
				if ev.Code() == kafka.ErrAllBrokersDown {
					logger.Error(ctx, ev.Error(), map[string]interface{}{
						"error": err,
						"tags":  []string{"kafka", "worker", q.Topic},
					})
					run = false
				}
			}
		}
	}

	return nil
}

func (q *QueueConsumer) Subscribe(message chan<- string) error {
	err := q.Consumer.Subscribe(q.Topic, nil)
	if err != nil {
		logger.Error(context.Background(), "error Subscribe", map[string]interface{}{
			"error": err.Error(),
		})
		return err
	}
	go func() {
		for {
			event := q.Consumer.Poll(q.PollTimeout)
			if event == nil {
				continue
			}
			ctx := context.Background()
			switch ev := event.(type) {
			case *kafka.Message:
				logger.Info(ctx, fmt.Sprintf("Processing message = ", string(ev.Value)), map[string]interface{}{
					"tags": []string{"kafka", "consumer", "process_message", q.Topic},
				})
				message <- string(ev.Value) // ev.Key
			case kafka.Error:
				if ev.Code() == kafka.ErrAllBrokersDown {
					logger.Error(ctx, ev.Error(), map[string]interface{}{
						"error": err,
						"tags":  []string{"kafka", " consumer", q.Topic},
					})
				}
			}
		}
	}()
	return nil
}

type QueuePublisher struct {
	Client *kafka.Producer
}

func (q *QueuePublisher) SyncProduce(message, routing_key interface{}, topic string) error {

	var key []byte
	if routing_key == nil {
		key = nil
	} else {
		key = []byte(fmt.Sprintf("%v", routing_key))
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	deliveryChan := make(chan kafka.Event)
	err = q.Client.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: messageBytes,
		Key:   key,
	}, deliveryChan)

	if err != nil {
		fmt.Println("error produce", err)
	}

	event := <-deliveryChan
	msg := event.(*kafka.Message)
	if msg.TopicPartition.Error != nil {
		fmt.Println("error msg.TopicPartition", msg.TopicPartition.Error)
		return msg.TopicPartition.Error
	}
	close(deliveryChan)
	return nil
}
