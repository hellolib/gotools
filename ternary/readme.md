# ternary

## Feature

### 1. 三元表达式
- demo
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
