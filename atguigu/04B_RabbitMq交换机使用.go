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

	// 创建两个队列
	queue1, _ := ch.QueueDeclare("first_queue", true, false, false, false, nil)
	queue2, _ := ch.QueueDeclare("second_queue", true, false, false, false, nil)

	// 创建两个交换机
	ch.ExchangeDeclare("frist_exchange", "direct", true, false, false, false, nil)
	ch.ExchangeDeclare("second_exchange", "direct", true, false, false, false, nil)

	// queue和交换机绑定
	ch.QueueBind(queue1.Name, "frist_routingKey", "frist_exchange", false, nil)
	ch.QueueBind(queue2.Name, "second_routingKey", "second_exchange", false, nil)
	// 第一个参数是队列名称，第二个参数是routingKey,第三个参数是交换机名称

	// 生产者：第一个参数是交换机名称，第二个参数routingKey
	ch.Publish("frist_exchange", "frist_routingKey", false, false, amqp.Publishing{
		ContentType:  "text/plain",
		Body:         []byte("hello world"),
		DeliveryMode: amqp.Persistent,
	})

	// // 消费者：使用queue name
	// msgs, _ := ch.Consume("first_queue", "my_consumer", false, false, false, false, nil)

	// for msg := range msgs { // chan类型
	// 	// DeliveryTag:唯一标识
	// 	fmt.Println(msg.DeliveryTag, string(msg.Body))
	// 	// msg.Ack(true) //消息消费成功后，返回true确认消费
	// }
}
