package _func

import (
	"errors"
	"strconv"
)

/**
 * @Description
 * @Date 2024/2/13 0:48
 **/

func Func4() {
	myFunc3 := Func3
	s, err := myFunc3(1, 2)
	println(s, err)
	_, _ = myFunc3(2, 3)
}
func Func3(a, b int) (string, error) {
	c := a + b
	return strconv.Itoa(c), errors.New("err")
}

// 这里注意的是两种写法意义不同
// _, _ = myFunc3(2, 3) 常规的方法调用
//myFunc3 := Func3 这里是方法赋值给变量

// 局部方法 方法内声明一个局部方法作用域就在本方法内

func Func5() {
	fn := func(name string) string {
		return "hello" + name
	}
	fn("大明")

}

// 方法作为返回值

func Func6() func(name string) string {
	return func(name string) string {
		return "hello" + name
	}
}

func Func6Invoke() {
	sayHello := Func6()
	sayHello("daming")
}

// 匿名方法发起调用

func Func7() {
	f := func(name string) string {
		return "hello," + name
	}("大明")
	println(f)

}
