package main

import "fmt"

func main() {
	var num1, num2 int //две переменные для математических операций
	var mop int        //переменная выбора математической операции
	fmt.Println("Учебная программа Калькулятор")

	fmt.Println("Введите первое число")
	fmt.Scan(&num1)
	fmt.Println("Вы ввели ", num1)

	fmt.Println("Введите второе число")
	fmt.Scan(&num2)
	fmt.Println("Вы ввели ", num2)

	fmt.Println(`Выберите нужную операцию:
	1. Сложение
	2. Вычитание
	3. Умножение
	4. Деление
	5. Остаток`)
	fmt.Scan(&mop)
	fmt.Println("Вы ввели ", mop)

	switch mop {
	case 1:
		fmt.Println("Первое число + Второе число = ", num1+num2)
	case 2:
		fmt.Println("Первое число - Второе число = ", num1-num2)
	case 3:
		fmt.Println("Первое число * Второе число = ", num1*num2)
	case 4:
		fmt.Println("Первое число / Второе число = ", num1/num2)
	case 5:
		fmt.Println("Первое число % Второе число = ", num1%num2)
	}

	/*
		//	fmt.Println(a * 5)
		for i := 0; i < a; i++ {
			fmt.Println("ntcn", i)
		}
		//var x [5]int

		x := make([]float64, 5)

		x[4] = 100
		x = append(x, 0.1)
		fmt.Println(x)
		fmt.Println(x[4] + 0.01)
		//m, n, o := x[2:4]
		//fmt.Println(m, n, o)
		fmt.Println()
	*/

}
