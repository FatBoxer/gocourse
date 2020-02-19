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

	//Сортируем массив по возрастанию

	// Все! дальше мозгов не хватает, перехожу к уроку массивы


}
