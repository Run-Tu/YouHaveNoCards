package filter

import "dqq/go/basic/type_func/oip/common"

type Filter interface {
	Filter([]*common.Product) []*common.Product //传入一批商品，返回过滤之后的商品
	Name() string
}
