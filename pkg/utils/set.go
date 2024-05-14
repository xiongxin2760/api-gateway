package utils

import (
	"reflect"
)

// Set 非线程安全 map[interface{}]struct{}
type Set map[interface{}]struct{}

// NewSet 创建一个set
func NewSet() Set {
	return make(Set)
}

// Has 判断set中是否有这个key存在
func (s Set) Has(key interface{}) bool {
	_, ok := s[key]
	return ok
}

// Add 将key存入set中。如果已存在返回false，不存且添加返回true
func (s Set) Add(key interface{}) bool {
	if s.Has(key) {
		return false
	}
	s[key] = struct{}{}
	return true
}

// AddList 在set中加入list，如果加入成功返回true，失败返回false
func (s Set) AddList(keyList interface{}) bool {
	switch reflect.TypeOf(keyList).Kind() {
	case reflect.Slice:
		input := reflect.ValueOf(keyList)
		for i := 0; i < input.Len(); i++ {
			s[input.Index(i).Interface()] = struct{}{}
		}
		return true
	default:
		return false

	}
}

// Delete 在set里面删除这个key
func (s Set) Delete(key interface{}) {
	delete(s, key)
}

// StrList 返回set中string类型的key，如果不是string类型，会返回unknown
func (s Set) StrList() []string {
	retList := []string{}
	for key := range s {
		switch key.(type) {
		case string:
			retList = append(retList, key.(string))
		default:
			retList = append(retList, "unknown")
		}
	}
	return retList
}

// Len 获取set长度
func (s Set) Len() int {
	return len(s)
}
