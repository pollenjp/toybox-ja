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

LOOP:
	for count := 1; count <= 2; count++ {
		var answer int
		for {
			fmt.Print("回答>")
			fmt.Scanln(&answer)
			if answer >= 1 && answer <= 4 {
				break
			}
			fmt.Println("1から4で入力してえください")
		}

		switch {
		case answer == 2:
			fmt.Println("正解!")
			// ラベルLOOPのついた繰り返しを抜け出す
			break LOOP

		case count == 1:
			fmt.Println("不正解!")
			fmt.Println("もう一度チャレンジ!")
		default:
			fmt.Println("不正解!")
			fmt.Println("答えは2です")
		}
	}
}
