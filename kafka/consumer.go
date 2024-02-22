package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var ConsumerId = 1

type ConsumerClient struct {
	Consumer   sarama.Consumer
	Topic      string
	ConsumerId int
}

func (c *ConsumerClient) InitConsumerClient() error {
	consumer, err := sarama.NewConsumer([]string{""}, nil)
	if err != nil {
		return err
	}

	c.Consumer = consumer
	c.Topic = Topic
	c.ConsumerId = ConsumerId
	ConsumerId++
	return nil
}

func (c *ConsumerClient) GetMessage(partitionId int32, offset int64) {
	if offset == -1 {
		offset = sarama.OffsetNewest
	}
	pc, err := c.Consumer.ConsumePartition(c.Topic, partitionId, offset)
	if err != nil {
		fmt.Printf("failed to start consumer for partition %d,err:%v", partitionId, err)
		//That topic/partition is already being consumed
		return
	}

	// 异步从每个分区消费信息
	go func(sarama.PartitionConsumer) {
		for msg := range pc.Messages() {
			fmt.Printf("ConsumerId:%d Partition:%d Offset:%d Key:%v Value:%v", c.ConsumerId, msg.Partition, msg.Offset, msg.Key, string(msg.Value))
		}
	}(pc)
}

// GetMessageToAll 遍历所有分区
func (c *ConsumerClient) GetMessageToAll(offset int64) {
	partitionList, err := c.Consumer.Partitions(c.Topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v", err)
		return
	}
	fmt.Printf("所有partition:%v", partitionList)

	for partition := range partitionList { // 遍历所有的分区
		c.GetMessage(int32(partition), offset)
	}
}
