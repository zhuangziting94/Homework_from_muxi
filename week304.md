3. 经典老题：交叉打印下面两个字符串（要求一个打印完，另一个会继续打印）
"ABCDEFGHIJKLMNOPQRSTUVWXYZ" "0123..."
得到："AB01CD23EF34..."
- 先来看数字切片长度大于字母切片长度
```Go
package main

import (
	"fmt"
	"sync"
)

var digitals = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "1", "1", "1", "2", "1", "3", "1", "4", "1", "5", "1", "6", "1", "7", "1", "8", "1", "9", "2", "0"}
var letters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var wait sync.WaitGroup
var tongBu = make(chan struct{})
var ch = make(chan struct{})

func lettersPrint(i int) {
	ch <- struct{}{}
	fmt.Print(letters[i], letters[i+1])
	wait.Done()
	tongBu <- struct{}{}
}

func digitalsPrinnt(i int) {
	<-ch
	fmt.Print(digitals[i], digitals[i+1])
	wait.Done()
}

func main() {
	wait.Add(24)
	for i := 0; i < 24; i += 2 {
		go lettersPrint(i)
		go digitalsPrinnt(i)
		<-tongBu
	}
	wait.Wait()
	fmt.Println()
	fmt.Println(digitals[24:])
}
```
打印输出结果：
```Go
AB01CD23EF45GH67IJ89KL11MN12OP13QR14ST15UV16WX17
[1 8 1 9 2 0]
```
- emmmmm,字母切片长度大于数字切片长度的情况有点复杂，这里就先不演示了
