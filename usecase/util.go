/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2024-01-20 11:47:27
 * @modify date 2024-01-20 11:47:27
 * @desc [description]
 */

package usecase

import (
	"github.com/samsul96maarif/auth-service/lib/queue"
	"github.com/samsul96maarif/auth-service/lib/worker_const"
)

type UsecaseUtil interface {
	DispatchWorker(routing_key string, message interface{}) error
}

type util struct {
	publisher queue.Publisher
}

func NewUsecaseUtil(publisher queue.Publisher) UsecaseUtil { return &util{publisher} }

func (u *util) DispatchWorker(routing_key string, message interface{}) error {
	return u.publisher.SyncProduce(message, routing_key, worker_const.WorkerTopic)
}
