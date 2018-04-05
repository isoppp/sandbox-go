package main

import (
	"fmt"
	"strconv"
)

//===================================================
// 関数・ポインタ・制御構文
//===================================================

// 関数
// func 関数名（引数変数 引数型）戻り値の型 {}
func add(x int, y int) int {
	return x + y
}

// 可変長引数
func fArgs(strArgs ...string) {
	for index, value := range strArgs {
		fmt.Println(index, value)
	}
}

// 複数の戻り値
func multiReturn(x int) (int, string) {
	return x, "multi return example"
}

// 関数型
type funcTemplate func(string) string

func greet(name string) string {
	return "hello, " + name
}

func typeTemplate(f funcTemplate) {
	fmt.Println(greet("name"))
}

// 関数の引数の渡し方とポインタ
func notPointerInc(x int) {
	x++
	fmt.Println(x)
}

func pointerInc(x *int) {
	*x++
	fmt.Println(*x)
}

//===================================================
// 構造体
//===================================================

// 構造体定義
type User struct {
	name string
	age  int
}

// ネスト構造体
type Account struct {
	User
	email string
}

// Goでは初期化関数を一緒に作ることが一般的？
func newUser(name string, age int) *User {
	u := new(User)
	u.name = name
	u.age = age
	return u
}

// 構造体のメソッド
func (u User) greetUser() string {
	return "hello" + u.name
}

//===================================================
// interface(メソッドの型だけを記述した型)
//===================================================
type Car interface {
	run(int) string
	stop()
}

type MyCar struct {
	name  string
	speed int
}

func (u *MyCar) run(speed int) string {
	u.speed = speed
	return strconv.Itoa(speed) + "kmで走ります"
}

func (u *MyCar) stop() {
	fmt.Println("停止します")
	u.speed = 0
}

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
	const (
		fuga     = "fuga"
		fugafuga = "fugafuga"
	)

	// iota 連続した型を持たない整数定数値
	const (
		first  = iota
		second = iota
		third
	)

	fmt.Println(i, str)
	fmt.Println(strHello)
	fmt.Println(x, y, z)
	fmt.Println(hoge)
	fmt.Println(fuga, fugafuga)
	fmt.Println(first, second, third)

	//===================================================
	// 関数・ポインタ の 実行
	//===================================================
	fmt.Println(add(1, 2))
	fArgs("hogehoge", "fugafuga", "unnunn")
	multiNum, multiStr := multiReturn(100)
	fmt.Println(multiNum, multiStr)
	typeTemplate(greet)
	pointerExample := 10
	notPointerInc(pointerExample)
	fmt.Println(pointerExample)
	pointerInc(&pointerExample)
	fmt.Println(pointerExample)

	//===================================================
	// 制御構文
	//===================================================

	// if and if local var
	boolVariable := 10
	if boolVariable == 10 {
		println("if basic")
	}

	if ifLocalValue := add(1, 2); ifLocalValue == 1 {

	} else {
		fmt.Println("if scope local value")
	}

	// for
	forSum := 0
	for i := 0; i < 10; i++ {
		forSum += i
		fmt.Println("for", forSum)
	}

	// like while
	likeWhile := 3
	for likeWhile > 0 {
		likeWhile--
		fmt.Println("like while:", likeWhile)
	}

	// switch breakがデフォルト breakさせない場合はfallthrough
	switchVar := 1
	switch switchVar {
	case 1:
		fmt.Println("switchVar == 1")
		fallthrough
	case 2, 3, 4:
		fmt.Println("switchVar == 2 or 3 or 4")
	default:
		fmt.Println("switchVar is another value")
	}

	//===================================================
	// 構造体
	//===================================================

	// 構造体の初期化1
	var u User
	u.name = "tarou"
	u.age = 30
	fmt.Println(u.name, u.age)

	// 構造体の初期化2
	u2 := User{"tarou", 30}
	fmt.Println(u2.name, u2.age)

	// 構造体の初期化3
	u3 := User{name: "tarou", age: 30}
	fmt.Println(u3.name, u3.age)

	// ポインタ型 up1はポインタ型
	up1 := &User{"tarou", 30}
	fmt.Println(up1.name, up1.age)

	// ポインタ型2 up2はポインタ型
	up2 := new(User)
	up2.name = "tarou"
	up2.age = 30
	fmt.Println(up2.name, up2.age)

	// Goでは初期化関数を一緒に作ることが一般的？
	uu := newUser("tarou", 30)
	fmt.Println(uu.name, uu.age)

	// ネスト構造体
	account := Account{User{"tarou", 30}, "hogehoge@example.com"}
	fmt.Println(account.User.name, account.User.age, account.email)

	// 構造体のメソッド
	fmt.Println(u.greetUser())

	//===================================================
	// interface
	//===================================================

	myCar := &MyCar{name: "マイカー", speed: 0}
	var car Car = myCar
	fmt.Println(car.run(50))
	car.stop()

	// 空のインターフェース
	// どの型のデータも代入可能
	var xxx interface{}
	num := 0
	strrr := "str"

	xxx = num
	xxx = strrr

	//	型を判定する
	if value, ok := xxx.(int); ok {
		fmt.Println("value is int", value)
	} else if value, ok := xxx.(string); ok {
		fmt.Println("value is string", value)
	} else {
		fmt.Println("value is other type")
	}

	xxx = num

	switch value := xxx.(type) {
	case int:
		fmt.Println("value is int", value)
	case string:
		fmt.Println("value is string", value)
	default:
		fmt.Println("value is other type")
	}

	//===================================================
	// よく使うデータ型
	//===================================================

	nums := []int{2, 3, 4}

	// forループ
	for i := 0; i < len(nums); i++ {
		fmt.Print(fmt.Sprintf("index: %d, value: %d\n", i, nums[i]))
	}

	// range
	for i, v := range nums {
		fmt.Print(fmt.Sprintf("index: %d, value: %d\n", i, v))
	}

	// Array [n]type [5]int と [10]int は違う型
	// 配列同士の代入はポインタではなく値渡しとなる
	var arr [5]int
	fmt.Println("arr", arr)
	arr[4] = 100
	fmt.Println("arr", arr)
	fmt.Println("arr[4]", arr[4])

	arr2 := [5]int{5, 6, 7, 8, 9}
	fmt.Println("arr2[:]", arr2[:])
	fmt.Println("arr2[1:4]", arr2[1:4])
	fmt.Println("arr2[:4]", arr2[:4])
	fmt.Println("arr2[1:]", arr2[1:])

	// Slice サイズを明示的に持たず柔軟なため使いやすい
	// ただ細かいメモリの割当の制御ができないのでそういったケースではArrayを使ったほうが良い
	// こちらは代入した場合にも値ではなくポインタが渡される
	// [n]type がArray型（長さによって型が異なる） / []type がslice型 ということが重要

	var slice1 []int
	slice2 := []string{"a", "b", "c"}
	fmt.Println(slice1)
	fmt.Println(slice2)

	// Map

	var map1 map[string]int = make(map[string]int)
	map1["x"] = 10
	map1["y"] = 100
	fmt.Println(map1)
	fmt.Println(map1["x"])

	map2 := map[string]int{"x": 10, "y": 1000}
	fmt.Println(map2)
	fmt.Println(map2["y"])

	// mapには2つ戻り値があり、値が存在するかしないかのboolが返ってくる
	mapval, mapok := map2["xxx"]
	fmt.Println(mapval, mapok)
}
