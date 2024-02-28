package main

import "fmt"

func main() {
	// 切片
	v := vector[int]{53, 64}
	PrintSlice(v)
	//map
	m := M[string, int]{"key": 1}
	m["key"] = 2
	fmt.Println(m)
	// chan
	c := make(C[int], 10)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	//
	PrintSlice[int]([]int{22, 5, 78, 52, 31, 1})
	PrintSlice[string]([]string{"你", "好", "啊", "我", "是", "GO"})
	PrintSlice[byte]([]byte{'a', 'd', 's', 'g'})
	PrintSlice[float64]([]float64{22.6, 5.1, 7.8, 5.2, 3.1, 1.2})

	// 演示泛型接口
	fmt.Println(add(3, 4))
	fmt.Println(add("hello", "world"))

	// 约束方法和类型结合
	fmt.Println(ShowPriceListV2([]Price2{1, 2}))
	fmt.Println(ShowPriceListV2([]Price3{"a", "b"}))

	// 可比较
	fmt.Println(findFunc([]int{1, 2, 3, 4, 5, 6}, 5))
}
