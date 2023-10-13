package optimus

import (
	"github.com/pjebs/optimus-go"
	"item-server/pkg/config"
	"sync"
)

type Optimus struct {
	Optimus optimus.Optimus
}

// once 确保 internalOptimus 对象只初始化一次
var once sync.Once

// internalOptimus 内部使用的 Optimus 对象
var internalOptimus Optimus

func NewOptimus() Optimus {
	once.Do(func() {
		// 初始化 Optimus 对象
		internalOptimus = Optimus{}
		internalOptimus.Optimus = optimus.New(config.GetUint64("optimus.prime"), config.GetUint64("optimus.inverse"), config.GetUint64("optimus.random"))
	})
	return internalOptimus

}

// Encode 混淆
func (o Optimus) Encode(n uint64) uint64 {
	return o.Optimus.Encode(n)
}

// Decode 解码
func (o Optimus) Decode(n uint64) uint64 {
	return o.Optimus.Decode(n)
}
