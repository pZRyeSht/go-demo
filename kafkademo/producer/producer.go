package producer

import (
	"encoding/json"
	"log"
	"time"

	"github.com/EscAlice/go-demo/kafkademo/consumer"
	"github.com/Shopify/sarama"
)

type MessageProducer struct {
	topic    string
	producer sarama.AsyncProducer
}

func NewMessageProducer(address []string, topic string) *MessageProducer {
	conf := sarama.NewConfig()
	conf.Producer.Return.Errors = true
	conf.Producer.Return.Successes = true
	conf.Producer.Retry.Max = 3 // 最大重试次数
	conf.Producer.Retry.Backoff = 100 * time.Millisecond // 两次重试之间等待群集稳定的时间，默认值为100ms
	conf.Producer.RequiredAcks = sarama.WaitForAll // 发送完数据需要leader和follow都确认，才会返回
	conf.Producer.Partitioner = sarama.NewHashPartitioner // 根据hash算法分区

	asyncProducer, err := sarama.NewAsyncProducer(address, conf)
	if err != nil {
		panic(err)
	}

	mp := &MessageProducer{
		topic:    topic,
		producer: asyncProducer,
	}
	go mp.dealMessage()
	return mp
}

func (p *MessageProducer) dealMessage() {
	for {
		select {
		case res := <-p.producer.Successes():
			log.Println("push message success: ", res)
		case err := <-p.producer.Errors():
			log.Println("push message error: ", err.Error())
		}
	}
}

func (p *MessageProducer) Producer(data *consumer.MessageData) {
	byt, err := json.Marshal(data)
	if err != nil {
		return
	}
	p.producer.Input() <- &sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(byt),
	}
}
