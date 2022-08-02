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

	for counter := 1; counter <= 2; counter++ {
		var answer int
		fmt.Print("回答>")
		fmt.Scanln(&answer)

		switch {
		// answerが2の時（正解の場合）
		case answer == 2:
			fmt.Println("正解!")
			// 1度目で正解しても2回答える必要がある
		// 1度目のチャレンジで不正解の場合
		case counter == 1:
			fmt.Println("不正解!")
			fmt.Println("もう一度チャレンジ!")
		// それ以外（2度目のチャレンジで不正解の場合）
		case counter == 2:
			fmt.Println("不正解!")
			fmt.Println("答えは2です")
		}
	}
}
