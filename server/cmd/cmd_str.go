package cmd

import (
	"../db"
	"fmt"
)
/**
 * @Author: Ember
 * @Date: 2022/5/5 12:51
 * @Description: 字符串命令
 **/

func newWrongNumOfArgsError(cmd string) error{
	return fmt.Errorf("wrong number of '#{cmd}' command...")
}
//添加字符串命令
//set key value
func set(db *db.GoDB,args []string)(res interface{},err error){
	//判断长度
	//如果长度大于2，抛错
	if(len(args) != 2){
		err = newWrongNumOfArgsError("set")
		return
	}
	//记录键值对
	key := args[0]
	value := args[1]
	//执行命令
	//使用[]byte()转化为字节数组
	res,err = db.StrIndex.Set([]byte(key),[]byte(value))
	if err != nil{
		res = "nil"
	}
	return
}

//获取字符串命令
//get key
func get(db *db.GoDB,args []string)(res interface{},err error){
	if(len(args) != 1){
		err = newWrongNumOfArgsError("get")
		return
	}
	key := args[0]
	res,err = db.StrIndex.Get([]byte(key))
	if err != nil{
		res = "nil"
	}
	return
}