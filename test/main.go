package main

import (
	"errors"
	"fmt"
)

/*
      实现切片的删除操作
实现删除切片特定下标元素的方法
*要求一 : 能够实现删除操作就可以
*要求二 : 考虑使用比较高性能的实现
*要求三 : 改造为泛型
*要求四 : 支持缩容，并且设计缩容机制

*/

type NumStr interface {
	Num | Str
}

type Str interface {
	string
}

// 加 ~ 的意思是衍生类型也可以  比如int的衍生类型 type Integer int

type Num interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~complex64 | ~complex128
}

func DeleteEle[T NumStr](idx int, arr []T) ([]T, error) {
	length := len(arr)
	if idx > length || idx < 0 {
		return nil, errors.New("索引超出范围")
	}
	for i := idx; i < length-1; i++ { // 也可以写成i+1<length
		arr[i] = arr[i+1]
	}
	// 去掉最后一个多余得元素
	arr = arr[:length-1]
	// 对其进行判断是否需要扩容

	return arr, nil

}

// Shrink对其判断是否需要缩容 不需要缩容的则返回原切片 需要的则创建新的切片然后append

func Shrink[T NumStr](arr []T) []T {
	c, l := cap(arr), len(arr)
	n, changed := calCapacity(c, l)
	if !changed {
		return arr
	}
	s := make([]T, 0, n)
	s = append(s, arr...)
	return s

}

// 具体的缩容判断方法
func calCapacity(c, l int) (int, bool) {
	// 这里可以是一个标准来判断
	if c <= 64 {
		return c, false
	}
	if c <= 2048 && (c/l >= 4) {
		return c / 2, true
	}
	if c > 2048 && (c/l >= 2) {
		factor := 0.625
		return int(float32(c) * float32(factor)), true
	}
	return c, false
}

func main() {
	oldVals := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("oldVals: %v, len: %d ,cap: %d \n", oldVals, len(oldVals), cap(oldVals))
	ele, err := DeleteEle(5, oldVals)
	if err != nil {
		return
	}
	fmt.Printf("ele: %v, len: %d ,cap: %d \n", ele, len(ele), cap(ele))
}
