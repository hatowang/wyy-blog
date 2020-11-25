### 实验笔记    

#### 1. namespace
    ip netns add netns1 //创建network namespace netns1
    cd /var/run/netns/ //查看netns1挂载
    ip netns exec netns1 ip link list //查看链路层
    ip netns list //查看network namespace
    ip netns exec netns1 ping 127.0.0.1 //进入，ping本地回环
    ip netns exec netns1 ip link set dev lo up //启动本地回环链路
    ip netns exec netns1 ping 127.0.0.1
    ip link add veth0 type veth peer name veth1 //创建veth pair
    ip link set veth1 netns netns1 //将veth1 set到netns1
    ip netns exec netns1 ifconfig veth1 10.1.1.1/24 up //设置veth1的ip
    ifconfig veth0 10.1.1.2/24 up //设置veth0的ip
    ifconfig //查看多了veth0网卡
    ip netns exec netns1 ifconfig //查看netns1中的veth1网卡
    ip netns exec netns1 ping 10.1.1.2 
    ip netns exec netns1 ping 10.1.1.1
    ip netns exec netns1 route //查看路由
    ip netns exec netns1 iptables -L //查看iptables
    ip netns exec netns1 ip link set veth1 netns 1 //netns1 network namespace下的veth1网卡挪到PID为1的进程（即init进程）所在的network namespace
    
#### 2. veth pair 
    ip link add veth0 type veth peer name veth1 //创建veth pair
    ip link list // 查看网卡
    ip link set dev veth1 up //打开网卡
    ifconfig veth0 ip设置网络
    
* veth pair设备的原理较简单，就是向veth pair设备的一端输入数据，数据通过内核协议栈后从veth pair的另一端出来
##### 容器与host veth pair的关系(veth pair+bridge)
````    
查看容器中网卡和宿主机的关系：
一、
1. 容器中： cat /sys/class/net/eth0/iflink
2. 主机中： /sys/claas/net/veth*/iflink 
3. 找到值相同的即为对应的网卡
二、
1.容器中执行： ip link show eth* 得到成对的网卡的index
2.在宿主机执行： ip link show | grep index 找到该网卡 
三、 ethtool-S
    