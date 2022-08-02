package main

import (
	"fmt"
	"os"
)

const tmplStr = `
Q. 以下のコードコンパイルして実行するとどうなるか？
package main
func main() {
	true := false
	println(true == false)
}

1: コンパイルエラー
2: trueと表示される
3: falseと表示される
4: パニックが起きる
`

func main() {

	fmt.Fprint(os.Stdout, tmplStr)

	var answer int

	fmt.Print("回答>")
	fmt.Scanln(&answer)

	if answer == 2 {
		fmt.Println("正解!")
	} else {
		fmt.Println("不正解!")
		fmt.Println("答えは2です")
	}
}
