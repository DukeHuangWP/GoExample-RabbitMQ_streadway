package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {

	setRoutingKey := "RoutingKeyTest.1"
	setExchangeName := "directExchangeTest"
	setQueueName := "fanoutQueueTest"

	//步驟1 : 連接 amqp host
	conn, err := amqp.Dial("amqp://user01:user01@10.211.55.2:5672/")
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	//步驟2 : 連接 Rabbitmq Channel
	channel, err := conn.Channel()
	if err != nil {
		log.Panic(err)
	}
	defer channel.Close()

	//步驟3 : 宣告 所使用Exchange
	err = channel.ExchangeDeclare(
		setExchangeName,
		"direct",
		false, // durable 關機是否保存Queue
		false, // autoDelete 當最後消費完成後自我刪除Queue 
		false, // exclusive :Queue是否獨佔(僅一個連線能使用其他人使用則回報RESOURCE_LOCKED錯誤)
		false, // no-wait
		nil,   // arguments
	)

	//步驟4 : 宣告 所使用Queue (如果所使用Queue未命名，則系統將會隨機給予新的Queue名稱)
	queue, err := channel.QueueDeclare(
		setQueueName,    // queue : Queue名稱
		false, // durable 關機是否保存Queue
		false,  // autoDelete 當最後消費完成後自我刪除Queue
		false, // exclusive :Queue是否獨佔(僅一個連線能使用其他人使用則回報RESOURCE_LOCKED錯誤)
		false, // no-wait
		nil,   // arguments
	)

	//步驟5 : 綁定Queue 與 Exchange (若已綁定則可跳過此步驟)
	err = channel.QueueBind(
		queue.Name,      // queue : Queue名稱
		setRoutingKey,      // routing key : Queue Routing Key
		setExchangeName, // exchange
		false,
		nil)

	err = channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	//步驟6 : 消費Queue中的message
	msgs, err := channel.Consume(
		queue.Name, // queue : Queue名稱
		"",         // consumer
		false,       // auto-ack  : 是否自動告知Queue Message已遭消費
		false,      // exclusive : Queue是否獨佔(僅一個連線能使用其他人使用則回報RESOURCE_LOCKED錯誤)
		false,      // no-local
		false,      // noWait
		nil,        // arguments
	)

	//步驟7 : 解析所消費的message
	go func() {
		for delivery := range msgs {
			log.Printf("Received a message: %s", delivery.Body)
			delivery.Ack(false)
		}
	}()

	log.Printf("> Exchange:%v Queue:%v\n", setExchangeName, queue.Name)
	select {}
}
