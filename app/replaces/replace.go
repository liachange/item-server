package replaces

import (
	"github.com/spf13/cast"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
	"strings"
)

func BetTime(str string) []int64 {
	betTime := make([]int64, 0)
	if strings.Contains(str, ",") {
		bet := strings.Split(str, ",")
		for _, v := range bet {
			if t := cast.ToInt64(v); t > 0 {
				betTime = append(betTime, t)
			}
		}
	}
	return betTime
}

func IdSlice(keys []uint64) []uint64 {
	ids := make([]uint64, 0)
	if len(keys) > 0 && keys[0] != 0 {
		opt := optimusPkg.NewOptimus()
		for _, v := range keys {
			if v > 0 {
				ids = append(ids, opt.Decode(v))
			}
		}
	}
	return ids
}
func IdString(str string) []uint64 {
	ids := make([]uint64, 0)
	if str != "" {
		opt := optimusPkg.NewOptimus()
		if strings.Contains(str, ",") {
			split := strings.Split(str, ",")
			for _, v := range split {
				if id := cast.ToUint64(v); id > 0 {
					ids = append(ids, opt.Decode(id))
				}
			}
		} else {
			ids = append(ids, opt.Decode(cast.ToUint64(str)))
		}
	}
	return ids
}

func IdPublicReplace(c []uint64) (ids []uint64, isPublic uint8) {
	//分类参数处理
	isPublic = 1
	opt := optimusPkg.NewOptimus()
	keys := make([]uint64, 0)
	if !helpers.Empty(c) && c[0] != 0 {
		for _, v := range c {
			keys = append(keys, opt.Decode(v))
		}
		isPublic = 2
	}
	return keys, isPublic
}
