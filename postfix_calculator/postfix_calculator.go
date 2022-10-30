package calculator

import (
	"errors"
	"math"
	"strconv"
	"strings"

	infToPost "github.com/snykk/infix-calculator/infix_to_postfix"
	"github.com/snykk/infix-calculator/stack"
)

func PostfixCalculate(equation string) (float64, error) {
	equation = strings.TrimSpace(equation)
	if equation == "" {
		return 0, errors.New("equation cannot be empty")
	}

	c := infToPost.InfixToPostfix(equation)
	tokens := strings.Fields(c)

	results := stack.NewFloatStack()
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/", "^":
			if len(results) < 2 {
				return 0, errors.New("missing required operand")
			}

			right := results.Pop()
			left := results.Pop()
			var res float64
			switch token {
			case "+":
				res = left + right
			case "-":
				res = left - right
			case "*":
				res = left * right
			case "/":
				res = left / right
			case "^":
				res = math.Pow(left, right)
			}
			results.Push(res)
		default:
			floatNum, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, errors.New("invalid operand")
			}
			results.Push(floatNum)
		}
	}

	return results[0], nil
}
