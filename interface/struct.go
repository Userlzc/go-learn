package main

/**
 * @Description
 * @Date 2024/2/13 18:49
 **/

// 注意方法接收器和指针接收器的区别
// 无论声明结构体的时候是指针类型还是普通类型 都取决于接收器的类型 只有指针接收器才能改动

// 结构体自引用只能用指针
//type node struct {
//	next *node
//}
