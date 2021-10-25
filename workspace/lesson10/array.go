package main

import "fmt"


func main() {
	// 指定数组大小
	var float_num_list1 [5]float32 = [5]float32{1.0, 2.0, 3.0, 4.0, 5.0}
	var float_num_list2 = [5]float32{1.0, 2.0, 3.0, 4.0, 5.0}
	int_num_list := [3]int{1, 2, 3}
    
    for index := range float_num_list1 {
        fmt.Println("[float_num_list1]index=", index)
    }
    
	for index, value := range float_num_list1 {
		fmt.Println("[float_num_list1]index=", index, "value=", value)
	}

	for index, value := range float_num_list2 {
		fmt.Println("[float_num_list2]index=", index, "value=", value)
	}

	for index, value := range int_num_list {
		fmt.Println("[int_num_list]index=", index, "value=", value)
	}

	// 编译器自行推导数组大小
	var balance1 []int = []int{1,2}
	var balance2 = []int{1,2,3}
	balance3 := []int{1, 2}
	fmt.Println("balance1=", balance1)
	fmt.Println("balance2=", balance2)
	fmt.Println("balance3=", balance3)

	balance := [5]int{1:10, 3:30} // 将数组下标为1和3的元素分别初始化为10和30
	fmt.Println(balance)

	num := balance[1]
	fmt.Println("num=", num)
	for i:=0; i<5; i++ {
		fmt.Printf("balance[%d]=%d\n", i, balance[i])
	}
}