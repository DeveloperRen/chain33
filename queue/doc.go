package queue

/*
=======
/*
>>>>>>> 16b55f7927e9031b05202f9eeee0bf8e538f339c
主要的功能：实现消息队列的功能。
设计这个模块的原因：
为系统的分布式化，微服务化做准备。
每个模块相对来说独立，不是通过接口调用，而是通过消息进行通信。
queue 的主要接口:
type IClient interface {
	Send(msg Message) (err error) //异步发送消息
	Wait()                        //等待消息处理完成
	Recv() chan Message
	Sub(topic string) (ch chan Message) //订阅消息
}
*/
