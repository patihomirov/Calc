package main

import "fmt"

func main() {
	var num1, num2 int //две переменные для математических операций
	var select1 int    //переменная выбора математической операции
	fmt.Println("Учебная программа Калькулятор")
	defer fmt.Println("До свидания!") //выпендримся и сделаем отложенную команду после завершения программы по Break
	//	for k:=0, k<10  k++ { //зацикливаем программу, так обычно нельзя без использования защиты от бесконечного цикла, но для учебных целей сделаем так. Выход из программы по Break в конце
	//запрашиваем данные для расчета
	fmt.Println("Введите первое целое число")
	fmt.Scan(&num1)
	fmt.Println("Вы ввели ", num1)

	fmt.Println("Введите второе целое число")
	fmt.Scan(&num2)
	fmt.Println("Вы ввели ", num2)

	fmt.Println(`Выберите нужную операцию:
	1. Сложение
	2. Вычитание
	3. Умножение
	4. Деление
	5. Остаток`)
	fmt.Scan(&select1)
	fmt.Println("Вы ввели ", select1)

	//Выводим результаты расчета
	switch select1 {
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

	//Предлагаем выход из программы
	var select2 int //Объявляем переменную выбора выйти из программы или остаться, но могли использовать повторно select1
	fmt.Println(`Желаете ли повторить расчет с новыми данными?
	1. Повторить расчет
	2. Выйти из программы`)
	fmt.Scan(&select2)
	switch select2 {
	case 1:
		fmt.Println(`Вы выбрали "1. Повторить расчет"`)
	case 2:
		fmt.Println(`Вы выбрали "2. Выйти из программы"`)
		//fmt.Println(`Вы выполнили расчет %d раза`, i)
		break
	}
	//	}
}
