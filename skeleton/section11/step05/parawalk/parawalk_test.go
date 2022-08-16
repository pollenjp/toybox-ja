package parawalk_test

import (
	"context"
	"io/fs"
	"path/filepath"
	"sync"
	"testing"

	"github.com/pollenjp/toybox-ja/skeleton/section11/step05/parawalk"
)

func skip(paths ...string) parawalk.WalkFunc {
	return func(ctx context.Context, path string, info fs.FileInfo, err error) error {
		for i := range paths {
			if path == paths[i] {
				return filepath.SkipDir
			}
		}
		return nil
	}
}

func TestWalk(t *testing.T) {
	defaultFn := parawalk.WalkFunc(func(_ context.Context, _ string, _ fs.FileInfo, _ error) error { return nil })
	cases := []struct {
		name string
		ctx  context.Context
		fn   parawalk.WalkFunc
	}{
		{"all", context.Background(), defaultFn},
		{"skipb", context.Background(), skip("testdata/b")},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// diffという名前のsync.Map型の変数を宣言
			var diff sync.Map

			filepath.Walk("testdata", func(path string, info fs.FileInfo, err error) error {
				// diffにキーがpathで値struct{}{}をストアする
				diff.Store(path, struct{}{})

				return tt.fn(context.Background(), path, info, err)
			})

			var unexpectedPaths sync.Map
			parawalk.Walk(tt.ctx, "testdata", func(ctx context.Context, path string, info fs.FileInfo, err error) error {
				_, loaded := diff.LoadAndDelete(path)
				if !loaded {
					unexpectedPaths.Store(path, struct{}{})
				}
				return tt.fn(ctx, path, info, err)
			})

			// diffマップに要素がある場合はRangeメソッドで要素を回ってエラーを出す
			// ヒント：unexpectedPathsの処理を参考にする
			diff.Range(
				func(path, _ interface{}) bool {
					t.Errorf("unexpected path: %s", path)
					return true
				},
			)

			unexpectedPaths.Range(func(path, _ interface{}) bool {
				t.Errorf("walked to unexpected path: %v", path)
				return true
			})
		})
	}
}
