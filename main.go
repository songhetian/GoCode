package main

import "fmt"

func main() {
	//数组长度是固定的
	/*定义方式1*/
	var array1 [4]int

	array1[0] = 1
	array1[1] = 2
	array1[2] = 3

	/*定义方式2*/

	array2 := [3]int {1,2,3}

	for k,v := range array2 {
		fmt.Println(k,"-",v)
	}

	fmt.Printf("%#v",array2)
}
