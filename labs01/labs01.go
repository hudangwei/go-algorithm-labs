package main

import (
	"fmt"
	"unsafe"
)

func ReverseSlice(s []byte, m, n int) {
	for m < n {
		temp := s[m]
		s[m] = s[n]
		m++
		s[n] = temp
		n--
	}
}

func LeftRotateSlice(s []byte, m int) {
	n := len(s)
	ReverseSlice(s, 0, m-1)
	ReverseSlice(s, m, n-1)
	ReverseSlice(s, 0, n-1)
}

func main() {
	str1 := "hello,world!"
	s1 := []byte(str1) //string必须转成[]byte才可以对string修改，但会重新分配内存，复制数据
	LeftRotateSlice(s1, 3)
	fmt.Println(str1)

	fmt.Println(*(*string)(unsafe.Pointer(&s1)))
	fmt.Println(string(s1))
}
