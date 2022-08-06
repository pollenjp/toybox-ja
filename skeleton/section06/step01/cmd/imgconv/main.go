package main

import (
	"flag"
	"fmt"
	"os"

	// imgconvパッケージをインポートする
	imgconv "github.com/pollenjp/toybox-ja/skeleton/section06/step01"
)

var (
	flagTo   = imgconv.PNG
	flagFrom = imgconv.JPEG
)

func init() {
	flag.Var(&flagTo, "to", "after format")
	flag.Var(&flagFrom, "from", "before format")
}

func main() {
	if err := imgconv.ConvertAll(os.Args[1], flagFrom, flagTo); err != nil {
		fmt.Fprintln(os.Stderr, "エラー:", err)
		os.Exit(1)
	}
}
