1.将程序输入VScode运行，发现error:slice bounds out of range [:5] with capacity 3,说明有切片越界，仔细阅读代码发现
```Go
 if len(consumeMSG) > 0 {
             //进行异步消费 
             go func() {
                m := consumeMSG[:ConsumeNum]
                fn(m)
             }()
             // 更新上次消费时间
             lastConsumeTime = time.Now()
             // 清空插入的数据
             consumeMSG = consumeMSG[ConsumeNum:]
          }
       }
```
代码中切片consumMSG的长度可能不足5，无法把容量只有3的切片的前五个元素赋给m,也无法将序号5之后的元素清空  
2. for range 语句中两个gorountine并发执行，同时对consumeMSG修改，抢占数据，需要对每个gorountine上互斥锁  
3. 此时运行代码会发现一直输出本次消费了0条消息，说明在协程并发时，协程外面的代码的consumeMSG直接被go语句使用，导致混乱，所以用管道阻塞读取条件将
使子协程运行完再运行主协程。  
综上，代码更改如下：
```Go
package main

import (
	"fmt"
	"sync"
	"time"
)

type message struct {
	Topic     string
	Partition int32
	Offset    int64
}

type FeedEventDM struct {
	Type    string
	UserID  int
	Title   string
	Content string
}

type MSG struct {
	ms        message
	feedEvent FeedEventDM
}

const ConsumeNum = 5

func main() {
	var consumeMSG []MSG
	var lastConsumeTime time.Time // 记录上次消费的时间
	msgs := make(chan MSG)
	var ch = make(chan struct{})
	//这里源源不断的生产信息
	go func() {
		for i := 0; ; i++ {
			msgs <- MSG{
				ms: message{
					Topic:     "消费主题",
					Partition: 0,
					Offset:    0,
				},
				feedEvent: FeedEventDM{
					Type:    "grade",
					UserID:  i,
					Title:   "成绩提醒",
					Content: "您的成绩是xxx",
				},
			}

			//每次发送信息会停止0.01秒以模拟真实的场景
			time.Sleep(100 * time.Millisecond)
		}
		close(msgs)
	}()
	var lock sync.Mutex
	//不断接受消息进行消费
	for msg := range msgs {
		// 添加新的值到events中
		consumeMSG = append(consumeMSG, msg)
		// 如果数量达到额定值就批量消费

		if len(consumeMSG) >= ConsumeNum {
			//进行异步消费
			go func() {
				ch <- struct{}{}
				lock.Lock()
				m := consumeMSG[:ConsumeNum]
				fn(m)
				lock.Unlock()
			}()
			<-ch
			// 更新上次消费时间
			lastConsumeTime = time.Now()
			// 清除插入的数据
			consumeMSG = consumeMSG[ConsumeNum:]
		} else if !lastConsumeTime.IsZero() && time.Since(lastConsumeTime) > 5*time.Minute {
			// 如果距离上次消费已经超过5分钟且有未处理的消息
			if len(consumeMSG) > 0 {
				//进行异步消费
				go func() {
					ch <- struct{}{}
					lock.Lock()
					m := consumeMSG
					fn(m)
					lock.Unlock()
				}()
				<-ch
				// 更新上次消费时间
				lastConsumeTime = time.Now()
				// 清空插入的数据
				consumeMSG = nil
			}
		}

	}

}

func fn(m []MSG) {
	fmt.Printf("本次消费了%d条消息\n", len(m))
}
```
