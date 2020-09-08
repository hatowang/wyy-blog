### RPC failed; curl 18 transfer closed with outstanding read data remaining

描述：在git clone项目时无论怎么配置git，都不行

解决：

````````
$git clone https://github.com/xxx/xxxxxxx.git --depth 1
此命令只能clone最近一次master的commit
若想clone其他分支，则使用-b指定clone某分支的最近一次commit：
$git clone -b remote_branch_name https://github.com/xxx/xxxxxxx.git --depth 1 
或者使用以下命令切换分支：
$git remote set-branches origin 'remote_branch_name'
$git fetch --depth 1 origin remote_branch_name
$git checkout remote_branch_name
