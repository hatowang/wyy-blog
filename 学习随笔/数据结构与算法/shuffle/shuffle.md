## 洗牌算法
### 一、
#### 1. 代码
```
//洗牌算法，用于负载均衡
func shuffle(slice []int) {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
}
````

##### 2. 错误的洗牌导致的负载不均衡
````
1. 没有随机种子。在没有随机种子的情况下，rand.Intn()返回的伪随机数序列是固定的。
   
2.   洗牌不均匀，会导致整个数组第一个节点有大概率被选中，并且多个节点的负载分布不均衡。
   
````

### 二、改进算法（fisher-yates算法）
#### 1. 思路：每次随机挑选一个值放在数组末尾。然后在n-1个元素的数组中再随机挑选一个值，放在数组末尾，以此类推.
````
func shuffle1(indexes  []int) {
	for i := len(indexes); i > 0; i-- {
		lastIdx := i -1
		sed := rand.Intn(i)
		indexes[lastIdx], indexes[sed] = indexes[sed], indexes[lastIdx]
	}
}
`````

