/*
Большие числа и операции
Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).
Комментарий: в Go тип int справится с такими числами,
но обратите внимание на возможное переполнение для ещё больших значений.
Для очень больших чисел можно использовать math/big.
*/
package main

import (
	"fmt"
	"math/big"
)

func Add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func Sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func Mul(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func Div(a, b *big.Int) *big.Int {
	if b.Cmp(big.NewInt(0)) == 0 {
		return nil
	}
	return new(big.Int).Div(a, b)
}

func main() {
	a := big.NewInt(1 << 25)
	b := big.NewInt(1 << 22)

	fmt.Println("Add:", Add(a, b))
	fmt.Println("Sub:", Sub(a, b))
	fmt.Println("Mul:", Mul(a, b))
	fmt.Println("Div:", Div(a, b))
}
