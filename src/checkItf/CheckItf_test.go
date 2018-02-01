package checkItf

import (
	"testing"
	"fmt"
)

func TestItf(t *testing.T)  {
	var p Animal= People{Name:"yanghai",Age:18,Sex:"man"}
	if p1, ok := p.(Animal2); ok {
		p1.Eat("apple")
	}else{
		fmt.Println("不是动物类型")
	}


	var a [12]int
	a[0] = 12


	b := a[:cap(a)/2]
	b[2] = 23

	fmt.Println("b",b)
}
