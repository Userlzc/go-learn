package main

/**
 * @Description
 * @Date 2024/2/14 17:14
 **/

// 遇事不决用指针
func main() {
	//// u1指向的是一个user对象的指针
	//u1 := &User{}
	//println(u1)
	//// 与上边性质一样（new 理解为自动分配内存，并把内存置为0）
	//u1 = new(User)
	//println(u1)
	//
	//// u2和u3的字段也都是零值
	//u2 := User{}
	//fmt.Printf("%+v\n", u2)
	//u2.Name = "Jerry"
	//println(u2.Name)
	//
	//var u3 User
	//println(u3)
	//
	//var u4 *User
	//println(u4) //输出0*0 代表nil

	ChangeUser()

}
