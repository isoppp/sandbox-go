package main

import "fmt"

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
}
