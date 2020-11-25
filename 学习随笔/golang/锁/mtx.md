## 锁
### 一、未加锁场景

#### 1. 直接而上代码：
```
package main

import "sync"



func main() {
	for j := 0; j < 3; j++ {
		var wg sync.WaitGroup
		// 变量
		var counter int
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				counter++
			}()
		}
		wg.Wait()
		println(counter)
	}
}

```

#### 2. 多次执行结果,count值不定
```
GOROOT=C:\Go #gosetup
GOPATH=E:\Goworkspace #gosetup
C:\Go\bin\go.exe build -i -o C:\Users\wangyeyu\AppData\Local\Temp\___go_build_main_go__3_.exe E:/Goworkspace/src/GoPro/test/main.go #gosetup
"D:\program\GoLand 2018.1\bin\runnerw.exe" C:\Users\wangyeyu\AppData\Local\Temp\___go_build_main_go__3_.exe #gosetup
966
930
921

Process finished with exit code 0
```

### 二、 加锁场景
#### 1. 上代码：

```
package main

import "sync"

var mtx sync.Mutex

func main() {
	for j := 0; j < 3; j++ {
		var wg sync.WaitGroup
		// 变量
		var counter int
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				mtx.Lock()
				counter++
				mtx.Unlock()
			}()
		}
		wg.Wait()
		println(counter)
	}
}
```

#### 2. 多地执行结果相同

```
GOROOT=C:\Go #gosetup
GOPATH=E:\Goworkspace #gosetup
C:\Go\bin\go.exe build -i -o C:\Users\wangyeyu\AppData\Local\Temp\___go_build_main_go__3_.exe E:/Goworkspace/src/GoPro/test/main.go #gosetup
"D:\program\GoLand 2018.1\bin\runnerw.exe" C:\Users\wangyeyu\AppData\Local\Temp\___go_build_main_go__3_.exe #gosetup
1000
1000
1000

Process finished with exit code 0
```

### 三、 trylock

#### 1.实现原理，利用channel读写阻塞的原理，在创建锁的时候给一个钥匙，lock的时候拿走，unlock的时候放入。拿到锁了count+1，否则返回

```
package main

import "sync"

// Lock try lock
type Lock struct {
	c chan struct{}
}

func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

func (l Lock) TryLock() bool {
	lockResult := false

	select {
	//channel队列内有钥匙，则可加锁，并取走钥匙
	case <-l.c:
		lockResult = true
	default:
	}
	return lockResult
}

func (l Lock) Unlock() {
	//放入钥匙
	l.c <- struct{}{}
}

// 变量
var counter int

func main() {
	var l = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.TryLock() {
				// log error
				println("lock failed")
				return
			}
			counter++
			println("current counter", counter)
			l.Unlock()
		}()
	}
	wg.Wait()
}
```

#### 2. 输出结果：多次输出同一count
```
GOROOT=C:\Go #gosetup
GOPATH=E:\Goworkspace #gosetup
C:\Go\bin\go.exe build -i -o C:\Users\wangyeyu\AppData\Local\Temp\___go_build_main_go__3_.exe E:/Goworkspace/src/GoPro/test/main.go #gosetup
"D:\program\GoLand 2018.1\bin\runnerw.exe" C:\Users\wangyeyu\AppData\Local\Temp\___go_build_main_go__3_.exe #gosetup
1000
1000
1000

Process finished with exit code 0
```

#### 3. 活锁：上述try lock场景会造成大量goroutine抢锁，CPU被浪费，单机情况不建议使用。

### 四、 分布式锁-基于redis的setnx

#### 1.代码
```
package main

import (
	"sync"
	"github.com/go-redis/redis"
	"time"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
		wg.Wait()
	}
}

func incr() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	var lockKey = "count_key"
	var countKey = "count"
	resp := client.SetNX(lockKey, 1, time.Second*5)
	lockSuccess, err := resp.Result()
	if err != nil || !lockSuccess {
		fmt.Println(err)
		fmt.Println("lock faild !")
		return
	}
	getResp := client.Get(countKey)
	cValue, err := getResp.Int64()
	if err == nil {
		cValue ++
		resp := client.Set(countKey, cValue, 0)
		_, err := resp.Result()
		if err != nil {
			fmt.Println("set count failed!")
		}
	}
	fmt.Print("current value is ")
	fmt.Println(cValue)
	delResp := client.Del(countKey)
	unLockSuccess, err := delResp.Result()
	if err != nil && unLockSuccess <= 0 {
		fmt.Println(err)
		fmt.Println("unlock failed!")
	}
}
```

#### 2. setnx
    远程调用setnx实际上和单机的trylock非常相似，如果获取锁失败，那么相关的任务逻辑就不应该继续向前执行;

    setnx很适合在高并发场景下，用来争抢一些“唯一”的资源。比如交易撮合系统中卖家发起订单，而多个买家会对其进行并发争抢。这种场景我们没有办法依赖具体的时间来判断先后，因为不管是用户设备的时间，还是分布式场景下的各台机器的时间，都是没有办法在合并后保证正确的时序的。哪怕是我们同一个机房的集群，不同的机器的系统时间可能也会有细微的差别。
    
