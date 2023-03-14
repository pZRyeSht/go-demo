# sarama 库的问题：阿里云的建议
阿里云官方文档不推荐使用 sarama 库，见此文：[为什么不推荐使用 Sarama Go 客户端收发消息](https://help.aliyun.com/document_detail/266782.html)
## 不推荐的原因
Sarama Go 客户端存在以下已知问题：
- 当 Topic 新增分区时，Sarama Go 客户端无法感知并消费新增分区，需要客户端重启后，才能消费到新增分区
- 当 Sarama Go 客户端同时订阅两个以上的 Topic 时，有可能会导致部分分区无法正常消费消息
- 当 Sarama Go 客户端的消费位点重置策略设置为 Oldest(earliest) 时，如果客户端宕机或服务端版本升级，由于 Sarama Go 客户端自行实现 OutOfRange 机制，有可能会导致客户端从最小位点开始重新消费所有消息
## 对应的解决方案
建议尽早将 Sarama Go 客户端替换为 [Confluent Go](https://github.com/confluentinc/confluent-kafka-go) 客户端，后续参考其封装下。
如果无法在短期内替换客户端，请注意以下事项：
- 针对生产环境，请将位点重置策略设置为 Newest（latest）
- 针对测试环境，或者其他明确可以接收大量重复消息的场景，设置为 Oldest（earliest）
- 如果发生了位点重置，产生大量堆积，可以使用消息队列 Kafka 版控制台提供的重置消费位点功能，手动重置消费位点到某一时间点，无需改代码或换 Consumer Group（如果不重置的话会产生大量的重复数据）

# TODO
- confluentinc/confluent-kafka-go Demo
- segmentio/kafka-go Demo