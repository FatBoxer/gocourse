package main
import "fmt"

/*Написать программу для конвертации рублей в доллары.
Программа запрашивает сумму  в рублях и выводит сумму в долларах.
Курс доллара задайте константой. */

func main() {
	var rubls float64
	var dolrs float64
	const curs float64 = 62.45


	fmt.Println("-------======$$$======-------")
	fmt.Println("Сегодняшний Курс Рубля к Доллару:", curs)
	fmt.Println("Введите сумму Рублей для продажи:")
	fmt.Scanln(&rubls)

	dolrs = rubls / curs

	fmt.Println("=====================")
	fmt.Println("Привет,сегодня в долларах вы получите:")
	fmt.Printf("%.1f", dolrs)
}
