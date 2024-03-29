package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readLine()
	str, operator, num := parseInput(input)
	result := calculate(str, operator, num)
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

func parseInput(input string) (string, string, string) {
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		panic("Неверный формат выражения")
	}
	str := strings.Trim(parts[0], "\"")
	operator := parts[1]
	num := parts[2]
	return str, operator, num
}

func calculate(str, operator, num string) string {
	switch operator {
	case "+":
		return str + strings.Trim(num, "\"")
	case "-":
		return strings.ReplaceAll(str, strings.Trim(num, "\""), "")
	case "*":
		n := parseInt(num)
		return strings.Repeat(str, n)
	case "/":
		n := parseInt(num)
		return str[:len(str)/n]
	default:
		panic("Неподдерживаемая операция")
	}
}

func parseInt(s string) int {
	var num int
	_, err := fmt.Sscanf(s, "%d", &num)
	if err != nil {
		panic("Неверный формат числа")
	}
	return num
}

func printResult(result string) {
	if len(result) > 40 {
		result = result[:40] + "..."
	}
	fmt.Println("\"" + result + "\"")
}
