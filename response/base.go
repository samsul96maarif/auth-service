/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2023-09-03 15:44:54
 * @modify date 2023-09-03 15:44:54
 * @desc [description]
 */

package response

type paginateResponse struct {
	Total int64 `json:"total"`
	Page  uint  `json:"page"`
	Limit uint  `json:"limit"`
}

type PaginateResponse interface {
	GetLimit() uint
	GetPage() uint
	GetTotal() int64
	GetData() interface{}
}

func (r *paginateResponse) GetLimit() uint  { return r.Limit }
func (r *paginateResponse) GetPage() uint   { return r.Page }
func (r *paginateResponse) GetTotal() int64 { return r.Total }
