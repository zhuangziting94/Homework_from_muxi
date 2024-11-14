```Go
  func main() {
      x := "hello!"
      for i := 0; i < len(x); i++ {
          x := x[i] 
          if x != '!' {
              x := x + 'A' - 'a'
              fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
         }
     }
  }
```
- line5: 把x长字符串的单个元素重新赋给x
- line7: 将x大写后重新赋给x
