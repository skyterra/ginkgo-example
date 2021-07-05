package calc

import "errors"

/*
 * 简单计算器实现，仅计算32位有符号整型
 */

func Add(a, b int32) int32 {
	return a + b
}

func Subtract(a, b int32) int32 {
	return a - b
}

func Multiply(a, b int32) int32 {
	return a * b
}

func Divide(a, b int32) (int32, error) {
	if b == 0 {
		return 0, errors.New("divide by 0")
	}

	return a / b, nil
}
