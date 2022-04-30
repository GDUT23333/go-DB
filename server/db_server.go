package server

import "./index"
import "./config"
/**
 * @Author: Ember
 * @Date: 2022/4/29 23:15
 * @Description: 服务器信息
 **/

type DBServer struct{
	//底层四种数据结构的索引
	strIndex index.StrIndex;
	listIndex index.ListIndex;
	hashIndex index.HashIndex;
	setIndex index.SetIndex;
	//配置
	config config.Config;
	//
}