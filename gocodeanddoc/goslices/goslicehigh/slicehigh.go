package main

import "fmt"

//可以进行多次切片操作,slice可以多次进行
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	//s1 有4个值
	s1 := arr[2:6]
	//len是当前切片可以取到的值，cap是切片可以用其他切片来将当前切片从第一个下标开始扩展到原数组最后一个下标
	fmt.Println(len(s1), cap(s1))
	//s1[4]取不出来会报错
	//s2 结果 5,6
	//s1底层知道5,6下标的存在，但取不出来，再用一个切片s2可以取出来
	//切片的结构分三个，ptr指向第一个切片的下标和len代表可以取出的切片，cap容量代表原数组从ptr到末尾所有值，因为再用一个切片可以取出这些s1取不出的值
	s2 := s1[3:5]
	fmt.Println(s2)
}
