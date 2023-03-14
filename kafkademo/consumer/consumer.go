package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

type MessageConsumer struct {
	topics  []string
	ctx     context.Context
	cancel  context.CancelFunc
	group   sarama.ConsumerGroup
	handler sarama.ConsumerGroupHandler
}

func NewMessageConsumer(address []string, group string, topics []string, handler *MessageHandler) *MessageConsumer {
	conf := sarama.NewConfig()
	conf.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRange} // 重平衡策略
	conf.Consumer.Offsets.Retry.Max = 3 // 最大重试次数
	conf.Consumer.Offsets.Initial = sarama.OffsetNewest // 从最新的offset偏移量开始
	conf.Consumer.Offsets.AutoCommit.Enable = true // 开启自动提交，需要手动调用MarkMessage才有效
	conf.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second // 自动提交间隔

	consumerGroup, err := sarama.NewConsumerGroup(address, group, conf)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &MessageConsumer{
		topics:  topics,
		ctx:     ctx,
		cancel:  cancel,
		group:   consumerGroup,
		handler: handler,
	}
}

func (c *MessageConsumer) Consume()  {
	for {
		select {
		case <-c.ctx.Done():
			if err := c.group.Close(); err != nil {
				return
			}
			return
		default:
			if err := c.group.Consume(c.ctx, c.topics, c.handler); err != nil {
				log.Println("Consume error:", err.Error())
			}
		}
	}
}

func (c *MessageConsumer) Stop() { c.cancel()}



type MessageHandler struct{}

func NewMessageHandler() *MessageHandler { return &MessageHandler{}}

func (m MessageHandler) Setup(sarama.ConsumerGroupSession) error { return nil }

func (m MessageHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (m MessageHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// 1.自动提交示例
	for message := range claim.Messages() {
		var msg MessageData
		if err := json.Unmarshal(message.Value, &msg); err != nil {
			return err
		}
		// handler data logic eg:mysql
		fmt.Println("consumer message success :", msg)
		// mark message and auto commit
		session.MarkMessage(message, "")
	}
	// 2.手动提交示例
	consumerCount := 0
	for message := range claim.Messages() {
		var msg MessageData
		if err := json.Unmarshal(message.Value, &msg); err != nil {
			return err
		}
		// handler data logic eg:mysql
		fmt.Println("consumer message success :", msg)
		// 手动提交模式下，也需要先进行标记
		session.MarkMessage(message, "")
		
		consumerCount++
		// 假设每消费 3 条数据 commit 一次
		if consumerCount%3 == 0 {
			// 手动提交，不能频繁调用
			t1 := time.Now().Nanosecond()
			session.Commit()
			t2 := time.Now().Nanosecond()
			fmt.Println("commit cost:", (t2-t1)/(1000*1000), "ms")
		}
	}
	return nil
}

type MessageData struct {
	ID   int64  `json:"id"`
	Data string `json:"data"`
}
