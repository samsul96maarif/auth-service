/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2023-11-03 16:00:51
 * @modify date 2023-11-03 16:00:51
 * @desc [description]
 */

package request

import "time"

type Cache struct {
	Data    interface{}   `json:"data"`
	Expires time.Duration `json:"expires"`
	Key     string        `json:"key"`
}
