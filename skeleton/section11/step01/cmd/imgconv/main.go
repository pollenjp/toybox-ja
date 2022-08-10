package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"runtime/trace"

	imgconv "github.com/pollenjp/toybox-ja/skeleton/section11/step01"
)

var (
	flagTo   = imgconv.PNG
	flagFrom = imgconv.TIFF
)

func init() {
	flag.Var(&flagTo, "to", "after format")
	flag.Var(&flagFrom, "from", "before format")
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "エラー:", err)
		os.Exit(1)
	}
}

func run() (rerr error) {
	f, err := os.Create("trace.out")
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil && rerr != nil {
			rerr = err
		}
	}()

	// 出力先はファイルfとしtrace.Startを呼ぶ
	if err = trace.Start(f); err != nil {
		panic(err)
	}
	defer trace.Stop()

	ctx, task := trace.NewTask(context.Background(), "imgconv")

	// deferでtask.Endを呼ぶ
	defer task.End()

	if err := imgconv.ConvertAll(ctx, os.Args[1], flagFrom, flagTo); err != nil {
		return err
	}

	return nil
}
