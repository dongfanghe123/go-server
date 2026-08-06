package dao

import (
	"sync"

	"github.com/go-juicedev/juice"
)

var (
	engine *juice.Engine
	once   sync.Once
)

// Init 初始化数据库引擎（整个应用只调用一次）
func Init(configPath string) error {
	var err error
	once.Do(func() {
		cfg, err := juice.NewXMLConfiguration(configPath)
		if err != nil {
			return
		}
		engine, err = juice.Default(cfg)
	})
	return err
}

// GetEngine 获取引擎
func GetEngine() *juice.Engine {
	return engine
}

// Close 关闭（程序退出时调用）
func Close() {
	if engine != nil {
		engine.Close()
	}
}
