package pkg

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type singleton struct {
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = new(singleton)
			fmt.Println("创建单个实例")
		} else {
			fmt.Println("已创建单个实例")
		}
	} else {
		fmt.Println("已创建单个实例")
	}
	return instance
}
