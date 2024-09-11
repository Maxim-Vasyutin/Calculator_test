package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToArabicMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
}

var arabicToRomanMap = map[int]string{
	1:   "I",
	4:   "IV",
	5:   "V",
	9:   "IX",
	10:  "X",
	40:  "XL",
	50:  "L",
	90:  "XC",
	100: "C",
}

func romanToArabic(romanNum string) (int, error) {
	result := 0
	for i := 0; i < len(romanNum); i++ {

		current := romanToArabicMap[string(romanNum[i])]

		if i+1 < len(romanNum) {
			next := romanToArabicMap[string(romanNum[i+1])]
			if current < next {
				result -= current
			} else {
				result += current
			}
		} else {
			result += current
		}
	}
	if result > 10 || result < 1 {
		panic("Используйте числа, эквивалент которых находится в диапазоне от I(1) до X(10) включительно. Повторите ввод ещё раз.")
	}
	return result, nil
}

func check(argument string) bool {
	romanChars := "IVXLC"

	for i := 0; i < len(argument); i++ {
		char := argument[i]
		res := strings.Index(romanChars, string(char))
		if res == -1 {
			return false
		}
	}
	return true
}

func main() {
	var input string
	var err error
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите выражение с римскими или арабскими числами: ")
	input, err = reader.ReadString('\n')
	if err != nil {
		panic("Ошибка при чтение ввода")
	}

	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")

	var result int
	//Проверка на наличие "всего что нужно" для начала работы с выражением
	if len(parts) != 3 {
		panic("Неправильный формат ввода. Ожидается формат: ЧИСЛО ОПЕРАТОР ЧИСЛО")
	}

	if check(parts[0]) && check(parts[2]) {
		//Блок кода с логикой вычисления для римских цифр (можно подключить вызов отдельной функции, чтобы не засорять main)

		numRoman1, err := romanToArabic(parts[0])
		numRoman2, err := romanToArabic(parts[2])

		if err != nil {
			fmt.Print(err)
			return
		}

		switch parts[1] {
		case "+":
			result = numRoman1 + numRoman2
		case "-":
			result = numRoman1 - numRoman2
			if result <= 0 {
				panic("Ошибка! Результат меньше I(1). Повторите ввод ещё раз")

			}
		case "*":
			result = numRoman1 * numRoman2
		case "/":
			if numRoman2 == 0 {
				panic("Ошибка! На ноль делить нельзя! Ну и '0' не существует в римском формате счисления)) Повторите ввод ещё раз")

			} else {
				result = numRoman1 / numRoman2
			}
		default:
			panic("В нашем калькуляторе не предусмотренн такой знак! Повторите ввод ещё раз.")

		}

		var romanNum string
		for result > 0 {
			for _, key := range []int{100, 90, 50, 40, 10, 9, 5, 4, 1} {
				if result >= key {
					result -= key
					romanNum += arabicToRomanMap[key]
					break
				}
			}
		}
		fmt.Println("Результат:", romanNum)
	} else if check(parts[0]) == false && check(parts[2]) == false {
		//Блок кода с логикой вычисления для арабских цифр
		numArabic1, _ := strconv.Atoi(parts[0])
		numArabic2, _ := strconv.Atoi(parts[2])

		if numArabic1 > 10 || numArabic2 > 10 || numArabic1 <= 0 || numArabic2 <= 0 {
			panic("Используйте числа, эквивалент которых находится в диапазоне от 1 до 10(включительно). Повторите ввод ещё раз.")
		}

		switch parts[1] {
		case "+":
			result = numArabic1 + numArabic2
		case "-":
			result = numArabic1 - numArabic2
		case "*":
			result = numArabic1 * numArabic2
		case "/":
			if numArabic2 == 0 {
				panic("Ошибка! На ноль делить нельзя! Повторите ввод ещё раз")

			} else {
				result = numArabic1 / numArabic2
			}
		default:
			panic("В нашем калькуляторе не предусмотренн такой знак! Введите выражение ещё раз.")

		}
		fmt.Println("Результат:", result)
	} else {
		panic("Ошибка! Вы ввели 2 разных системы счисления")
	}
}
