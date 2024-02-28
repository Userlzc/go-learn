package main

// 闭包 ：方法+绑定的上下文  注意：对象被闭包引用的话不会被垃圾回收导致会引起内存泄漏
func main() {
	ClosureInvoke()

}
func Closure(name string) func() string {

	return func() string {
		return "hello," + name
	}

}
func ClosureInvoke() {
	c := Closure("大明")
	println(c())
}
