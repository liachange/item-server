package config

import "item-server/pkg/config"

func init() {
	config.Add("optimus", func() map[string]interface{} {
		return map[string]interface{}{
			// 素数
			"prime": config.Env("OPTIMUS_PRIME", 938860369),

			// 逆数
			"inverse": config.Env("OPTIMUS_INVERSE", 294308273),

			// 随机数
			"random": config.Env("OPTIMUS_RANDOM", 499995894),
		}
	})
}
