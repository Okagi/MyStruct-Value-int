package main

import "fmt"

type MyStruct struct {
	id    int
	name  string
	Value int
}

func (ms MyStruct) ChangeValue(newValue int) {
	ms.Value = newValue
}

func main() {

	myStruct := MyStruct{Value: 3}
	fmt.Println("Ваша переменная:", myStruct.Value)

	myStruct.ChangeValue(20)
	fmt.Println("Новая структура:", myStruct.Value)
}
