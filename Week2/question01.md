利用闭包实现一个计数器
```Go
package main

import "fmt"

func counter() func() int {
	a := 0
	return func() int {
		a++
		return a
	}
}

func main() {
	count := counter()
	for i := 0; i <= 10; i++ {
		fmt.Println(count())
	}
}
```
