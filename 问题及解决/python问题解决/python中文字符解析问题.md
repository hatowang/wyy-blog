### 问题描述
````
将str类型的keyvalue，执行keyvalue.decode()
UnicodeDecodeError: 'ascii' codec can't decode byte 0xe7 in position 5587: ordinal not in range(128)
`````

### 原因：
````
python的str默认是ascii编码，和unicode编码冲突。
````

### 解决方法：
````
# 导入sys模块
import sys
# 重新加载sys模块
reload(sys)
# 设置编码格式为utf-8
sys.setdefaultencoding('utf8')
`````