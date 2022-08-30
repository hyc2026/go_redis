# 使用Golang实现简易redis服务器

## TCP服务器

[代码实现](https://github.com/hyc2026/go_redis/tree/master/tcp)

早期的 Tomcat/Apache 服务器使用的是阻塞 IO 模型。它使用一个线程处理一个连接，在没有收到新数据时监听线程处于阻塞状态，直到数据就绪后线程被唤醒。因为阻塞 IO 模型需要开启大量线程并且频繁地进行上下文切换，所以效率很差。

IO 多路复用技术为了解决上述问题采用了一个线程监听多路连接的方案。一个线程持有多个连接并阻塞等待，当其中某个连接可读写时线程被唤醒进行处理。因为多个连接复用了一个线程所以 IO 多路复用需要的线程数少很多。

Golang 的 `netpoller` 基于IO多路复用和 goroutine scheduler 构建了一个简洁高性能的网络模型，并给开发者提供了 `goroutine-per-connection` 风格的极简接口。



## 备注
`go mod init 项目名`后可以引入当前路径下的包