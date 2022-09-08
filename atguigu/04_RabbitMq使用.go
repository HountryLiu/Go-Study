package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	//连接
	//  amqp://账号:密码@ip:port//vhost名称
	conn, err := amqp.Dial("amqp://admin:123123@127.0.0.1:5672//myvhost")
	defer conn.Close()
	fmt.Println(err)

	//打开通道
	ch, err := conn.Channel()
	defer ch.Close()

	//限流：确保消费者每次只能消费一个任务，消费完成后再分配任务，ack后再继续接收任务
	//设置每次从消息队列获取任务的数量
	err = ch.Qos(
		1,     //预取任务数量，这里可以理解为线程的数量
		0,     //预取大小
		false, //全局设置
	)

	if err != nil {
		//无法设置Qos
		return
	}
	// 声明队列
	//											durable队列持久化
	queue, err_q := ch.QueueDeclare("mysql_queue", true, false, false, false, nil)
	fmt.Println(err_q)

	// 生产任务：生产者
	ch.Publish("", queue.Name, false, false, amqp.Publishing{
		ContentType:  "text/plain",
		Body:         []byte("hello world"),
		DeliveryMode: amqp.Persistent, //消息持久化
	})

	// 消费者
	msgs, err_c := ch.Consume("mysql_queue", "my_consumer", false, false, false, false, nil)
	fmt.Println(err_c)

	for msg := range msgs { // chan类型
		// DeliveryTag:唯一标识
		fmt.Println(msg.DeliveryTag, string(msg.Body))
		msg.Ack(true) //消息消费成功后，返回true确认消费
	}
}
