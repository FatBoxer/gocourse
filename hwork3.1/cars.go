package main
/*Описать несколько структур — любой легковой автомобиль и грузовик.
Структуры должны содержать марку авто, год выпуска, объем багажника/кузова,
запущен ли двигатель, открыты ли окна, насколько заполнен объем багажника.*/

import "fmt"

type carSedan struct {
	brand string
	series string
	age int
	engineStatus string
	windowStatus string
	trunkVolume float64
	trunkStatus string
}


type carTruck struct {
	brand, series string
	age int
	engineStatus, windowStatus string
	trunkVolume float64
	trunkStatus string
}


func main() {
	var car1 = carSedan {"Audi", "A5", 2011,  "ON","OFF", 450, "FULL"}
	var car2 = carSedan {"BMW", "F1", 2010,  "OFF","OFF", 550.5, "FULL"}
	// дргой вариант инициализации типа автомобиля
	car3 := carTruck {"IVECO", "EuroCargo", 2016,  "OFF","OFF", 1884, "EMPTY"}
	car4 := carTruck {"IVECO", "Stralis", 2006,  "ON","OFF", 1650, "FULL"}


	fmt.Println("Автомобиль:", car1)
	fmt.Println("Автомобиль:", car2)
	fmt.Println("Автомобиль:", car3)
	fmt.Println("Автомобиль:", car4)

	// делаем бвм и ауди моложе

	car1.age = 2018 // меняем год выпуска на 2018
	car2.age = 2020 // меняем год выпуска на 2020

	fmt.Println("------------------chit------- \n")
	fmt.Println("Автомобиль:", car1)
	fmt.Println("Автомобиль:", car2)

}