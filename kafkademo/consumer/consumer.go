package consumer

import (
	"context"
	"encoding/json"
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
	var bsList []sarama.BalanceStrategy
	bsList = append(bsList, sarama.BalanceStrategyRange)
	conf.Consumer.Group.Rebalance.GroupStrategies = bsList
	conf.Consumer.Offsets.Retry.Max = 3
	conf.Consumer.Offsets.Initial = sarama.OffsetOldest
	conf.Consumer.Offsets.AutoCommit.Enable = true
	conf.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second

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
	for message := range claim.Messages() {
		var msg MessageData
		if err := json.Unmarshal(message.Value, &msg); err != nil {
			return err
		}
		// handler data logic
		log.Println("consumer message success :", msg)
		// mark message and auto commit
		session.MarkMessage(message, "")
	}
	return nil
}

type MessageData struct {
	ID   int64  `json:"id"`
	Data string `json:"data"`
}
