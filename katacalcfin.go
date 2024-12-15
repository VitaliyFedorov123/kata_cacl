package main

import (
	"fmt"
	"strings"
)

// Таблица для преобразования римских чисел (расширена до 100)
var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	"XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20,
	"XXX": 30, "XL": 40, "L": 50, "LX": 60, "LXX": 70, "LXXX": 80, "XC": 90, "C": 100,
}
var arabicToRoman = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
	11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV", 16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX", 20: "XX",
	30: "XXX", 40: "XL", 50: "L", 60: "LX", 70: "LXX", 80: "LXXX", 90: "XC", 100: "C",
}

func romanToNumber(roman string) int {
	return romanToArabic[roman]
}

// Преобразование арабского числа в римское
func numberToRoman(number int) string {
	return arabicToRoman[number]
}

// Выполнение математической операции
func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}

func main() {
	// Ввод от пользователя
	fmt.Println("Введите выражение (например, 'V + V' или '10 * 10'):")
	var input string
	fmt.Scanln(&input)

	// Убираем пробелы
	input = strings.ReplaceAll(input, " ", "")

	// Определяем оператор
	var operator string
	for _, op := range "+-*/" {
		if strings.Contains(input, string(op)) {
			operator = string(op)
			break
		}
	}

	// Разделяем числа по оператору
	parts := strings.Split(input, operator)
	if len(parts) != 2 {
		fmt.Println("Ошибка: введите два числа и оператор.")
		return
	}

	num1Str, num2Str := parts[0], parts[1]

	// Определяем, римские числа или арабские
	isRoman := strings.ContainsAny(num1Str, "IVXLCDM") && strings.ContainsAny(num2Str, "IVXLCDM")

	var num1, num2 int
	if isRoman {
		num1 = romanToNumber(num1Str)
		num2 = romanToNumber(num2Str)
	} else {
		fmt.Sscanf(num1Str, "%d", &num1)
		fmt.Sscanf(num2Str, "%d", &num2)
	}

	// Проверяем диапазон ввода (от 1 до 10)
	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		fmt.Println("Ошибка: числа должны быть от 1 до 10.")
		return
	}

	// Выполняем операцию
	result := calculate(num1, num2, operator)

	// Проверяем, чтобы результат не превышал 100
	if result > 100 {
		fmt.Println("Ошибка: результат превышает 100.")
		return
	}

	// Выводим результат
	if isRoman {
		fmt.Println("Результат:", numberToRoman(result))
	} else {
		fmt.Println("Результат:", result)
	}
}
