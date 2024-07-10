# GoTools

go语言开发实用工具。

## Import
- 安装
    ```shell
    go get github.com/hellolib/gotools
    ```
- 引入
    
    ```go
    import "github.com/hellolib/gotools"
    ```

## Example

```go
package main

import (
  "fmt"
  "github.com/hellolib/gotools/ternary"
)

func main() {
  fmt.Println(ternary.If(true, "true", "false"))  // true
  fmt.Println(ternary.If(false, "true", "false")) // false
  fmt.Println(ternary.If(1 > 0, 1, 0))  // 1
  fmt.Println(ternary.If(1 < 0, 1, 0))  // 0
}
```

## Feature

* #### [三元表达式](./ternary/readme.md)
  * If 三元运算
* #### [切片](./slice/readme.md)
  * Chunk 将切片切分为大小为chunkSize的切片
  * SplitExecute 切片分块执行
  * SplitAsyncExecute 切片分块异步执行
  * SplitAsyncExecuteWithCtx 切片分块异步执行