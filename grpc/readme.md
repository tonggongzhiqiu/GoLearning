# grpc 模块介绍
学习 grpc 基本功能

# 基本流程
1. 编写 helloworld proto 文件，通过 proto 命令生成 pb go 文件
2. 编写服务端代码
   1. 定义结构体，实现 pb go 中接口方法
   2. 监听端口
   3. 注册 server
   4. 执行 Serve 函数
3. 编写客户端代码
   1. 创建连接
   2. 创建 client
   3. 调用 sayHello 方法，打印结果

# proto 使用
## 生成命令解释
```
protoc --go_out=plugins=grpc:. *.proto
```

- --go_out: 指定 go 语言代码
- plugins=grpc 子选项，启动 grpc 插件。这个设置意味着要生成支持 grpc 插件的代码
- ：输出路径分隔符
- 。表示要生成到当前目录中

- *.proto 要编译的 proto 文件，可以指定具体的，也可以使用通配符