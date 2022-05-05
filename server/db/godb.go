package db

import (
	index "../index"

	config "../config"
	"sync"
)
/**
 * @Author: Ember
 * @Date: 2022/5/3 16:56
 * @Description: 数据库实例
 **/

type GoDB struct{
	//索引
	StrIndex *index.StrIndex
	ListIndex *index.ListIndex
	HashIndex *index.HashIndex
	SetIndex *index.SetIndex
	//只读配置，所以不需要指针
	Config config.Config
	//全局锁
	gbLock *sync.RWMutex

}

//启动数据库
func StartDB(config config.Config)(*GoDB,error){
	//加载持久化文件进入内存

	//创建GoDB结构体
	godb := &GoDB{
		StrIndex : index.CreateStrIndex(),
		ListIndex : index.CreateListIndex(),
		HashIndex : index.CreateHashIndex(),
		SetIndex : index.CreateSetIndex(),
		Config : config,
		//创建锁
		gbLock : new(sync.RWMutex),
	}
	//开启协程进行定时Compacting
	return godb,nil
}