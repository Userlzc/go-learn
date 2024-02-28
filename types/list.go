package main

/**
 * @Description
 * @Date 2024/2/14 17:11
 **/

// 接口的声明

type List interface {
	Add(index int, val any)
	Append(val any) error
	Delete(index int) error
}
type LinkedList struct {
	head node
}

func (l LinkedList) Add(index int, val any) {
	//TODO implement me
	panic("implement me")
}

func (l LinkedList) Append(val any) error {
	//TODO implement me
	panic("implement me")
}

func (l LinkedList) Delete(index int) error {
	//TODO implement me
	panic("implement me")
}

type node struct {
}
