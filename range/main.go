package main

import "fmt"

/**
因为for range会创建每个元素的副本，而不是返回每个元素的引用，所以它的地址是不会发生变化的。
*/

type girl struct {
	Name string
	Age  int
}

func main() {
	gl := make(map[string]*girl)
	studs := []girl{
		{Name: "Lili", Age: 23},
		{Name: "Lucy", Age: 24},
		{Name: "Han Mei", Age: 21},
	}

	for _, v := range studs {
		gl[v.Name] = &v
		fmt.Printf("%s addr: %v\n", v.Name, &v.Age)
	}
	// Lili addr: 0xc0000a4028
	//Lucy addr: 0xc0000a4028
	//Han Mei addr: 0xc0000a4028

	for mk, mv := range gl {
		fmt.Println(mk, "=>", mv.Age)
	}
}

// 运行结果是
//Lili => 21
//Lucy => 21
//Han Mei => 21

// 解决方案
/**
1.在每次循环时，创建一个临时变量tem，将v的值赋给tem，相当于copy一个v到tem，再将tem地址传给指针即可
for _, v := range studs {
	temp := v
	gl[v.Name] = &temp
}

2.不使用指针作为map的value，同理下面赋值也不使用地址赋值
gl := make(map[string]girl)

for _, v := range studs {
		gl[v.Name] = v
	}
*/
