package index

import (
	"./struct"
	"sync"
)

/**
 * @Author: Ember
 * @Date: 2022/4/29 23:37
 * @Description: 字符串索引
 **/

type StrIndex struct{
	//底层的跳表
	skiplist _struct.SkipList
	//全局锁，读写锁
	mu *sync.RWMutex
}
