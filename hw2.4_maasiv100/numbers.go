package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Создаём массив
	var array [100]int

	//Заполняем массив рандомными числами от 0 до 100
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(100)
	}

	fmt.Println("Заполненный массив различными числами: \n", array)

	//сортируем по возростанию
	var b int
	var c bool
	for !c {
		c = true
		for i := 0; i < len(array); i++ {
			if i < len(array)-1 {
				if array[i] > array[i+1] {
					b = array[i]
					array[i] = array[i+1]
					array[i+1] = b
					c = false
				}
			}
		}
	}
	fmt.Println("Вывод подряд чисел из массива по возрастанию числа: \n", array)



	}



	// Все! дальше мозгов не хватает, перехожу к уроку массивы
