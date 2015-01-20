# pool

[![Go Walker](https://img.shields.io/badge/Go%20Walker-API%20Documentation-green.svg?style=flat)](https://gowalker.org/github.com/Bluek404/pool)
[![GoDoc](https://img.shields.io/badge/GoDoc-API%20Documentation-blue.svg?style=flat)](http://godoc.org/github.com/Bluek404/pool)

goroutine执行完毕后，自动补充goroutine

# Example

```go
package main

import (
	"fmt"
	"time"

	p "github.com/Bluek404/pool"
)

func main() {
	pool := p.NewPool(10, func() {
		time.Sleep(time.Millisecond * 100)
		fmt.Print(0)
	})

	fmt.Println("Max Thread:", pool.GetMaxThread())
	pool.Run()
	time.Sleep(time.Second * 1)
	pool.Exit()
}
```