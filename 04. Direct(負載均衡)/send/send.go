package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {

	setRoutingKey := "RoutingKeyTest.1"
	setExchangeName := "directExchangeTest"

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

	//步驟3 : 宣告 所使用Exchange (Rabbitmq依情況狀況自動創建Exchange)
	err = channel.ExchangeDeclare(
		setExchangeName,
		"direct", // exchange Type : 交換機模式 fanout/direct/topic/headers
		false,    // durable :關機是否保存Queue
		false,    // autoDelete :當最後消費完成後自我刪除Queue 
		false,    // exclusive :Queue是否獨佔(僅一個連線能使用其他人使用則回報RESOURCE_LOCKED錯誤)
		false,    //noWait
		nil,      //arguments
	)

	//步驟4 : 發送訊息
	for index := 0; index < 10; index++ {
		err = channel.Publish(
			setExchangeName, // exchange : 交換機名稱
			setRoutingKey,      // routing key : queue 綁定
			false,           // mandatory
			false,           // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte("全體廣播! > setRoutingKey : " + setRoutingKey + " > " + fmt.Sprint(+index)),
			})
	}

	log.Printf("Message to %v sent.\n", setExchangeName)
}
