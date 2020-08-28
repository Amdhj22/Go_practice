package main

import (
	"fmt"
)

func main() {
	var a int
	var f float32 = 11.

	fmt.Println("Variable definition (int, float32) :", a, f)

	a = 10
	f = 12.0

	fmt.Println("Var value change (int, float32) :", a, f)

	var i, j, k int = 1, 2, 3

	fmt.Println("Multi-Def var with init :", i, j, k)

	const c int = 10
	const s string = "Hi"

	fmt.Println("const def with init :", c, s)

	const (
		Visa   = "Visa"
		Master = "Master Card"
		Amex   = "American Express"
	)

	fmt.Println("const def with bracket :", Visa, Master, Amex)

	const (
		Apple = iota
		Grape
		Orange
	)

	fmt.Println("const def with bracket using iota :", Apple, Grape, Orange)

	rawLiteral := `good \n good`
	interLiteral := "great \n great"
	fmt.Println("Raw Literal :", rawLiteral)
	fmt.Println("Interpreted Literal :", interLiteral)

	var ni int = 100
	var u uint = uint(ni)
	var nf float32 = float32(ni)
	fmt.Println("Type casting (int, uint, float32) :", ni, u, nf)

}
