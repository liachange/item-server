package filter

import (
	"gorm.io/gorm"
	"item-server/pkg/helpers"
	"strings"
)

// QueryFilter 查询过滤 eq等于 lt小于 gt大于 neq不等于 egt大于等于 elt小于等于 bet 两个值之间 in 包含在给定的值列表中
func QueryFilter(query *gorm.DB, filters map[string]interface{}) *gorm.DB {

	for k, v := range filters {
		rng := strings.Split(k, ",")
		switch rng[1] {
		case "in":
			query.Where(rng[0]+" in ?", v)
			break
		case "like":
			query.Where(rng[0]+" like ?", "%"+v.(string)+"%")
			break
		case "egt":
			query.Where(rng[0]+" >= ?", v)
			break
		case "elt":
			query.Where(rng[0]+" <= ?", v)
			break
		case "bet_time":
			bet := v.([]int64)
			query.Where(rng[0]+" between ? and ?", helpers.TimeStr(bet[0], "second"), helpers.TimeStr(bet[1], "second"))
			break
		case "neq":
			query.Where(rng[0]+" <> ?", v)
			break
		case "eq":
			query.Where(rng[0]+" = ?", v)
			break
		case "lt":
			query.Where(rng[0]+" < ?", v)
			break
		case "gt":
			query.Where(rng[0]+" > ?", v)
			break
		default:
			query.Where("id =?", "0")
		}
	}
	return query
}
