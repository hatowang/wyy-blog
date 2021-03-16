## golang context包
### 一、概念及作用
````
context用于并发控制，一方面用于当请求超时获取小时，goroutine马上退出并释放资源，另一方面context本身含义就是上下文，可以用在多个goroutine间传递共享信息。
````

### 二、Context使用示例
#### 1. 超时控制
    //使用context.WithTimeout定义超时时间
    package main
    
    import (
    	"context"
    	"time"
    	"fmt"
    )
    
    func main() {
    	now := time.Now().UTC()
    	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
    	defer  cancel()
    	go loopForTimeOut(ctx,  now)
    	time.Sleep(5 * time.Second)
    }
    
    func loopForTimeOut(ctx context.Context, now time.Time) {
    	select {
    	case <-ctx.Done():
    		since := time.Since(now)
    		fmt.Println("since:", since)
    		fmt.Println("ctx timeout!", ctx.Err())
    	}
    }
````
    //结果
    since: 3.0006888s
    ctx timeout! context deadline exceeded
````
#### 2. context超时，读取github信息
````
    package main
    
    import (
        "context"
        "time"
        "net/http"
        "io/ioutil"
        "fmt"
    )
    
    func main() {
        ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
        defer  cancel()
        select {
        case <- ctx.Done():
            fmt.Println("ctx timeout!")
        case data := <-queryGitInfo(ctx,"hatowang/wyy-blog"):
            fmt.Println(string(data))
        }
    }
    
    func queryGitInfo(ctx context.Context, pro string) <- chan string{
        stop := make(chan string, 1)
        go func() {
            url := "https://api.github.com/repos/" + pro
            req, err := http.NewRequest("GET", url, nil)
            if err != nil {
                return
            }
            req = req.WithContext(ctx)
            client := &http.Client{}
            resp, err := client.Do(req)
            if err != nil {
                return
            }
            data, err := ioutil.ReadAll(resp.Body)
            if err != nil {
                return
            }
            defer resp.Body.Close()
            stop <- string(data)
        }()
        return stop
    }
````

#### 3. 利用context的WithValue方法传递值
##### 3.1 Context接口
```
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
Deadline：若ctx调用WithTimeout设置超时时间，则返回超时时间和true，否则返回当前时间和false；
Done：返回关闭的通道
Err： 若ctx通道未关闭，则返回nil，如果超时返回context deadline exceeded,若被取消，返回context canceled
Value：返回WithValue绑定的KV数据
```
#### 4. 实现Context接口的类型

    Context一共有4个类型实现了Context接口, 分别是emptyCtx, cancelCtx,timerCtx,valueCtx。每个类型都关联的创建方法。
    emptyCtx：用于创建根context，一般用于主函数，初始化和测试中，使用Background或TODO创建
    cancelCtx：cancelCtx支持取消操作，取消同时也会对实现了canceler接口的子代进行取消操作
    timerCtx: 可取消的context，过期或超时会自动取消
    valueCtx: 可存储共享信息的context
    
#### 5. Context使用规范
    1. 不要将context作为结构体的一个字段存储，将其作为第一个参数并命名为ctx
    2. 不要传递nil 作为context给一个函数，如果不确定则传递TODO
    3. context时并发安全的（互斥锁），相同的context可以运行在不同的goroutine中