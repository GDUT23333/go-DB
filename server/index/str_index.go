package index

import (
	"./struct"
	"../db"
	"sync"
)

/**
 * @Author: Ember
 * @Date: 2022/4/29 23:37
 * @Description: 字符串索引
 **/

type StrIndex struct{
	//底层的跳表
	skiplist *_struct.SkipList
	//全局锁，读写锁
	mu *sync.RWMutex
}

//创建字符串索引
func CreateStrIndex() *StrIndex{
	//创建StrIndex
	return &StrIndex{
		mu : new(sync.RWMutex),
		skiplist : _struct.CreateSkipList(),
	}
}

//添加字符串对象
func (strIndex *StrIndex)Set(key []byte,value []byte)(res interface{},err error){
	//进行编码
	strIndex.doSet(key,value)
	return "ok",nil
}

func (strIndex *StrIndex)doSet(key []byte,value []byte)(err error){
	//上锁
	strIndex.mu.Lock()
	defer strIndex.mu.Unlock()
	//保存在Index处
	strIndex.skiplist.Put(key,value)
	return
}

//查找字符串对象
func (strIndex *StrIndex) Get(key[] byte)(res interface{},err error){
	res = strIndex.doGet(key)
	return res,nil
}
func (strIndex *StrIndex) doGet(key []byte)(value interface{}){
	//上锁
	strIndex.mu.Lock()
	strIndex.mu.Unlock()
	//从StrIndex处进行搜索
	value = strIndex.skiplist.Get(key).Value()
	return
}
