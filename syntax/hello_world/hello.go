package main

/**
 * @Description
 * @Date 2024/2/11 18:01
 **/

func Hello() {
	println("hello hi")
}

// const可以支持类型推断
const a int = 123
const Ex = "12你好"

// const如果是重新从某一个词判断的话最好另起一个const
const (
	Status = iota
	Status1
	Status2
	Status3
	Status4
	Sa6 = 6
	Sam // 这里之所以是6 是因为上边拍下来的而不是从6变7 故要另起一个const

)
const ()
