package main

import (
	"fmt"
	"untitled/until"
)

func test() {
	//定义数组
	var scores = [5]int{3, 6, 7, 8, 2}
	fmt.Printf("数组的类型为：%T\n", scores)
	for i := 0; i < len(scores); i++ {
		fmt.Printf("请录入第一个%d学生的成绩", i+1)
		fmt.Scanln(scores[i])
	}
	//	数组便利
	for i := 0; i < len(scores); i++ {
		fmt.Printf("第一个%d学生的成绩为：%d\n", i+1, scores[i])
		fmt.Scanln(scores[i])
	}
	//	for range循环
	for key, value := range scores {
		fmt.Printf("第一个%d学生的成绩为:%d\n", key+1, value)
	}
	//	for range循环(忽略某个值用下划线就可以了)
	for _, value := range scores {
		fmt.Printf("第一个学生的成绩为:%d\n", value)
	}

}
func test1(arr *[3]int) {
	(*arr)[0] = 7
	//	引用传递靠指针完成
}
func test2() {
	//	定义二维数组
	var arr = [2][3]int{{1, 4, 7}, {2, 3, 4}}
	var arr2 = [...]int{1, 4, 5, 7}
	fmt.Println(arr)
	fmt.Println(arr2)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Println("遍历后的值", arr[i][j])
		}
	}
	//for range循环
	for key, value := range arr {
		for k, val := range value {
			fmt.Printf("arr[%v][%v]=%v\t", key, k, val)
		}
	}
}
func main() {
	var arr = [3]int{1, 6, 7}
	test1(&arr) //传入arr数组的地址
	fmt.Println(arr)
	until.GetConn()
	test()
	fmt.Println("成功")
	test2()
}
