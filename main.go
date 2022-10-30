package main

import (
	"bufio"
	"fmt"
	"os"

	postfixCalc "github.com/snykk/infix-calculator/postfix_calculator"
)

func main() {
	fmt.Print("~> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	equation := scanner.Text()
	result, err := postfixCalc.PostfixCalculate(equation)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("result: ", result)
}
