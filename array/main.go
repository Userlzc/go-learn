package main

import "fmt"

func main() {

	// 数组  数组的长度和容量是相等的
	a1 := [3]int{9, 8, 0}

	a2 := [3]int{9, 8}

	var a3 [3]int
	fmt.Printf("a1长度为%d，容量为%d \n", len(a1), cap(a1))
	fmt.Printf("a2长度为%d，容量为%d \n", len(a2), cap(a2)) // 这里虽然只有两个元素 但是长度是3 因为最后一个元素被填充了0
	fmt.Printf("a3长度为%d，容量为%d ", len(a3), cap(a3))   // 同理  也是被填充了0

	// 切片的声明最好采用这种

	a4 := make([]int, 3, 4) // 无序的
	fmt.Printf("%v的长度为%d，容量为%d \n", a4, len(a1), cap(a1))
	// append扩容机制  尽量避免
	// 子切片 左闭右开[ : ](包左不包右)  ,如果发生扩容的话旧不在共享扩容 没有发生扩容的话就共享数组

	// map声明同样适用make（map[string]string,4）
	// map判断存不存在时的方法  val,ok:=m[key],如果不存在的话,那么val返回的是对应类型的默认值 ok则为false
	// map的delete方法没有返回值 仍然需要上面的方法进行判断

	// 数组可比较  切片不可比较

	// map[any]any 作为参数传递，但是你传递一个map[int]string
}
