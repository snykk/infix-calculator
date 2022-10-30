package infix_to_postfix

import (
	"strings"

	"github.com/snykk/infix-calculator/stack"
)

var precedence = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
	"^": 3,
}

var isLeftAssociative = map[string]bool{
	"+": true,
	"-": true,
	"*": true,
	"/": true,
	"^": false,
}

func InfixToPostfix(equation string) string {
	results := stack.NewStringStack()
	operators := stack.NewStringStack()

	tokens := strings.Fields(equation)

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		currentPrecedence, isOperator := precedence[token]
		if !isOperator {
			results.Push(token)
			continue
		}

		if len(operators) == 0 {
			operators.Push(token)
			continue
		}

		for i := len(operators) - 1; i >= 0; i-- {
			if precedence[operators[i]] < currentPrecedence || !isLeftAssociative[token] {
				break
			}

			top := operators.Pop()
			results.Push(top)
		}

		operators.Push(token)
	}

	for len(operators) != 0 {
		topOperator := operators.Pop()
		results.Push(topOperator)
	}

	return strings.Join(results, " ")
}
