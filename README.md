# pool

[![Go Walker](https://img.shields.io/badge/Go%20Walker-API%20Documentation-green.svg?style=flat)](https://gowalker.org/github.com/Bluek404/pool)
[![GoDoc](https://img.shields.io/badge/GoDoc-API%20Documentation-blue.svg?style=flat)](http://godoc.org/github.com/Bluek404/pool)

开启指定数量的 goroutine，并限制同时运行的 goroutine 数量

# Example

```go
package main

import (
	"fmt"
	"time"

	p "github.com/Bluek404/pool"
)

func main() {
	pool := p.NewPool(100, 10, func() {
		time.Sleep(time.Second * 1)
		fmt.Print(0)
	})

	pool.Run()
	fmt.Println(pool.Run())
	// Output:
	// already running

	pool.Wait()
	fmt.Println("\nAll done")
}
```