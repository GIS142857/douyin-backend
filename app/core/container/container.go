package container

import (
	"douyin-backend/app/global/my_errors"
	"douyin-backend/app/global/variable"
	"log"
	"sync"
)

// 定义一个全局键值对存储容器
var sMap sync.Map

type Containers struct {
}

func CreateContainersFactory() *Containers {
	return &Containers{}
}

func (c *Containers) Set(key string, value interface{}) (res bool) {
	if _, exists := c.KeyIsExists(key); exists == false {
		sMap.Store(key, value)
		return true
	} else {
		if variable.ZapLog == nil {
			log.Fatal(my_errors.ErrorsContainerKeyAlreadyExists + ",请解决键名重复问题,相关键：" + key)
		} else {
			variable.ZapLog.Warn(my_errors.ErrorsContainerKeyAlreadyExists + ", 相关键：" + key)
		}
	}
	return
}

func (c *Containers) Get(key string) interface{} {
	if value, exists := c.KeyIsExists(key); exists {
		return value
	}
	return nil
}

func (c *Containers) KeyIsExists(key string) (interface{}, bool) {
	return sMap.Load(key)
}
