package main

import "fmt"

func main() {

	//===================================================
	// 基本構文
	//===================================================

	// var 変数名 型名で宣言
	var i int
	var str string

	// 型を静的に推論
	var strHello = "hello"

	// 一括代入
	var x, y, z = 10, 20, 30

	// := は var と型定義を省略できる（グローバルでは使用できない）
	hoge := "a"

	// const定数と（）による連続定義
	const(
		fuga = "fuga"
		fugafuga = "fugafuga"
	)

	// iota 連続した型を持たない整数定数値
	const(
		first = iota
		second = iota
		third
	)


	fmt.Println(i, str)
	fmt.Println(strHello)
	fmt.Println(x,y,z)
	fmt.Println(hoge)
	fmt.Println(fuga, fugafuga)
	fmt.Println(first, second, third)
}
