package slice

import (
	"context"
	"golang.org/x/sync/errgroup"
)

// Chunk 将切片切分为大小为chunkSize的切片
// chunkSize <= 0 时返回空切片
func Chunk[T any](sli []T, chunkSize int) [][]T {
	sliLen := len(sli)
	if sliLen == 0 || chunkSize <= 0 {
		return [][]T{}
	}
	chunked := make([][]T, 0, (sliLen+chunkSize-1)/chunkSize)
	for i := 0; i < len(sli); i += chunkSize {
		end := i + chunkSize
		if end > len(sli) {
			end = len(sli)
		}
		chunked = append(chunked, sli[i:end])
	}
	return chunked
}

// SplitExecute 切片分块执行
// bulk : 分块大小, <= 0 时不执行 fn
func SplitExecute[T any](sli []T, bulk int, fn func(sub []T) error) error {
	l := len(sli)
	if l == 0 {
		return nil
	}

	if bulk <= 0 {
		return nil
	}

	if l <= bulk {
		return fn(sli)
	}

	for i := 0; i < len(sli); i += bulk {
		end := i + bulk
		if end > len(sli) {
			end = len(sli)
		}

		err := fn(sli[i:end])
		if err != nil {
			return err
		}
	}

	return nil
}

// SplitAsyncExecute 切片分块异步执行
// bulk : 分块大小
// concurrency: 并发数, concurrency <= 0 不执行 fn
func SplitAsyncExecute[T any](sli []T, bulk int, concurrency int, fn func(sub []T) error) error {
	return splitAsyncExecuteWithCtx(context.Background(), sli, bulk, concurrency, fn)
}

// SplitAsyncExecuteWithCtx 切片分块异步执行
// bulk : 分块大小
// concurrency: 并发数, concurrency <= 0 不执行 fn
func SplitAsyncExecuteWithCtx[T any](ctx context.Context, sli []T, bulk int, concurrency int, fn func(sub []T) error) error {
	return splitAsyncExecuteWithCtx(ctx, sli, bulk, concurrency, fn)
}

func splitAsyncExecuteWithCtx[T any](ctx context.Context, sli []T, bulk int, concurrency int, fn func(sub []T) error) error {
	l := len(sli)
	if l == 0 || concurrency <= 0 {
		return nil
	}

	if bulk <= 0 {
		return nil
	}

	if l <= bulk {
		return fn(sli)
	}

	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(concurrency)

	for i := 0; i < len(sli); i += bulk {
		start := i
		end := start + bulk
		if end > len(sli) {
			end = len(sli)
		}

		g.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				return fn(sli[start:end])
			}
		})
	}

	return g.Wait()
}
