package main

/**
defer可以分为两大类
1.有参（defer是在返回之前执行，在此之前修改返回值）-->返回值
 1.1返回值只定义类型 没有名字 那么是改不掉的
 1.2 返回值定义了类型和名字那么就能改掉


2.无参
  2.1 作为参数传入时在定义defer的时候就确定了
  2.2 作为闭包引入时执行defer对应的方法的时候才确定

3.特殊情况



*/

func main() {
	println(DeferReturnV2().name)
}

// 1.1 返回值只定义类型

func DeferReturn() int {
	i := 0
	defer func() {
		i = 1
	}()
	return i
}

// 1.2 返回值类型和名字都有

func DeferReturnV1() (i int) {
	i = 0
	defer func() {
		i = 1
	}()
	return i
}

// 2.1 作为参数传递

func DeferCloser() {
	i := 0
	defer func(val int) {
		println(val)
	}(i)
	i = 1
	// 最终输出0
}

//2.2 作为闭包传入

func DeferCloserV1() {
	i := 0
	defer func() {
		println(i)
	}()
	i = 1
	// 最终输出1
}

// 3.特殊

func DeferReturnV2() *MyStruct {
	a := &MyStruct{
		name: "Jerry",
	}
	defer func() {
		a.name = "Tom"
	}()
	return a
	// 输出TOm
}

type MyStruct struct {
	name string
}
