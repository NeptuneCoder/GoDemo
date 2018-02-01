package main


import (
	"fmt"
	"time"
)

func counting(c chan<- int) {

	for i := 0; i < 10; i++ {
		msg:= "starting a gofunc id  = "
		fmt.Println("msg:", msg,string(i))
		time.Sleep(1 * time.Second)
		c <- i
	}
	close(c)//当写完以后，通知不能再写入了。开始读取
}
func main() {
	bus := make(chan int)
	go counting(bus)
	for count := range bus {
		fmt.Println("count:", count)
	}
}
