package _struct

import (
	"bytes"
	"math"
	"math/rand"
	"time"
)

/**
 * @Author: Ember
 * @Date: 2022/4/30 16:07
 * @Description: 跳表，为什么要使用调表呢？
 **/

const(
	//最高层级
	default_maxLevel int = 18
	default_probability float64 = 1/math.E
)
type (
	//底层双向链表的结点
	Level struct{
		//指向下一个结点的层次结构
		next []*Node
	}
	//存储键值对信息
	Node struct{
		key []byte
		//object类型的value
		value interface{}
		//层次结构所属的结点
		//可以找到下一个层次结构
		level Level
	}
	//跳表
	SkipList struct{
		//哨兵傀儡结点
		Sentinel Node
		maxLevel int
		//每一层高度的可能性
		levelProbabilities  []float64
		//高度上升的可能性
		probability float64
		//随机数
		randSource rand.Source
	}
)

//key & value的get方法
func (e *Node) Key() []byte{
	return e.key
}

func (e *Node) Value() interface{}{
	return e.value
}

//key & value的set方法
func (e *Node) SetKey(key []byte){
	e.key = key
}

func (e *Node) SetValue(value interface{}){
	e.value = value
}

//获取该结点的下一个结点
//也就是下一个层级结构的底层
func (e *Node) Next() *Node{
	return e.level.next[0]
}

//创建新的跳表
func CreateSkipList() *SkipList{
	return &SkipList{
		//哨兵结点
		Sentinel: Node{
			key:nil,
			value:nil,
			level : Level{
				next:make([]*Node,default_maxLevel),
			},
		},
		//最高层次
		maxLevel : default_maxLevel,
		//上升一层的可能性
		probability : default_probability,
		//每一层的概率
		levelProbabilities : initialLevelProbability(default_probability,default_maxLevel),
		//随机数种子，用来生成概率的
		randSource : rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

//初始化每一层的概率
func initialLevelProbability(probability float64,maxLevel int)([]float64){
	levelTable := make([]float64,maxLevel)
	for i := 0;i < maxLevel;i++{
		//计算当前层的概率
		curProb := math.Pow(probability, float64(i))
		levelTable = append(levelTable,curProb)
	}
	return levelTable
}

//生成随机概率
func (s *SkipList) RandomLevel() int{
	//获取概率
	r := float64(s.randSource.Int63()) / (1 << 63)
	level := 1
	for level < s.maxLevel && r < s.levelProbabilities[level] {
		level++
	}
	return level
}
//获取哨兵结点
func (s *SkipList) Front() *Node{
	return &s.Sentinel
}

//往跳表中添加结点
//并且返回新添加的结点
func (s *SkipList) Put(key []byte,value interface{}) *Node{
	//第一步：找到前面的一个pre结点
	preNode := s.GetPreNode(key)
	//第二步：判断前结点下一个的key值，如果相等，只需要替换value即可
	nextNode := preNode.level.next[0]
	if nextNode != nil && bytes.Compare(nextNode.key,key) == 0{
		nextNode.value = value
		return nextNode
	}
	//第三步：创建Node，并且决定高度
	//要地址，因为数组中存的也是地址
	curNode := &Node{
		key : key,
		value : value,
		//随机生成自身的层级高度去连接下一层
		level : Level{
			next : make([]*Node,s.RandomLevel()),
		},
	}
	//第四步：当前结点连接上pre结点的下一个结点，pre结点连接上当前结点
	for i := range curNode.level.next{
		curNode.level.next[i] = preNode.level.next[i]
		preNode.level.next[i] = curNode
	}
	return curNode
}

//搜索结点
func (s *SkipList) Get(key []byte) *Node{
	preNode := &s.Sentinel
	var curNode *Node
	for i := s.maxLevel;i >=0 ;i--{
		//当前结点是pre的下一个
		curNode = preNode.level.next[i]
		//遍历，如果当前的curNode小于key，要找下一个
		//还是在当前level寻找
		for curNode != nil && bytes.Compare(key,curNode.key) > 0 {
			preNode = curNode
			curNode = curNode.level.next[i]
		}
	}
	//判断找到的结点是不是值相等
	if curNode != nil && bytes.Compare(curNode.key,key) == 0{
		return curNode
	}
	return nil
}
func (s *SkipList) GetPreNode(key []byte) *Node{
	//前一个结点
	preNode := &s.Sentinel
	var curNode *Node
	//一开始从最高层出发搜索
	//使用遍历
	for i := s.maxLevel;i >= 0;i--{
		//找到当前层的所属结点
		curNode = preNode.level.next[i]
		//继续循环去寻找，直到
		//如果当前结点不为Null，并且当前的key大于当前结点的key
		//继续找下一层，从小到大进行排序
		for curNode != nil && bytes.Compare(key,curNode.key) > 0 {
			preNode = curNode
			//迭代更新curNode，
			//找下一个层级的结点
			curNode = curNode.level.next[i]
		}
		//如果出现等于或者大于，找preNode的下一个结点的，下一层
	}
	return preNode
}


