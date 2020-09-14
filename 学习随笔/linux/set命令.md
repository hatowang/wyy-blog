## 一、 set -e（set -o errexit）
正常情况下遇到错误会继续执行下条命令，加上set -e 遇到错误即停止执行：set.sh
````
#!/usr/bin/env bash
set -e
foo
echo hello
`````
set.sh: line 3: foo: command not found

错误提示，然后停止执行
````
#!/usr/bin/env bash
set +e
foo
echo hello
set -e
````
set -e根据返回值来判断，一个命令是否运行失败。但是，某些命令的非零返回值可能不表示失败，或者开发者希望在命令失败的情况下，脚本继续执行下去。这时可以暂时关闭set -e，该命令执行结束后，再重新打开set -e
`

