package server

import (
	"./db"
	"./config"
	"sync"
)
/**
 * @Author: Ember
 * @Date: 2022/5/5 13:38
 * @Description: TODO
 **/

type Server struct{
	DB *db.GoDB
	mu sync.Mutex
}

//创建服务器
func StartServer(config config.Config)(server *Server,err error){
	db,err := db.StartDB(config)
	if err != nil{
		server = nil
		return
	}

	lock := new(sync.Mutex)
	server = &Server{
		DB : db,
		mu : *lock,
	}
	return
}