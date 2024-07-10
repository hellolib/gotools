package slice

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestChunk(t *testing.T) {
	type args[T any] struct {
		sli       []T
		chunkSize int
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want [][]T
	}
	tests := []testCase[any]{
		{
			name: "Chunk-1",
			args: args[any]{
				sli:       []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				chunkSize: 3,
			},
			want: [][]any{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10},
			},
		},
		{
			name: "Chunk-2",
			args: args[any]{
				sli:       []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				chunkSize: 5,
			},
			want: [][]any{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
			},
		},
		{
			name: "Chunk-3",
			args: args[any]{
				sli:       []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				chunkSize: 10,
			},
			want: [][]any{
				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.args.sli, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceChunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitAsyncExecute(t *testing.T) {
	type args[T any] struct {
		sli         []T
		bulk        int
		concurrency int
		fn          func(sub []T) error
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		wantErr bool
	}
	tests := []testCase[any]{
		{
			name: "SplitAsyncExecute-1",
			args: args[any]{
				sli:         []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				bulk:        3,
				concurrency: 2,
				fn: func(sub []any) error {
					t.Log(sub)
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "SplitAsyncExecute-2",
			args: args[any]{
				sli:         []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				bulk:        3,
				concurrency: 1,
				fn: func(sub []any) error {
					t.Log(sub)
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "SplitAsyncExecute-3",
			args: args[any]{
				sli:         []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				bulk:        3,
				concurrency: 0,
				fn: func(sub []any) error {
					t.Log(sub)
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "SplitAsyncExecute-4",
			args: args[any]{
				sli:         []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				bulk:        0,
				concurrency: 1,
				fn: func(sub []any) error {
					t.Log(sub)
					return nil
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SplitAsyncExecute(tt.args.sli, tt.args.bulk, tt.args.concurrency, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("SplitAsyncExecute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSplitAsyncExecuteWithCtx(t *testing.T) {
	type args[T any] struct {
		ctx         context.Context
		sli         []T
		bulk        int
		concurrency int
		fn          func(sub []T) error
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		wantErr bool
	}

	getTc := func() context.Context {
		to, _ := context.WithTimeout(context.Background(), 1*time.Second)
		return to
	}
	tests := []testCase[any]{
		{
			name: "SplitAsyncExecuteWithCtx-1",
			args: args[any]{
				ctx:         context.Background(),
				sli:         []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				bulk:        3,
				concurrency: 2,
				fn: func(sub []any) error {
					t.Log(sub)
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "SplitAsyncExecuteWithCtx-2",
			args: args[any]{
				sli:         []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				bulk:        3,
				concurrency: 1,
				fn: func(sub []any) error {
					t.Log(sub)
					time.Sleep(9 * time.Millisecond)
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "SplitAsyncExecuteWithCtx-3",
			args: args[any]{
				sli:         []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				bulk:        1,
				concurrency: 1,
				fn: func(sub []any) error {
					t.Log(sub)
					time.Sleep(200 * time.Millisecond)
					return nil
				},
			},
			wantErr: true,
		},
		{
			name: "SplitAsyncExecuteWithCtx-4",
			args: args[any]{
				sli:         []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				bulk:        1,
				concurrency: 2,
				fn: func(sub []any) error {
					time.Sleep(200 * time.Millisecond)
					return nil
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SplitAsyncExecuteWithCtx(getTc(), tt.args.sli, tt.args.bulk, tt.args.concurrency, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("SplitAsyncExecuteWithCtx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSplitExecute(t *testing.T) {
	type args[T any] struct {
		sli  []T
		bulk int
		fn   func(sub []T) error
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		wantErr bool
	}
	tests := []testCase[any]{
		{
			name: "SplitExecute-1",
			args: args[any]{
				sli:  []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				bulk: 3,
				fn: func(sub []any) error {
					t.Log(sub)
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "SplitExecute-2",
			args: args[any]{
				sli:  []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				bulk: 1,
				fn: func(sub []any) error {
					t.Log(sub)
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "SplitExecute-2",
			args: args[any]{
				sli:  []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				bulk: 0,
				fn: func(sub []any) error {
					t.Log(sub)
					return nil
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SplitExecute(tt.args.sli, tt.args.bulk, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("SplitExecute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
