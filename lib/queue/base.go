/*
 * Author: Samsul Ma'arif <samsulma828@gmail.com>
 * Copyright (c) 2023.
 */

package queue

import "context"

type Consumer interface {
	RegisterWorker(key string, handler func(ctx context.Context, payload, key []byte) error)
	Start(ctx context.Context) error
	Subscribe(message chan<- string) error
}

type Publisher interface {
	SyncProduce(message, routing_key interface{}, topic string) error
}
