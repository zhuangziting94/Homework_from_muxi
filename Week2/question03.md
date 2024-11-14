  - 比较字符串"hello，世界"的长度 和for range 该字符串的循环次数
```Go
package main

import "fmt"

func main() {
	var str string = "hello, 世界"
	fmt.Println(len(str))
	for i, _ := range str {
		fmt.Printf("循环第 %d 次\n", i+1)
	}
}
```
- 运行结果
```
13
循环第 1 次
循环第 2 次
循环第 3 次
循环第 4 次
循环第 5 次
循环第 6 次
循环第 7 次
循环第 8 次
循环第 11 次
```
