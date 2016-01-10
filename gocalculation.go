package gocalculation

import (
	"strconv"
	"strings"
)

func Cal(formula string, sourceVals map[string]float64) float64 {

	form := formula
	form = strings.Replace(form, "+", " + ", -1)
	form = strings.Replace(form, "-", " - ", -1)
	form = strings.Replace(form, "*", " * ", -1)
	form = strings.Replace(form, "/", " / ", -1)
	form = strings.Replace(form, "(", " ( ", -1)
	form = strings.Replace(form, ")", " ) ", -1)

	stack := NewStack()
	fields := strings.Fields(form)
	rpn := make([]string, 0)
	for _, field := range fields {
		if isOperator(field) {

			if !stack.Empty() {

				peak := stack.Peak().(string)
				for peak != "(" && comparePriority(field, peak) <= 0 {
					rpn = append(rpn, stack.Pop().(string))
					if stack.Empty() {
						break
					}

					peak = stack.Peak().(string)
				}
			}

			stack.Push(field)
		} else if field == "(" {
			stack.Push(field)
		} else if field == ")" {
			pop := stack.Pop().(string)
			for pop != "(" && !stack.Empty() {
				rpn = append(rpn, pop)
				pop = stack.Pop().(string)
			}
		} else {
			rpn = append(rpn, field)
		}
	}

	for !stack.Empty() {
		rpn = append(rpn, stack.Pop().(string))
	}

	// 对rpn进行求值
	for _, rp := range rpn {
		if isOperator(rp) {
			var ret float64
			calR := parseFloat64(stack.Pop())
			calL := parseFloat64(stack.Pop())
			switch rp {
			case "+":
				ret = calL + calR
			case "-":
				ret = calL - calR
			case "*":
				ret = calL * calR
			case "/":
				if calR == 0 {
					ret = 0
				} else {
					ret = calL / calR
				}
			}

			stack.Push(ret)
		} else {

			var val float64
			if _, ok := sourceVals[rp]; ok {
				val = sourceVals[rp]
			} else {
				val = stringToFloat64(rp)
			}
			stack.Push(val)
		}
	}

	result := stack.Pop().(float64)

	return result
}

func isOperator(c string) bool {

	return c == "+" || c == "-" || c == "*" || c == "/"
}

func comparePriority(op1 string, op2 string) int {

	if op1 == "*" || op1 == "/" {
		if op2 == "*" || op2 == "/" {
			return 0
		} else {
			return 1
		}
	} else {
		if op2 == "*" || op2 == "/" {
			return -1
		} else {
			return 0
		}
	}
}

func parseFloat64(val interface{}) float64 {

	var re float64
	switch val.(type) {
	case int64:
		re = float64(val.(int64))
	case float64:
		re = val.(float64)
	case int:
		re = float64(val.(int))
	case string:
		re = stringToFloat64(val.(string))
	default:
		re = 0
	}

	return re
}

func stringToFloat64(str string) float64 {

	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		f = 0.0
	}

	return f
}
