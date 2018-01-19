package main

import (
	"fmt"
)

type X struct {
}

//func (x *X) test()  {
//	fmt.Println("hi!,",x)
//}

func (x X) test()  {
		fmt.Println("hi!,",x)
}

func main() {

	X{}.test()
}