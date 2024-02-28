package main

import "fmt"

/**
 * @Description
 * @Date 2024/2/13 18:50
 **/

type node struct {
}
type LinkedList struct {
	next *node
}

func UserList() {
	l1 := LinkedList{}
	l1Ptr := &l1               // 取地址
	var l2 LinkedList = *l1Ptr // 解引用
	fmt.Printf("%+v \n", l2)
}
