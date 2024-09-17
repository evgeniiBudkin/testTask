package main

import (
	"fmt"
	"strconv"
	"strings"
)

var rim10 = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
var rim100 = []string{"", "X", "XX", "XXX", "LX", "L", "LX", "LXX", "LXXX", "XC", "C"}
var err = fmt.Errorf("не соответствует ТЗ, ошибка")

func main() {
	var inputUser string
	fmt.Println("Введите арифметическую операцию")
	fmt.Scanln(&inputUser)
	if len(inputUser) < 3 { //проверка длины введёной арифметической операции
		panic(err)
	} else {
		//поиск знака, и некоторые проверки
		flag, operation := finOperation(inputUser)
		if !flag {
			panic(err)
		} else {
			//делим числа по знаку
			numbers := strings.Split(inputUser, string(inputUser[operation]))

			// 1.1) проверяем арбские числа
			greenFlag, resalt := numbersOrNot(numbers)
			if greenFlag {
				//1.2) Вычисляем
				resultarab := myOperation(resalt, string(inputUser[operation]))
				fmt.Println("Вывод арабские: ", resultarab)
			} else {
				//----------------------------------------------------------//
				//2.1) проверяем римские цыфры
				resultrim := myOperation(rimCeck(numbers), string(inputUser[operation]))

				//2.3) Вычисляем
				if resultrim > 0 {
					fmt.Println(rimOut(resultrim))
				} else {
					panic(err)
				}
			}
		}
	}
	fmt.Println("Шалость удалась")
}

// поиск перебором O(n)
func finOperation(find string) (bool, int) {
	operation := []string{"+", "-", "*", "/"}
	resalt := []int{}
	for i, v := range find {
		for _, y := range operation {
			if y == string(v) {
				resalt = append(resalt, i)
			}
		}
	}
	//Проверяем наличие знаков, количество знаков, наличие знаков в начале и конце строки
	if len(resalt) == 0 || len(resalt) > 1 || len(find)-1 == resalt[len(resalt)-1] || resalt[0] == 0 {
		return false, 0
	} else {
		return true, resalt[0]
	}
}

// проверка арабских цифр
func numbersOrNot(numbers []string) (bool, []int) {
	resalt := []int{}
	for i, _ := range numbers {
		val, noDigital := strconv.Atoi(numbers[i])
		resalt = append(resalt, val)
		if noDigital != nil || val > 10 {
			return false, resalt
		}
	}
	return true, resalt
}

// вычисления
func myOperation(resalt []int, operation string) int {
	switch operation {
	case "+":
		return resalt[0] + resalt[1]
	case "-":
		return resalt[0] - resalt[1]
	case "*":
		return resalt[0] * resalt[1]
	case "/":
		if resalt[1] != 0 {
			return resalt[0] / resalt[1]
		}
	}
	panic(err)
}

// вывод римских цифр
func rimOut(number int) string {

	if number <= 9 {
		return rim10[number]
	} else if strconv.Itoa(number)[1] == 0 {
		return rim100[(number / 10)]
	} else {
		return rim100[(number/10)%10] + rim10[(number-((number/10)%10)*10)]
	}

}

// проверка на римские цифры.
func rimCeck(number []string) []int {
	resalt := []int{}
	for _, v := range number {
		for j, y := range rim10 {
			if y == strings.ToUpper(v) {
				resalt = append(resalt, j)
			}
		}
	}
	if len(resalt) < 2 {
		panic(err)
	}
	return resalt
}
