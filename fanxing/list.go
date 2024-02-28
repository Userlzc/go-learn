package main

import (
	"fmt"
	"strconv"
)

/**
 * @Description
 * @Date 2024/2/21 18:01
 **/

// ListV1 泛型定义接口
type ListV1[T any] interface {
	Add(index int, val T)
	Append(val T) error
	Delete(index int) error
}

// 切片
type vector[T any] []T

// map
type M[K string, V any] map[K]V

// chan
type C[T any] chan T

// 泛型定义函数

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Printf("%v\n", v)
	}
}

// 泛型定义总的数据类型接口（约束类型）

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

func add[T NumStr](a, b T) T {
	return a + b
}

// 约束方法

type Price int

func (f Price) String() string {
	return strconv.Itoa(int(f))
}

type ShowPrice interface {
	String() string
}

// 注意的是 这里的T对应的是接口 因为实现了接口就是实现了该结构体的方法

func ShowPriceList[T ShowPrice](s []T) (res []string) {
	for _, v := range s {
		res = append(res, v.String())
	}
	return res

}

// 两种结合的方法

type Price2 int

func (f Price2) String() string {
	return strconv.Itoa(int(f))
}

type Price3 string

func (n Price3) String() string {
	return string(n)
}

type ShowPriceAth interface {
	String() string
	~int | ~string
}

func ShowPriceListV2[T ShowPriceAth](s []T) (res []string) {
	for _, v := range s {
		res = append(res, v.String())
	}
	return res

}

// 可比较
func findFunc[T comparable](a []T, v T) int {
	for i, e := range a {
		if e == v {
			return i
		}

	}
	return -1
}

// 写一个查找的方法

func Find[T any](val []T, filter func(t T) bool) T {
	for _, v := range val {
		if filter(v) {
			return v
		}
	}

	var t T
	return t

}
