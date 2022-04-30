package config

/**
 * @Author: Ember
 * @Date: 2022/4/29 23:21
 * @Description: 配置
 **/
import "time"
type Config struct{
	//最大Key大小
	MaxKeySize int
	//最大Value大侠
	MaxValueSize int
	//根目录
	Path string
	//Compaction间隔
	Compaction time.Duration

}