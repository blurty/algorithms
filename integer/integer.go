package integer

import (
	"errors"
	"fmt"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

// Divide two integers without using multiplication, division and mod operator.
// If it is overflow, return MAX_INT.
func Divide(dividend int, divisor int) int {
	if divisor == 0 {
		panic("divisor is zero")
	}
	if dividend == MinInt && divisor == -1 {
		return MaxInt
	}

	flag := true
	if dividend < 0 && divisor > 0 || dividend > 0 && divisor < 0 {
		flag = false
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	ans := 0
	for dividend >= divisor {
		x := divisor
		y := 1
		for dividend >= (x << 1) {
			x <<= 1
			y <<= 1
		}
		dividend -= x
		ans += y
	}
	if flag {
		return ans
	} else {
		return -ans
	}
}

// a simple arithmetic device
func Arithmetic(expression string) (int, error) {
	dS := make([]int, 0, len(expression))  // store datas
	sS := make([]byte, 0, len(expression)) // store signs
	ret := 0                               // store result
	var x, y int                           // operand
	var op, op2 byte                       // operator
	for i := 0; i < len(expression); {
		switch {
		case expression[i] == '(':
			i++
			sS = append(sS, '(')
		case expression[i] == ')':
			i++
			for {
				op = sS[len(sS)-1]
				if op == '(' {
					sS = sS[:len(sS)-1]
					break
				} else {
					y, x = dS[len(dS)-1], dS[len(dS)-2]
					dS = dS[:len(dS)-2]
					op = sS[len(sS)-1]
					sS = sS[:len(sS)-1]
					ret = calculate(x, y, op)
					dS = append(dS, ret)
				}
			}
		case expression[i] >= '0' && expression[i] <= '9':
			n, err := fmt.Sscanf(expression[i:], "%d", &x)
			if n == 0 || err != nil {
				return ret, errors.New("invalid expression")
			}
			dS = append(dS, x)
			count := 0
			for x != 0 {
				count++
				x /= 10
			}
			i += count
		default:
			op = expression[i]
			i++
			for {
				if len(sS) == 0 || sS[len(sS)-1] == '(' {
					sS = append(sS, op)
					break
				}
				op2 = sS[len(sS)-1]
				if optMax(op, op2) > 0 {
					sS = append(sS, op)
					break
				} else {
					y, x = dS[len(dS)-1], dS[len(dS)-2]
					dS = dS[:len(dS)-2]
					sS = sS[:len(sS)-1]
					ret = calculate(x, y, op2)
					dS = append(dS, ret)
				}
			}
		}
	}
	op = sS[len(sS)-1]
	if op == '*' || op == '/' {
		y, x = dS[len(dS)-1], dS[len(dS)-2]
		dS = dS[:len(dS)-2]
		sS = sS[:len(sS)-1]
		ret = calculate(x, y, op)
		dS = append(dS, ret)
	}
	y, x = dS[len(dS)-1], dS[len(dS)-2]
	dS = dS[:len(dS)-2]
	op = sS[len(sS)-1]
	sS = sS[:len(sS)-1]
	ret = calculate(x, y, op)
	return ret, nil
}

// calculate x op y
func calculate(x, y int, op byte) int {
	switch op {
	case '+':
		return x + y
	case '-':
		return x - y
	case '*':
		return x * y
	case '/':
		return x / y
	}
	return 0
}

// return priority of op1 and op2
// if op1 > op2, return 1, equals return 0, otherwise, return -1
func optMax(op1, op2 byte) int {
	if op1 == '*' || op1 == '/' {
		if op2 == '+' || op2 == '-' {
			return 1
		}
		return 0
	} else {
		if op2 == '+' || op2 == '-' {
			return 0
		}
		return -1
	}
}
