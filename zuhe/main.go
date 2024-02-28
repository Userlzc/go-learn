package main

// 组合 （组合里结构体里的方法时可以随便使用）
type Outer struct {
	Inner
}
type Inner struct {
}

func (i Inner) Hello() {
	println("我是inner")
}

// 如何组合里边放的是指针的话就需要注意初始化

type Outer1 struct {
	*Inner
}

func main() {
	var o Outer
	o.Hello()

	// 因为是指针所以要初始化掉
	m := Outer1{
		Inner: &Inner{},
	}
	m.Hello()

	var b = B{}
	b.Change()
}

// 也可以在结构体里组合接口
// 这里就是说如果组合的结构体实现了所有接口那么意味着B也实现了此接口

// 组合不是继承 没有多态 （如果B重写了这个change方法那么输出的结果也是a的change）

type Name interface {
	Change(name string) bool
}
type B struct {
	A
}
type A struct {
}

func (a A) Change() bool {
	return true
}
