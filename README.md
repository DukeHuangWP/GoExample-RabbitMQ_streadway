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

![Alt text]("./01. Fanout(全體廣播).svg")
![Alt text]("./02. Fanout(分區廣播+負載均衡)")
![Alt text]("./03. Direct(指定路由).svg")
![Alt text]("./04. Direct(負載均衡).svg")
![Alt text]("./05. Topic(訂閱消息).svg")
d

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