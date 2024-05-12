package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	input := readLine()
	v1, operator, v2 := parseInput(input)
	result := calculate(v1, operator, v2)

	printResult(result)
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return ""
	}
	return strings.TrimSpace(str)
}

var r = regexp.MustCompile(`(".*?"|\d+)\s*([+\-*\/])\s*(".*?"|\d+)`)

func parseInput(input string) (string, string, string) {
	// if len([]rune(input)) > 10 {
	// 	panic("Слишком длинное выражение")
	// }

	parts := r.FindAllStringSubmatch(input, -1)
	if len(parts) != 1 || len(parts[0]) != 4 {
		panic("Неверный формат выражения")
	}

	return strings.Trim(parts[0][1], "\""), parts[0][2], strings.Trim(parts[0][3], "\"")
}

func calculate(v1, operator, v2 string) string {
	switch operator {
	case "+":
		return v1 + v2
	case "-":
		return strings.ReplaceAll(v1, strings.Trim(v2, "\""), "")
	case "*":
		n := parseInt(v2)
		return strings.Repeat(v1, n)
	case "/":
		n := parseInt(v2)
		return string([]rune(v1)[:len(v1)/n])
	default:
		panic("Неподдерживаемая операция")
	}
}

func parseInt(s string) int {
	var num int
	_, err := fmt.Sscanf(s, "%d", &num)
	if err != nil || num > 10 {
		panic("Неверный формат числа")
	}
	return num
}

func printResult(result string) {
	runeResult := []rune(result)
	if len(runeResult) > 40 {
		runeResult = append(runeResult, []rune("...")...)
	}
	runeResult = []rune("\"" + string(runeResult) + "\"")
	fmt.Println(string(runeResult))
}
