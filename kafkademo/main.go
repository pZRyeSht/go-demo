package main

import (
	"github.com/EscAlice/go-demo/kafkademo/consumer"
	"github.com/EscAlice/go-demo/kafkademo/producer"
)

func main() {
	topic := "test_topic"
	address := "127.0.0.1:9092"
	group := "test_group"
	handler := consumer.NewMessageHandler()
	cons := consumer.NewMessageConsumer([]string{address}, group, []string{topic}, handler)
	defer cons.Stop()
	go cons.Consume()
	
	prod := producer.NewMessageProducer([]string{address}, topic)
	for i := 0; i < 10; i++ {
		msg := &consumer.MessageData{
			ID:   int64(i),
			Data: "this is a test message",
		}
		prod.Producer(msg)
	}
	select {}
}