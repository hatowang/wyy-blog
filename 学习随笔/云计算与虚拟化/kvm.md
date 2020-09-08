### KVM
概念
- KVM（kernel－based Virtual Machine)：开源免费，架构简单，性能卓越，社区活跃
- 基于硬件辅助的开源虚拟化解决方案
- 基于Linux内核的虚拟化技术, 可以直接将Linux内核转换为Hypervisor,）从而使得Linux内核能够直接管理虚拟机,直接调用Linux内核中的内存管理、进程管理子系统来管理虚拟机

架构
- 两大方面：KVM模块以及QEMU-KVM
- kvm-ko(主要模块),kvm_intel.ko,kvm_amd.ko
- QEMU-KVM通过修改qemu代码而得出的专门为管理和创建虚机的管理工具
- /dev/kvm :Linux系统下kvm提供的驱动接口
- kvm(内核空间)提供VCPU、vMEM，既支持linux，也支持windows
- QEMU-KVM(用户空间)提供IO Device

KVM和QEMU-KVM交互
- QEMU-KVM是kvm团队针对qemu的改善和二次开发的一套工具
- /dev/kvm是kvm内核模块提供给用户空间的一个接口，这个接口被qemu-kvm调用，通过ioctl系统调用给用户提供删除、创建、管理虚机的工具
- qemu-kvm就是通过open()、close()、ioctl()，等方法去打开，关闭和调用这个接口，实现跟KVM的互动

[qemu-kvm 调用kvm的过程](picture/qemu-kvm调用kvn.png)
- 打开/dev/kvm设备
- 通过KVM_CREATE_VM创建一个虚拟机对象
- 通过KVM_CREATE_VCPU为虚机创建VCPU对象
- 通过KVM_RUN设置VCPU运行起来


搭建KVM
- 确认硬件服务器是否支持
 egrep 'svm|vmx' /proc/cpuinfo
- 安装包准备，yum源配置
- yum安装
- 创建vm