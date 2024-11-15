1.在tool目录下用go mod init muxi-backend/tool,然后go get github.com/spf13/afero
2.补全请求GET代码，定义err := Body.Close(),如果有错误则返回，
注意到Read()输入是（p []byte),
3.secret/main.go中GET请求是加密的论文和解密的密钥存放的根目录，我需要把它俩分开
> 这里先进行加密的论文
4. 从 HTTP 响应体中读取数据并关闭响应体,`defer func(Body io.ReadCloser)`是匿名函数，遇到错误`return`,这里`io.ReadCloser`是接口， 
当使用 http.Get 发起 HTTP 请求时，返回的 response.Body 就是一个 io.ReadCloser，表示 HTTP 响应的主体。
可以读取响应体内容，然后关闭响应体。
用buf,[]byte类型来缓冲读取响应内容，用n指代response.Body.Read(buf)读取后的切片长度
5. 将buf[:n]强行转换成string类型，并赋给encodedPaper
> 复制操作得到密钥key


6. 解密环节涉及到引入包，引入包后可以运行函数,用go.work,
   save环节同样使用了别包的函数  
  
7. 终端运行go work init ./secret ./tool,引用包名“muxi-backend/tool/safePaper"
## 最后main.go代码 ##
```Go
package main

import (
	"io"
	"log"
	"muxi-backend/tool/getDecryptedPaper"
	"muxi-backend/tool/savePaper"
	"net/http"
)

func main() {
	// 目标根URL
	url := "http://121.43.151.190:8000/"
	// 发送 GET 请求,返回的结果还需要进行处理才能得到你需要的结果
	response, err := http.Get(url + "paper")
	if err != nil {

		log.Fatalf("Failed to send request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}

	}(response.Body)
	buf := make([]byte, 1024)
	n, _ := response.Body.Read(buf)
	encodedPaper := string(buf[:n])

	resp, err := http.Get(url + "secret")
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}

	}(resp.Body)
	buf2 := make([]byte, 1024)
	n, _ = resp.Body.Read(buf2)
	key := string(buf2[:n])

	//解密
	data := getDecryptedPaper.GetDecryptedPaper(encodedPaper, key)
	//将文件内容保存到相关路径中
	savePaper.SavePaper("C:\\Users\\zhuangziting\\Downloads\\muxi-backend\\paper\\Academician Sun's papers.txt", data)
}
```
![image](https://github.com/user-attachments/assets/f2f072f7-8b72-4ef7-9287-6f32a08513f7)

### go work 的用法
**应用场景：同一个代码仓库里有多个互相依赖的Go Module**
当我们在同一个代码仓库里开发多个互相依赖的Go Module时，我们可以使用go.work，而不是在go.mod里使用replace指令。  
为你的workspace(工作区)创建一个目录。  
Clone仓库里的代码到你本地。代码存放的位置不一定要放在工作区目录下，因为你可以在go.work里使用use指令来指定Module的相对路径。在工作区目录运行  `go work init [path-to-module-one] [path-to-module-two]` 命令。  
示例: 你正在开发 example.com/x/tools/groundhog 这个Module，该Module依赖 example.com/x/tools 下的其它Module。你Clone仓库里的代码到你本地，然后在工作区目录运行命令 `go work init tools tools/groundhog`   
go.work文件内容如下：
```Go
go 1.18

use (
./tools
./tools/groundhog
)
```
