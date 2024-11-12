1.在tool目录下用go mod init muxi-backend/tool,然后go get github.com/spf13/afero
2.补全请求GET代码，定义err := Body.Close(),如果有错误则返回，
注意到Read()输入是（p []byte),
3.secret/main.go中GET请求是加密的论文和解密的密钥存放的根目录，我需要把它俩分开
* 这里先进行加密的论文
4. 从 HTTP 响应体中读取数据并关闭响应体
当使用 http.Get 发起 HTTP 请求时，返回的 resp.Body 就是一个 io.ReadCloser，表示 HTTP 响应的主体。
可以读取响应体内容，然后关闭响应体。
用buf,[]byte类型来缓冲读取响应内容，用n指代response.Body.Read(buf)读取后的切片长度
5.将buf[:n]强行转换成string类型，并赋给encodedPaper
* 复制操作得到密钥


6.解密环节涉及到调用包，调用包后可以运行函数？用go.work?
  save环节同样使用了别包的函数
  go work init ./secret ./tool,引用包名“muxi-backend/tool/safePaper"
