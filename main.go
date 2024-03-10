package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var romNumbers = map[string]int{
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

func main() {
	fmt.Println("Hi. Get it\n" +
		"Введите пример:")
	intType, first, second, sign, err := readInput()
	if err != nil {
		fmt.Println("Ошибка при вводе данных:\n", err)
		return
	}
	if intType == "arab" {
		firstNum, err1 := strconv.Atoi(first)
		if err1 != nil {
			fmt.Println("Exeption\n", err1)
			return
		}
		secondNum, err2 := strconv.Atoi(second)
		if err2 != nil {
			fmt.Println("Exeption\n", err2)
			return
		}
		res, err3 := calc(firstNum, secondNum, sign)
		if err3 != nil {
			fmt.Println("Exeption\n", err3)
			return
		} else {
			fmt.Println("Ответ: ", res)
		}
	} else {
		firstNum := romToInt(first)
		secondNum := romToInt(second)
		res, err1 := calc(firstNum, secondNum, sign)
		if err1 != nil {
			fmt.Println("Exeption\n", err1)
			return
		} else {
			final, err2 := arabToRom(res)
			if err2 != nil {
				fmt.Println("Exeption\n", err2)
				return
			}
			fmt.Println("Ответ: ", final)
		}
	}
}

func readInput() (string, string, string, string, error) {
	stdin := bufio.NewReader(os.Stdin)
	usInput, _ := stdin.ReadString('\n')
	usInput = strings.TrimSpace(usInput)
	intType, first, second, sign, err := checkInput(usInput)
	if err != nil {
		return "", "", "", "", err
	}
	return intType, first, second, sign, err
}

func checkInput(input string) (string, string, string, string, error) {
	r := regexp.MustCompile("\\s+")
	replace := r.ReplaceAllString(input, "")
	arr := strings.Split(replace, "")
	var intType, first, second, sign string
	for index, value := range arr {
		isN := isNumber(value)
		isS := isSign(value)
		isR := isRomanNumber(value)
		if !isN && !isS && !isR {
			return "", "", "", "", errorHandler(1)
		}
		if isS {
			if sign != "" {
				return "", "", "", "", errorHandler(6)
			} else {
				sign = arr[index]
			}
		}
		if (isN && intType != "rom") || (isR && intType != "arab") {
			if intType == "" {
				if isN {
					intType = "arab"
				} else {
					intType = "rom"
				}
			}
			if first == "" && !(index+1 == len(arr)) && isSign(arr[index+1]) {
				slice := arr[0:(index + 1)]
				first = strings.Join(slice, "")
			} else if index+1 == len(arr) && first != "" {
				slice := arr[(len(first) + 1):]
				second = strings.Join(slice, "")
			}
		} else if (intType == "arab" && isR) || (intType == "rom" && isN) {
			return "", "", "", "", errorHandler(2)
		}
	}
	if second == "" || first == "" || sign == "" {
		return "", "", "", "", errorHandler(3)
	}
	return intType, first, second, sign, nil
}

func isNumber(c string) bool {
	if c >= "0" && c <= "9" {
		return true
	} else {
		return false
	}
}

func isSign(c string) bool {
	if c == "+" || c == "-" || c == "/" || c == "*" {
		return true
	} else {
		return false
	}
}
func isRomanNumber(n string) bool {
	_, ok := romNumbers[n]
	if ok {
		return true
	} else {
		return false
	}
}

func romToInt(n string) int {
	var out int
	array := strings.Split(n, "")
	for i, value := range array {
		if i+1 != len(array) && romNumbers[value] < romNumbers[array[i+1]] {
			out -= romNumbers[value]
		} else {
			out += romNumbers[value]
		}
	}
	return out
}

func arabToRom(n int) (string, error) {
	var out string
	if n <= 0 {
		return "", errorHandler(7)
	}
	arrArab := [9]int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	arrRom := [9]string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for n > 0 {
		for i := 0; i < 13; i++ {
			if arrArab[i] <= n {
				out += arrRom[i]
				n -= arrArab[i]
				break
			}
		}
	}
	return out, nil
}

func calc(first int, second int, sign string) (int, error) {
	if first > 10 || second > 10 {
		return 8, errorHandler(8)
	}
	switch {
	case sign == "+":
		return first + second, nil
	case sign == "-":
		return first - second, nil
	case sign == "*":
		return first * second, nil
	case sign == "/" && second != 0:
		return first / second, nil
	case sign == "/" && second == 0:
		return 4, errorHandler(4)
	default:
		return 5, errorHandler(5)
	}
}

func errorHandler(e int) error {
	return errors.New(errorsDist[e])
}

var errorsDist = map[int]string{
	1: "Нераспознанные символы",
	2: "Некорректный ввод. Только арабские или только римские цифры",
	3: "Некорректный ввод. Введите 2 числа и знак",
	4: "Делить на 0 запрещено",
	5: "Calc Exeption",
	6: "Некорректный ввод",
	7: "В римском счислении не существует отрицательных чисел и нуля",
	8: "Ввод от 0 до 10",
}
