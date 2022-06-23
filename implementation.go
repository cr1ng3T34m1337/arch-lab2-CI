package lab2

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// isOperator ckecks if provided string is operator
func isOperator(str string) bool {
	operators := []string{"+", "-", "*", "/", "^"}
	for _, operator := range operators {
		if str == operator {
			return true
		}
	}
	return false
}

// prefixValidate checks if provided expression is valid
func prefixValidate(expr []string) bool {
	if !isOperator(expr[0]) {
		return false
	}
	var (
		operatorCount int
		operandCount  int
	)
	for _, str := range expr {
		if isOperator(str) {
			operatorCount++
		} else {
			operandCount++
		}
	}
	return operandCount == operatorCount+1
}

// PrefixEvaluate evaluates given prefix expression.
// Operators and operands must be separated by space
func PrefixEvaluate(expr string) (float64, error) {
	var stack Stack
	values := strings.Split(expr, " ")
	res := prefixValidate(values)
	if !res {
		return 0.0, errors.New("invalid input expression")
	}
	for i := len(values) - 1; i >= 0; i-- {
		value := values[i]
		if num, err := strconv.ParseFloat(value, 64); err == nil {
			stack.Push(num)
		} else {
			operand1, err1 := stack.Pop()
			operand2, err2 := stack.Pop()
			if err1 != nil || err2 != nil {
				return 0.0, errors.New("something went wrong, maybe check your input")
			}
			switch value {
			case "+":
				stack.Push(operand1 + operand2)
			case "-":
				stack.Push(operand1 - operand2)
			case "*":
				stack.Push(operand1 * operand2)
			case "/":
				stack.Push(operand1 / operand2)
			case "^":
				stack.Push(math.Pow(operand1, operand2))
			default:
				return 0.0, errors.New("available operators: +, -, *, /, ^")
			}
		}
	}
	return stack.Pop()
}
