package main

/**
 * @Description
 * @Date 2024/2/13 18:55
 **/
// 衍生类型
// 想使用第三方库 但有没有办法修改源码，又想扩展第三方库的结构体方法时

type Fish struct {
}

func (f Fish) Swim() {

}

type FakeFish Fish // 衍生类型

func (fa FakeFish) Swim() {

}

func UseFish() {
	var u Fish
	u.Swim()

	f2 := FakeFish{}
	f2.Swim()

	// 类型转换

	f3 := FakeFish(u)
	f1 := Fish(f2)
	println(f3, f1)

	// 别名

	type Yu = Fish
	y := Yu{}
	y.Swim()
}
