package checkItf

import "fmt"


type Animal interface {
	 Run() error
	 Eat(foodName string)
}

type Animal2 interface {
	Run() error
	Eat(foodName string)
}


type People struct {
	Name string
	Age int
	Sex string
}

func (p People) Run() error {
	return nil
}
func (p People) Eat(foodNmae string)  {
	fmt.Println("name = ",p.Name ," eat = ",foodNmae)
}