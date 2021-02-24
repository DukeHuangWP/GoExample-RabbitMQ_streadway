A golang example of RabbitMQ tutorials.

## 前言
<!-- <img src="https://raw.githubusercontent.com/DukeHuangWP/HackMD/2021-01-01/RabbitMQ%20%E5%AD%B8%E7%BF%92%E7%AD%86%E8%A8%98/RabbitMQ_Flow_Diagram.svg">-->

此為RabbitMQ Golnag使用範例:
```
01. Fanout(全體廣播)
02. Fanout(全體廣播+負載均衡)
03. Direct(指定路由)
04. Direct(負載均衡)
05. Topic(訂閱消息)
```
<!-- <img src="https://raw.githubusercontent.com/DukeHuangWP/HackMD/2021-01-01/RabbitMQ%20%E5%AD%B8%E7%BF%92%E7%AD%86%E8%A8%98/RabbitMQ_Flow_Diagram.svg">-->
<!-- <img src="https://raw.githubusercontent.com/DukeHuangWP/GoExample-RabbitMQ_streadway/master/02.%20Fanout(%E5%85%A8%E9%AB%94%E5%BB%A3%E6%92%AD%2B%E8%B2%A0%E8%BC%89%E5%9D%87%E8%A1%A1).svg">-->
<!-- <img src="https://raw.githubusercontent.com/DukeHuangWP/GoExample-RabbitMQ_streadway/master/03.%20Direct(%E6%8C%87%E5%AE%9A%E8%B7%AF%E7%94%B1).svg">-->
<!-- <img src="https://raw.githubusercontent.com/DukeHuangWP/GoExample-RabbitMQ_streadway/master/04.%20Direct(%E8%B2%A0%E8%BC%89%E5%9D%87%E8%A1%A1).svg">-->
<!-- <img src="https://raw.githubusercontent.com/DukeHuangWP/GoExample-RabbitMQ_streadway/master/05.%20Topic(%E8%A8%82%E9%96%B1%E6%B6%88%E6%81%AF).svg">-->


## 設置負載均衡關鍵
```go
err = channel.Qos(
    1,     // prefetch count
    0,     // prefetch size
    false, // global
)

msgs, err := channel.Consume(
...
false,       // auto-ack  : 是否自動告知Queue Message已遭消費
...


for delivery := range msgs {
    ...
    delivery.Ack(false)
    ...
}
```
