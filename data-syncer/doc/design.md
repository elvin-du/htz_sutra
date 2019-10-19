## 模块
### service

主要逻辑模块。调用`mdu/rpc`模块获取区块链数据，然后进行一定的逻辑运算后，把数据组合调用`common/model`模块把数据写入到数据库。

### task

定时任务。定时调用`service`模块的方法进行数据同步。

### mdu/rpc

主要负责从区块链那边获取数据。

### 其他

* 日志
    + `output`: 只能是`file`或者`std`。填写`file`时，系统会可执行程序程序的根目录生成一个名为`log.txt`的文件用来存放日志。
    + `level`：只能是`panic`, `fatal`,`error`,`warn`,`info`,`debug`,`trace`的其中之一。危险级别从高到低排序。
            