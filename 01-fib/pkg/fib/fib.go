package fib

import "errors"

// Number - возвращает число фибоначи с номером n.
func Number(n int) (int, error) {

	if n < 1 {
		return 0, errors.New("Вычисление числа фибоначи только для числел больше 0")
	}

	last := 0
	res := 1
	for i := 1; i < n; i++ {
		lastRes := res
		res = res + last
		last = lastRes
	}

	return res, nil
}
