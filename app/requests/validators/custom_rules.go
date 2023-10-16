package validators

import (
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"github.com/thedevsaddam/govalidator"
	"item-server/pkg/database"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
	"strconv"
	"strings"
	"unicode/utf8"
)

// 此方法会在初始化时执行，注册自定义表单验证规则
func init() {

	// 自定义规则 not_exists，验证请求数据必须不存在于数据库中。
	// 常用于保证数据库某个字段的值唯一，如用户名、邮箱、手机号、或者分类的名称。
	// not_exists 参数可以有两种，一种是 2 个参数，一种是 3 个参数：
	// not_exists:users,email 检查数据库表里是否存在同一条信息
	// not_exists:users,email,32 排除用户掉 id 为 32 的用户
	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

		// 第一个参数，表名称，如 users
		tableName := rng[0]
		// 第二个参数，字段名称，如 email 或者 phone
		dbFiled := rng[1]

		// 第三个参数，排除 ID
		var exceptID uint64
		if len(rng) > 2 {
			id := cast.ToUint64(rng[2])
			if id > 0 {
				exceptID = optimusPkg.NewOptimus().Decode(id)
			} else {
				return errors.New(message)
			}
		}

		// 用户请求过来的数据
		requestValue := value.(string)

		// 拼接 SQL
		query := database.DB.Table(tableName).Where(dbFiled+" = ?", requestValue)

		// 如果传参第三个参数，加上 SQL Where 过滤
		if exceptID > 0 {
			query.Where("id != ?", exceptID)
		}

		// 查询数据库
		var count int64
		query.Count(&count)

		// 验证不通过，数据库能找到对应的数据
		if count != 0 {
			// 如果有自定义错误消息的话
			if message != "" {
				return errors.New(message)
			}
			// 默认的错误消息
			return fmt.Errorf("%v 已被占用", requestValue)
		}
		// 验证通过
		return nil
	})

	// max_cn:8 中文长度设定不超过 8
	govalidator.AddCustomRule("max_cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_cn:"))
		if valLength > l {
			// 如果有自定义错误消息的话，使用自定义消息
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度不能超过 %d 个字", l)
		}
		return nil
	})

	// min_cn:2 中文长度设定不小于 2
	govalidator.AddCustomRule("min_cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "min_cn:"))
		if valLength < l {
			// 如果有自定义错误消息的话，使用自定义消息
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度需大于 %d 个字", l)
		}
		return nil
	})

	// 自定义规则 exists，确保数据库存在某条数据
	// 一个使用场景是创建话题时需要附带 category_id 分类 ID 为参数，此时需要保证
	// category_id 的值在数据库中存在，即可使用：
	// exists:categories,id
	govalidator.AddCustomRule("exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "exists:"), ",")

		// 第一个参数，表名称，如 categories
		tableName := rng[0]
		// 第二个参数，字段名称，如 id
		dbFiled := rng[1]
		// 用户请求过来的数据
		var requestValue any
		if dbFiled == "id" {
			requestValue = optimusPkg.NewOptimus().Decode(value.(uint64))
		} else {
			requestValue = value.(string)
		}

		// 查询数据库
		var count int64
		database.DB.Table(tableName).Where(dbFiled+" = ?", requestValue).Count(&count)
		// 验证不通过，数据不存在
		if count == 0 {
			// 如果有自定义错误消息的话，使用自定义消息
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%v 不存在", requestValue)
		}
		return nil
	})
	// 自定义规则 exist_key，确保数据库存在某条数据
	// id 的值在数据库中存在，即可使用：
	// exist_key:categories
	govalidator.AddCustomRule("exist_key", func(field string, rule string, message string, value interface{}) error {
		// 第一个参数，表名称，如 categories
		tableName := strings.TrimPrefix(rule, "exist_key:")

		// 用户请求过来的数据
		requestValue := value.([]uint64)
		if len(requestValue) == 0 {
			return errors.New(message)
		}
		ids := make([]uint64, 0)
		opt := optimusPkg.NewOptimus()
		for _, v := range requestValue {
			if v > 0 {
				ids = append(ids, opt.Decode(v))
			}
		}
		if len(ids) != len(requestValue) {
			return errors.New(message)
		}
		// 查询数据库
		var count int64
		database.DB.Table(tableName).Where("id in ?", ids).Count(&count)
		// 验证不通过，数据不存在
		if cast.ToInt(count) != len(ids) {
			// 如果有自定义错误消息的话，使用自定义消息
			return errors.New(message)
		}
		return nil
	})
	// 自定义规则 slice_time，确保该字段为切片并且值为时间戳
	govalidator.AddCustomRule("slice_time", func(field string, rule string, message string, value interface{}) error {
		valSlice := strings.Split(value.(string), ",")
		if len(valSlice) == 2 {
			for _, v := range valSlice {
				if len(v) != 10 {
					return errors.New(message)
				}
				str := helpers.TimeStr(cast.ToInt64(v), "second")
				if strings.Contains(str, "1970-01") {
					return errors.New(message)
				}
			}
		} else {
			return errors.New(message)
		}
		return nil
	})
	// 自定义规则 required_without:foo,bar，确保该字段为切片并且值为时间戳
	govalidator.AddCustomRule("required_without", func(field string, rule string, message string, value interface{}) error {
		valSlice := strings.Split(value.(string), ",")
		if len(valSlice) == 2 {
			for _, v := range valSlice {
				if len(v) != 10 {
					return errors.New(message)
				}
				str := helpers.TimeStr(cast.ToInt64(v), "second")
				if strings.Contains(str, "1970-01") {
					return errors.New(message)
				}
			}
		} else {
			return errors.New(message)
		}
		return nil
	})
	govalidator.AddCustomRule("exist_user", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "exist_user:"), ",")

		// 第一个参数，表名称，如 categories
		tableName := rng[0]
		// 用户请求过来的数据
		requestValue := value.(string)
		query := database.DB.Table(tableName)
		if len(rng) == 2 {
			id := cast.ToUint64(rng[1])
			if id > 0 {
				query.Where("id <> ?", optimusPkg.NewOptimus().Decode(id)).Where(
					database.DB.Or("phone = ?", requestValue).
						Or("email = ?", requestValue).
						Or("name = ?", requestValue),
				)
			} else {
				return errors.New(message)
			}
		} else {
			query.Where("phone = ?", requestValue).
				Or("email = ?", requestValue).
				Or("name = ?", requestValue)
		}
		// 查询数据库
		var count int64
		query.Count(&count)
		// 数据存在
		if count != 0 {
			return errors.New(message)
		}
		return nil
	})
}
