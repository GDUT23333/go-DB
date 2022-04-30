package entry

import "time"
/**
 * @Author: Ember
 * @Date: 2022/4/29 23:21
 * @Description: 封装的键值对
 **/

/**
键值对里面的元数据
*/
type Meta struct{
	//键值对，使用字节数组
	Key []byte
	Value []byte

	//键值对的长度
	KeySize int
	ValueSize int
}

type Entry struct{
	//元数据
	Meta *Meta
	//创建时间
	CreateTime int64
}

/**
创建一个新Entry
 */
func CrateEntry(key []byte,value []byte) *Entry{
	//创建Meta
	meta := Meta{
		Key:key,
		Value:value,
	};
	//创建Entry
	entry := Entry{
		Meta: &meta,
		CreateTime: time.Now().Unix(),
	}
	return &entry
}

