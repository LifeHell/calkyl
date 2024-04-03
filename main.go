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

var r = regexp.MustCompile(`^"(.*?)"\s*([-+])\s*"(.*?)"$`)

func parseInput(input string) (string, string, string) {
	parts := r.FindAllStringSubmatch(input, -1)
	if len(parts) != 1 || len(parts[0]) != 4 {
		panic("Неверный формат выражения")
	}
	return parts[0][1], parts[0][2], parts[0][3]
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
