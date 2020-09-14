## 操作系统原理
1. 进程：有序的双向链表（Task_struct）
````
    进程描述符：进程元数据
    进程控制块：
    1. 进程切换上下文：Suspend、Resume（stack Pointer、other registers、EIP registers etc）
    类别：交互式进程（IO密集型）、批处理进程（cpu密集）、实时（Real-time）
       CPU密集型：时间片长、优先级低
       IO密集型：时间片短、优先级高
   
   2. Linux优先级：priority
    2.1 实时优先级：1-99，数字越小，优先级越低（RT）：ps -e -o class,cmd|grep sshd
        SCHED_FIFO:先进先出
        SCHED_RR:Round Robin
    2.2 静态优先级：100-139，数字越小、优先级越高：ps -e -o class,rtprio,nice,cmd
        SCHED_Other：用来调度100-139    
        实时优先级比静态优先级高
        nice值:调整静态优先级（-20-19）
             
   2.3 动态优先级：临时性调整优先级（长时间未运行，防止饿死，一直运行的进程，给予惩罚措施）
        dynamic priority = max（100，min（static priority - bonus + 5 ，139）
        bonus:0-10
        
   
   2.4 手动调整优先级：
       100-139：nice  
       chrt:调整实时优先级进程的优先级
       chrt -f -p [prio] pid
       chrt -r -p [prio] pid
       chrt -f -p [prio] cmd
   
   3. 优先级队列（同优先级的吸纳成放在一个队列），（活队列，过期队列，来回轮换 O(1)）
        CFS：Complete Fair Scheduler（SCHED_Other）
        
   4. COW:
    Kernel->initial
        init
            fork():系统调用
            task_struct
                Memory->Parent
                COW:Copy on Write
````