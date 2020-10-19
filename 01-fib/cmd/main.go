package main

import (
	"flag"
	"fmt"

	"1.go.core/pkg/fib"
)

func main() {
	inputNum := flag.Int("num", 10, " номер числа Фибоначи должен быть в диапазоне от 0 до 20")
	flag.Parse()
	num := *inputNum
	if num > 20 {
		fmt.Printf("номер числа фибоначи не должен превышать 20")
		return
	}
	fmt.Printf("Расчет числа Фибоначи для n=%v\n", num)
	res, err := fib.Number(num)
	if err != nil {
		fmt.Printf("ошибка при расчете числа Фибоначи: %v", err)
		return
	}
	fmt.Printf("Число Фибоначи с номером %d = %v\n", num, res)
}
