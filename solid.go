package main

import "fmt"

//S : Single responsibility principle (принцип единой ответственности) / Делай что-то одно и делай это хорошо
//O : Open/Closed principle (Принцип открытия/закрытия) / Модуль должен быть открыт для расширений, но закрыт для модификации
//L : Liskov substitution principle (Принцип подстановки/замены Барбары Лисков) Производные методы не должны требовать большего и предоставлять не меньше
//I : Interface segregation principle (принцип разделения интерфейса)
//D :  Dependency inversion principle (Принцип инверсии зависимости)
//каждая зависимость между модулями должна быть нацелена на абстрактный класс или интерфейс.
//Ни одна зависимость не должна быть нацелена на конкретный класс.

func main() {

	arr := [4]string{
		"a1", "a2", "a3", "a4",
	}

	fmt.Printf("%#v\n", arr)

	s1 := arr[:2]
	fmt.Printf("%#v\n", s1)

	s2 := arr[1:]
	fmt.Printf("%#v\n", s2)

	s2[0] = "bbb"

	fmt.Printf("%#v\n", arr)
	fmt.Printf("%#v\n", s1)

}

func plus(a *int, c int) {
	*a = *a + c
}
