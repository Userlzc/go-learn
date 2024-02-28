package main

import "fmt"

/**
 * @Description
 * @Date 2024/2/14 17:14
 **/

type User struct {
	Name     string
	Age      int
	NickName string
}

func (u *User) ChangeAge(age int) {
	u.Age = age

}

// 值传递

func (u User) ChangeName(name string) {
	fmt.Printf("Change name的地址 %p \n", &u)
	u.NickName = name

}

func ChangeUser() {
	u1 := User{
		Name: "Tom",
		Age:  18,
	}

	fmt.Printf("%+v \n", u1)
	fmt.Printf("u1 address %p \n", &u1)

	u1.ChangeName("Jerry")
	u1.ChangeAge(35)
	fmt.Printf("%+v \n", u1)

	//这里u2是取址 但是掉用changeName的方法仍然没有发生变化
	// 故 和初始化无关 取决于的是指针接收器和方法接收器，只有指针接收器才可以
	// 对其进行修改

	println("-------u2--------------")
	u2 := &User{
		Name: "Tom",
		Age:  18,
	}
	fmt.Printf("%+v \n", u2)
	fmt.Printf("u1 address %p \n", &u2)

	u1.ChangeName("Jerry")
	u1.ChangeAge(35)
	fmt.Printf("%+v \n", u2)
}
