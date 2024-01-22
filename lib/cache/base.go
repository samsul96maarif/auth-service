/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2023-11-02 20:39:34
 * @modify date 2023-11-02 20:39:34
 * @desc [description]
 */

package cache

import "time"

type CacheUtil interface {
	Close() error
	GetFromCache(key string) (string, error)
	SetToCache(key string, data interface{}, expires time.Duration) error
}
