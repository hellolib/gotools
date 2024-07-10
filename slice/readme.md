# slice

## Feature
### 1. Chunk 将切片切分为大小为chunkSize的切片
- demo
  ```go
  fmt.Println(slice.Chunk([]int{1, 2, 3, 4, 5}, 2))
  // [[]int{1, 2}, []int{3, 4}, []int{5}]
  ```
### 2. SplitExecute 切片分块执行
  - demo
    ```go
    fmt.Println(slice.SplitExecute([]int{1, 2, 3, 4, 5}, 2, func(sub []int) error {
        fmt.Println(sub)
        return nil
    }))
    // [1 2]
    // [3 4]
    // [5]
    // nil
    ```
### 3. SplitAsyncExecute 切片分块异步执行
  - demo
    ```go
    fmt.Println(slice.SplitAsyncExecute([]int{1, 2, 3, 4, 5}, 2, 2, func(sub []int) error {
        fmt.Println(sub)
        return nil
    }))
    // [3 4]
    // [1 2]
    // [5]
    // nil
    ```
  
### 4. SplitAsyncExecuteWithCtx 切片分块异步执行
  - demo
    ```go
    fmt.Println(slice.SplitAsyncExecuteWithCtx(context.Background(), []int{1, 2, 3, 4, 5}, 2, 2, func(sub []int) error {
    fmt.Println(sub)
    return nil
    }))
    // [3 4]
    // [5]
    // [1 2]
    // nil
    ```

